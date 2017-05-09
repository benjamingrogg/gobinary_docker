Run a GO Binary inside a Docker-Container
=========================================

Use FROM scratch to provide a small footprint


Compile
-------

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pong`


Build
-----

`sudo docker build -t benjamingrogg/pong .`

Run
---

`sudo docker run -p 8888:8888 -ti benjamingrogg/pong`

