# ConsulENV

Pulls key/value pairs from Consul KV store and returns it to stdout.

Alternative to envconsul.

- It doesn't keep connection to Consul opened.
- Env variable change is not triggering service restart, that's handled by rolling update to secure availability.
- Env variables are loaded in order specified by prefixes, interpolation supported.

## Building

```
go build
```

## Configuring

use config.example.yml or env variables.

- CONSUL_HTTP_ADDR (localhost:1234)
- CONSUL_HTTP_TOKEN
- CONSUL_HTTP_AUTH (user:pass)
- CONSUL_HTTP_SSL (true|false)

## Running

```
eval "$(./consulenv -p staging/env/ -p staging/MyApp/env/)"
```
