# ConsulENV

Alternative to envconsul.

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
eval $(./consulenv -p staging/env/MyApp/ -p environment/env/MyApp/staging1)
```
