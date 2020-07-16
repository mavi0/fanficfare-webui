# A lightweight Dockerised web frontend for FanFicFare

Written in go! Uses the most recent release of FanFicFare at time of building.

The site demo: [fanficfare.eleanor.servicies](https://fanficfare.eleanor.services) - a proper domain is coming soon, promise 100%

Docker hub: https://hub.docker.com/r/mavi0/fanficfare

Docker run example: `docker run -p 80:80 mavi0/fanficfare:latest`

Docker-compose deployment example with Traefik:

```yaml
version: "3.7"

services:
  fanficfare:
    image: mavi0/fanficfare:latest
    container_name: fanficfare
    networks:
      - traefik-network
      - default
    environment:
      - PUID=${PUID}
      - PGID=${PGID}
      - TZ=${TZ}
    restart: unless-stopped
    labels:
      - traefik.enable=true
      - traefik.http.routers.fanficfare.entrypoints=web
      - traefik.http.routers.fanficfare-sec.entrypoints=websecure
      - traefik.http.routers.fanficfare.rule=Host(`fanficfare.${DOMAIN}`)
      - traefik.http.routers.fanficfare-sec.rule=Host(`fanficfare.${DOMAIN})
      - traefik.http.services.fanficfare-sec.loadbalancer.server.port=80
      - traefik.http.routers.fanficfare.middlewares=basic-http
      - traefik.http.routers.fanficfare-sec.middlewares=basic
      - traefik.http.routers.fanficfare-sec.tls=true
      - traefik.http.routers.fanficfare-sec.tls.certresolver=cfdns

networks:
  traefik-network:
    external: true
  default:
    driver: bridge
```



Uses FanFicFare - see their Original Repo: [FanFicFare](https://github.com/JimmXinu/FanFicFare)
