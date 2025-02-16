## Authentication Stack
Replace all authentication steps with a COTS solution. With that said, we should replace out in house authentication solution with a commerical off the shelve solution such as OAuth2 Proxy using Redis and possibly Keycloak to provide authentication/authorization for our application. This will reduce development time and allow for ease of use for authentication. May benefits to make our application work with this stack over making our own and quite possibly creating multiple layers of vulnerabilities.

## Conerns
- There might be an issue with Redis going closed source. We want to replace Redis with another in memory database such as Valkey (evaluation might be needed)

## Documentation
- https://github.com/oauth2-proxy/oauth2-proxy
- https://redis.io/
- https://github.com/valkey-io/valkey
- https://www.keycloak.org/