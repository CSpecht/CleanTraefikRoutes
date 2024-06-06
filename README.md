# Clean Traefik Routes


# Description

A simple api wrapper for the static file traefik configuration, to dynamically add http routes where a wildcard certificate is impossible due to missing dns challenge, but only one target service is needed

# Run

Get the latest container from docker pull cspecht/traefik-static-file-api:latest, setup with with following environment variables:

API_PORT default: "8080"
FILENAME default: "services.yml"
ROUTER default: "http"
DOMAIN default: ".example.com"
SERVICE default: "svc@docker"
CERT_RESOLVER default: "le"
ENTRY_POINT default: "websecure"
Simple execute a post command to localhost:8080as a json including 'subdomain' key, e.g.

{"subdomain":"john"}

Hints

If using traefik inside a docker, always mount a folder including the config files into the container when using the watch functionality, otherwise no changes are detected
There is no security mechanism for this service, so dont expose this service!
# License

- under [MIT license](http://opensource.org/licenses/mit-license.php)

