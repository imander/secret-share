# Secret Share

Secret share is a basic web application that enables users to securely share secrets with one another. The functionality of this application is similar to [One-Time Secret](https://github.com/onetimesecret/onetimesecret). The difference is that all encryption happens client-side so the secret server never sees any sensitive data.

## Usage

### Docker

The easiest way to run secret share is with docker.

**Build locally:**

```
$ make build
```

**Pull pre-built image:**

```
$ docker pull imander/secret-share
```

**Running Docker:**

```
$ docker run --rm -p 8443:8443 imander/secret-share
```

### TLS

By default the docker image serves a self-signed TLS certificate. The certificate can be changed by mounting a new certificate and key pair into the docker image. By default, the application looks for `key.pem` and `cert.pem` in the `/secret-share/ssl/` directory. The certificate and key path can be changed using environment variables.

```
docker run --rm -p 8443:8443 \
  -v "/path/to/cert.pem:/secret-share/ssl/cert.pem" \
  -v "/path/to/key.pem:/secret-share/ssl/key.pem" \
  imander/secret-share
```

### Environment Variables

|Name             |Default                  |Description                                        |
|-----------------|-------------------------|---------------------------------------------------|
|`ADDR`           |`127.0.0.1`              |The address to bind to (set to `0.0.0.0` in docker)|
|`APP_ROOT`       |`/`                      |The root path where the application is served.     |
|`PORT`           |`http = 8080, https=8443`|The port to listen on.                             |
|`SERVE_TLS`      |`true`                   |Set to "`false`" to listen on HTTP.                |
|`TLS_CERTIFICATE`|`./ssl/cert.pem`         |Path to TLS certificate file.                      |
|`TLS_PRIVATE_KEY`|`./ssl/key.pem`          |Path to TLS private key.                           |

