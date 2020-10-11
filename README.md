# caddy-maxmind-geolocation

Caddy v2 module to filter requests based on source IP geographic location. This was a feature provided by the V1 `ipfilter`
middleware. 

## Installation

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

- Allow access to the website only from Italy and France:
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

- Deny access to the website from Russia:
```
test.example.org {
  @mygeofilter {
    maxmind_geolocation {
      db_path "/usr/share/GeoIP/GeoLite2-Country.mmdb"
      deny_countries IT FR
    }
  }

   file_server @mygeofilter {
     root /var/www/html
   }
}

```

### API/JSON

- Allow access to the website only from Italy and France:
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

- Deny access to the website from Russia
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

2. As an handler within a route for commands that get triggered by an http endpoint.

```jsonc

{
  ...
  "routes": [
    {
      "handle": [
        // exec configuration for an endpoint route
        {
          // required to inform caddy the handler is `exec`
          "handler": "exec",
          // command to execute
          "command": "git",
          // command arguments
          "args": ["pull", "origin", "master"],

          // [optional] directory to run the command from. Default is the current directory.
          "directory": "/home/user/site/public",
          // [optional] if the command should run on the foreground. Default is false.
          "foreground": true,
          // [optional] timeout to terminate the command's process. Default is 10s.
          "timeout": "5s"
        }
      ],
      "match": [
        {
          "path": ["/generate"]
        }
      ]
    }
  ]
}
```
