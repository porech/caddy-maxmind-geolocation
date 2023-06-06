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

You will specifically need the `GeoLite2-Country.mmdb` file, or the `GeoLite2-City.mmdb` if you're matching on subdivisions and metro codes.

## Usage

You can use this module as a matcher to blacklist or whitelist a set of countries, subdivisions or metro codes. 

You'll find the detailed explanation of all the fields on the [Caddy website's plugin page](https://caddyserver.com/docs/modules/http.matchers.maxmind_geolocation).

Here are some samples:

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

2. Deny access to the website from Russia or from IPs with an unknown country:
```
test.example.org {
  @mygeofilter {
    maxmind_geolocation {
      db_path "/usr/share/GeoIP/GeoLite2-Country.mmdb"
      deny_countries RU UNK
    }
  }

   file_server @mygeofilter {
     root /var/www/html
   }
}

```

3. Allow access from US and CA, but exclude the NY subdivision (note that you'll need the City database here):
```
test.example.org {
  @mygeofilter {
    maxmind_geolocation {
      db_path "/usr/share/GeoIP/GeoLite2-City.mmdb"
      allow_countries US CA
      deny_subdivisions NY
    }
  }

   file_server @mygeofilter {
     root /var/www/html
   }
}

```

4. Allow access from US, but only to TX subdivision excluding the metro code 623 and the not-recognized metro codes:
```
test.example.org {
  @mygeofilter {
    maxmind_geolocation {
      db_path "/usr/share/GeoIP/GeoLite2-City.mmdb"
      allow_countries US
      allow_subdivisions TX
      deny_metro_codes 623 UNK
    }
  }

   file_server @mygeofilter {
     root /var/www/html
   }
}

```

5. Deny access from AS64496 (note that you'll need the ASN database here):
```
test.example.org {
  @mygeofilter {
    maxmind_geolocation {
      db_path "/usr/share/GeoIP/GeoLite2-ASN.mmdb"
      deny_asn 64496
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

2. Deny access to the website from Russia or from IPs with an unknown country:
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
                    "deny_countries": [ "RU", "UNK" ]
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

3. Allow access from US and CA, but exclude the NY subdivision (note that you'll need the City database here):
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
                    "db_path": "/usr/share/GeoIP/GeoLite2-City.mmdb",
                    "allow_countries": [ "US", "CA" ],
                    "deny_subdivisions": [ "NY" ]
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

4. Allow access from US, but only to TX subdivision excluding the metro code 623 and the not-recognized metro codes:
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
                    "db_path": "/usr/share/GeoIP/GeoLite2-City.mmdb",
                    "allow_countries": [ "US" ],
                    "allow_subdivisions": [ "TX" ],
                    "deny_metro_codes": [ "623", "UNK" ]
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

5. Deny access from AS64496 (note that you'll need the ASN database here):
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
                    "db_path": "/usr/share/GeoIP/GeoLite2-ASN.mmdb",
                    "deny_asn": [ "64496" ]
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
