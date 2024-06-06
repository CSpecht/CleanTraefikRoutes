# Clean Traefik Routes


# Description

A simple api wrapper for the static file traefik configuration, to dynamically add http routes where a wildcard certificate is impossible due to missing dns challenge, but only one target service is needed

# Run

Get the latest container from docker pull cspecht/clean-traefik-routes:latest, setup with with following environment variables:

- FILENAME default: "8080"
- INTERVAL default: "* * * * *"
- PERIOD default: "MONTHLY"
- TICKS default: "[3,9]"
- DB_HOST default: "localhost"
- DB_PORT default: "5432"
- DB_USERNAME default: "postgres"
- DB_NAME default: "postgres"
- DB_PW default: "postgres"
- DB_TABLE default: "principles" 

# License

- under [MIT license](http://opensource.org/licenses/mit-license.php)

