# PixelPusher
API for PixelCrasher and the UDPX project

# Installation and Usage

## Source

Running or installing from source requires golang installed, the GOPATH set, and [dep installed](https://golang.github.io/dep/docs/introduction.html). First, dependencies are installed with the command `dep ensure` (or `dep ensure -v` if you want to see what's going on) than PixelPusher can be ran with `go run main.go`.  


## Docker

To create a docker image of this project, simply clone the git into `$GOPATH/src/github.com/IoTPanic/` like any other go project ensure Docker is installed and the Docker daemon is running, than execute `docker build -t pixelpusher .` and Docker will create a image from that, installing dependencies as it goes that will be ready to run.

To run the image, simply execute `docker run -p 8080:8080 --name pusher pixelpusher`

## Docker Compose

Docker compose allows for the entire software stack to be easily deployed. Simply build the docker image as instructed in the Docker section above, than run `docker-compose up`.

# TODO

* Add JSON schema validation - https://github.com/xeipuuv/gojsonschema
* Add HTTP or TCP connection options

# Routes

`/api/status`

Responds with if the MQTT is connected, the number of connected devices, the number of registered devices, the current average frame rate, number of open client streams open, and the current configuration.

`/api/healthcheck`

Should always return HTTP code `200` to a GET response.

`/api/devices`

Lists all the devices registered, with their current status and channel configuration.

`/api/devices/add`

Allows for a `POST` request that contains a valid JSON representation of a device which will be added to storage. 

`/api/devices/remove`



`/api/devices/modify`
