# A lightweight Dockerised web frontend for FanFicFare

The site demo: [fanficfare.eleanor.servicies](https://fanficfare.eleanor.services) - a proper domain is coming soon, promise!

Deployment example with Traefik:

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
      - traefik.http.routers.fanficfare.rule=Host(`fanficfare.eleanor.services`)
      - traefik.http.routers.fanficfare-sec.rule=Host(`fanficfare.eleanor.services`)
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


## Readme From JimmXinu/FanFicFare
Go see their Original Repo: [FanFicFare](https://github.com/JimmXinu/FanFicFare)
==========

FanFicFare makes reading stories from various websites much easier by helping
you download them to EBook files.

FanFicFare was previously known as FanFictionDownLoader (AKA
FFDL, AKA fanficdownloader).

Main features:

- Download FanFiction stories from over [100 different sites](https://github.com/JimmXinu/FanFicFare/wiki/SupportedSites). into ebooks.

- Update previously downloaded EPUB format ebooks, downloading only new chapters.

- Get Story URLs from Web Pages.

- Support for downloading images in the story text. (EPUB and HTML
  only -- download EPUB and convert to AZW3 for Kindle) More details on
  configuring images in stories and cover images can be found in the
  [FAQs] or [this post in the old FFDL thread].

- Support for cover image. (EPUB only)

- Optionally keep an Update Log of past updates (EPUB only).

There's additional info in the project [wiki] pages.

There's also a [FanFicFare maillist] for discussion and announcements and a [discussion thread] for the Calibre plugin.

Getting FanFicFare
==================

### Official Releases

This program is available as:

- A Calibre plugin from within Calibre or directly from the plugin [discussion thread], or;
- A Command Line Interface (CLI) [Python
  package](https://pypi.python.org/pypi/FanFicFare) that you can
  install with:
```
pip install FanFicFare
```
- _As of late November 2019, the web service version is shutdown.  See the [Wiki Home](https://github.com/JimmXinu/FanFicFare/wiki#web-service-version) page for details._

### Other Releases

Other versions may be available depending on your OS.  I(JimmXinu) don't directly support these:

- **Arch Linux**: The CLI can also be obtained on Arch Linux from the OS repositories:

```
pacman -S fanficfare
```

...or from git via the [AUR package](https://aur.archlinux.org/packages/fanficfare-git)
(which will also update the calibre plugin, if calibre is installed).



[this post in the old FFDL thread]: https://www.mobileread.com/forums/showthread.php?p=1982785#post1982785
[FAQs]: https://github.com/JimmXinu/FanFicFare/wiki/FAQs#can-fanficfare-download-a-story-containing-images
[FanFicFare maillist]: https://groups.google.com/group/fanfic-downloader
[wiki]: https://github.com/JimmXinu/FanFicFare/wiki
[discussion thread]: https://www.mobileread.com/forums/showthread.php?t=259221
