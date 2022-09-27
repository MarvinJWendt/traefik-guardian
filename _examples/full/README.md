# Full example of using Traefik Guardian

This is an example `docker-compose.yml` that uses Traefik Guardian to protect a service (`whoami`) that is behind the Traefik reverse proxy.

```yaml
version: '3'

services:
  traefik: # Your traefik service
    image: "traefik:latest"
    container_name: "traefik"
    networks:
      - proxy
    command:
      - "--providers.docker=true"
      - "--providers.docker.exposedbydefault=false"
      - "--entrypoints.web.address=:80"
    ports:
      - "80:80"
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock:ro"

  traefik-guardian: # Your traefik-guardian service
    image: marvinjwendt/traefik-guardian:latest
    environment:
      - AUTH_HOST=auth.test.localhost # Replace with your auth host (e.g.: auth.example.com).
      - PASSWORDS=plaintext:test1234|test1337 # Replace with your passwords. See the docs for more info at: https://github.com/MarvinJWendt/traefik-guardian#password-management
    networks:
      - proxy
    labels:
      - traefik.enable=true
      - traefik.docker.network=proxy
      - traefik.http.routers.auth.entrypoints=web
      - traefik.http.routers.auth.rule=Host(`auth.test.localhost`) || Path(`/traefik-guardian-session-share`) # Replace auth.test.localhost with your auth host defined above.
      - traefik.http.middlewares.traefik-guardian.forwardauth.address=http://traefik-guardian/auth

  whoami: # A demo whoami service that is protected with traefik-guaridan
    image: containous/whoami
    container_name: "whoami"
    networks:
      - proxy
    labels:
      - traefik.enable=true
      - traefik.docker.network=proxy
      - traefik.http.routers.whoami.entrypoints=web
      - traefik.http.routers.whoami.rule=Host(`whoami.test.localhost`)
      - traefik.http.routers.whoami.middlewares=traefik-guardian # Add this to services that you want to guard

networks:
  proxy:
```
