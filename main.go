package caddy_maxmind_geolocation

import (
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/oschwald/maxminddb-golang"
	"go.uber.org/zap"
	"net"
	"net/http"
)

// Interface guards
var (
	_ caddy.Module             = (*MaxmindGeolocation)(nil)
	_ caddyhttp.RequestMatcher = (*MaxmindGeolocation)(nil)
	_ caddy.Provisioner        = (*MaxmindGeolocation)(nil)
	_ caddy.CleanerUpper       = (*MaxmindGeolocation)(nil)
	_ caddyfile.Unmarshaler    = (*MaxmindGeolocation)(nil)
)

func init() {
	caddy.RegisterModule(MaxmindGeolocation{})
}

type MaxmindGeolocation struct {
	DbPath         string   `json:"db_path"`
	AllowCountries []string `json:"allow_countries"`
	DenyCountries  []string `json:"deny_countries"`

	dbInst *maxminddb.Reader
	logger *zap.Logger
}

/*
	The matcher configuration will have a single block with the following parameters:
	- `db_path`: required, is the path to the GeoLite2-Country.mmdb file
	- `allow_countries`: a space-separated list of allowed countries
	- `deny_countries`: a space-separated list of denied countries.

	You will want specify just one of `allow_countries` or `deny_countries`. If you
	specify both of them, denied countries will take precedence over allowed ones.
	If you specify none of them, all requests will be denied.
 */
func (m *MaxmindGeolocation) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	current := 0
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "db_path":
				current = 1
			case "allow_countries":
				current = 2
			case "deny_countries":
				current = 3
			default:
				switch current {
				case 1:
					m.DbPath = d.Val()
					current = 0
				case 2:
					m.AllowCountries = append(m.AllowCountries, d.Val())
				case 3:
					m.DenyCountries = append(m.DenyCountries, d.Val())
				default:
					return fmt.Errorf("unexpected config parameter %s", d.Val())
				}
			}
		}
	}
	return nil
}

func (MaxmindGeolocation) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.matchers.maxmind_geolocation",
		New: func() caddy.Module { return new(MaxmindGeolocation) },
	}
}

func (m *MaxmindGeolocation) Provision(ctx caddy.Context) error {
	var err error
	m.logger = ctx.Logger(m)
	m.dbInst, err = maxminddb.Open(m.DbPath)
	if err != nil {
		return fmt.Errorf("cannot open database file %s: %v", m.DbPath, err)
	}
	return nil
}

func (m *MaxmindGeolocation) Cleanup() error {
	if m.dbInst != nil {
		return m.dbInst.Close()
	}
	return nil
}

func (m *MaxmindGeolocation) Match(r *http.Request) bool {
	var record Record

	// If both the allow and deny fields are empty, let the request pass
	if len(m.AllowCountries) < 1 && len(m.DenyCountries) < 1 {
		return false
	}

	remoteIp, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		m.logger.Warn("cannot split IP address", zap.String("address", r.RemoteAddr), zap.Error(err))
	}

	// Get the record from the database
	addr := net.ParseIP(remoteIp)
	if addr == nil {
		m.logger.Warn("cannot parse IP address", zap.String("address", r.RemoteAddr))
		return false
	}
	err = m.dbInst.Lookup(addr, &record)
	if err != nil {
		m.logger.Warn("cannot lookup IP address", zap.String("address", r.RemoteAddr), zap.Error(err))
		return false
	}

	m.logger.Debug("Detected country", zap.String("country", record.Country.ISOCode))

	for _, denyCountry := range m.DenyCountries {
		if denyCountry == record.Country.ISOCode {
			return false
		}
	}

	for _, allowCountry := range m.AllowCountries {
		if allowCountry == record.Country.ISOCode {
			return true
		}
	}

	// If there are no allowed countries, the default is true (pass if the country is not denied).
	// If there are allowed countries, and they didn't match, return false.
	return len(m.AllowCountries) < 1
}
