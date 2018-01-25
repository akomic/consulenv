# ConsulENV

Alternative to envconsul.

## Building

```
go build
```

## Configuring

use config.example.yml or env variables.

- CONSUL_HTTP_ADDR
- CONSUL_HTTP_TOKEN
- CONSUL_HTTP_AUTH
- CONSUL_HTTP_SSL

## Running

```
eval $(./consulenv -p staging/env/MyApp/ -p -p environment/env/MyApp/staging1)
```
