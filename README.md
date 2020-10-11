# caddy-maxmind-geolocation

Caddy v2 module to filter requests based on source IP geographic location. This was a feature provided by the V1 `ipfilter`
middleware. 

## Installation

You can download a Caddy build with this plugin inside directly from the [official Caddy page](https://caddyserver.com/download).

If you prefer, you can build Caddy by yourself by [installing xcaddy](https://github.com/caddyserver/xcaddy) and running:
```
xcaddy build --with github.com/porech/caddy-maxmind-geolocation
```

## Requirements 

To be able to use this module you will need to have a Maxmind GeoLite2 database, that can be downloaded for free
by creating an account. More information about this are available on the
[official website](https://dev.maxmind.com/geoip/geoip2/geolite2/).

You will specifically need the `GeoLite2-Country.mmdb` file.

## Usage

You can use this module as a matcher to blacklist or whitelist a set of countries. Here are some samples:

### Caddyfile

1. Allow access to the website only from Italy and France:
```
test.example.org {
  @mygeofilter {
    maxmind_geolocation {
      db_path "/usr/share/GeoIP/GeoLite2-Country.mmdb"
      allow_countries IT FR
    }
  }

   file_server @mygeofilter {
     root /var/www/html
   }
}

```

2. Deny access to the website from Russia:
```
test.example.org {
  @mygeofilter {
    maxmind_geolocation {
      db_path "/usr/share/GeoIP/GeoLite2-Country.mmdb"
      deny_countries RU
    }
  }

   file_server @mygeofilter {
     root /var/www/html
   }
}

```

### API/JSON

1. Allow access to the website only from Italy and France:
```jsonc
{
  "apps": {
    "http": {
      "servers": {
        "myserver": {
          "listen": [":443"],
          "routes": [
            {
              "match": [
                {
                  "host": [
                    "test.example.org"
                  ],
		  "maxmind_geolocation": {
                    "db_path": "/usr/share/GeoIP/GeoLite2-Country.mmdb",
                    "allow_countries": [ "IT", "FR" ]
                  }
                }
              ],
              "handle": [
                {
                  "handler": "file_server",
                  "root": "/var/www/html"
                }
              ]
            }
          ]
        }
      }
    }
  }
}

```

2. Deny access to the website from Russia
```jsonc
{
  "apps": {
    "http": {
      "servers": {
        "myserver": {
          "listen": [":443"],
          "routes": [
            {
              "match": [
                {
                  "host": [
                    "test.example.org"
                  ],
		  "maxmind_geolocation": {
                    "db_path": "/usr/share/GeoIP/GeoLite2-Country.mmdb",
                    "deny_countries": [ "RU" ]
                  }
                }
              ],
              "handle": [
                {
                  "handler": "file_server",
                  "root": "/var/www/html"
                }
              ]
            }
          ]
        }
      }
    }
  }
}

```
