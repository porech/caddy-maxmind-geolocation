package caddy_maxmind_geolocation

import "strings"

type Names struct {
	De   string `maxminddb:"de"`
	En   string `maxminddb:"en"`
	Es   string `maxminddb:"es"`
	Fr   string `maxminddb:"fr"`
	PtBr string `maxminddb:"pt-BR"`
	Ru   string `maxminddb:"ru"`
	ZhCn string `maxminddb:"zh-CN"`
}

type Country struct {
	GeonameId         int    `maxminddb:"geoname_id"`
	IsInEuropeanUnion bool   `maxminddb:"is_in_european_union"`
	ISOCode           string `maxminddb:"iso_code"`
	Names             Names  `maxminddb:"names"`
}

type Subdivision struct {
	ISOCode string `maxminddb:"iso_code"`
}

type Subdivisions []Subdivision

func (s Subdivisions) GetISOCodes() []string {
	var result []string
	for _, sub := range s {
		result = append(result, sub.ISOCode)
	}
	return result
}
func (s Subdivisions) CommaSeparatedISOCodes() string {
	return strings.Join(s.GetISOCodes(), ",")
}

type Location struct {
	MetroCode int `maxminddb:"metro_code"`
}

type Record struct {
	Country                Country      `maxminddb:"country"`
	Location               Location     `maxminddb:"location"`
	Subdivisions           Subdivisions `maxminddb:"subdivisions"`
	AutonomousSystemNumber int          `maxminddb:"autonomous_system_number"`
}
