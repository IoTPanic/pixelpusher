# PixelPusher
API for PixelCrasher and the UDPX project. This project involves driving a lot of LEDs in addressable strips stared by @hputzek and @martinberlin. For the firmware to use with this application, consult the [UDP Repo](https://github.com/martinberlin/udpx).

# Installation and Usage

## Source

Running or installing from source requires [golang](https://golang.org/dl/) installed, the GOPATH set, and [dep installed](https://golang.github.io/dep/docs/introduction.html). 

Lots of [tutorials](https://www.google.com/search?newwindow=1&ei=5CFtXYv_As-zsAfMjqboDA&q=setup+go+environment) for setting up the environment are available.

First, dependencies are installed with the command `dep ensure` (or `dep ensure -v` if you want to see what's going on) than PixelPusher can be ran with `go run main.go`.  

## Docker

To create a docker image of this project, simply clone the git into `$GOPATH/src/github.com/IoTPanic/` like any other go project ensure Docker is installed and the Docker daemon is running, than execute `docker build -t pixelpusher .` and Docker will create a image from that, installing dependencies as it goes that will be ready to run.

To run the image, simply execute `docker run -p 8080:8080 --name pusher pixelpusher`

## Docker Compose

Docker compose allows for the entire software stack to be easily deployed. Simply build the docker image as instructed in the Docker section above, than run `docker-compose up`.

### Mosquitto MQTT Service

For a MQTT broker, the mosquitto MQTT service is used, there is a simple `mosquitto.conf` used for configuration, and a `mqttauth` file that's used as a user password list for MQTT. Create a new `mqttauth` file with the command `mosquitto_passwd -c mqttauth USERNAME` after installing the `mosquitto` package. 

# TODO

* Add JSON schema validation - https://github.com/xeipuuv/gojsonschema
* Add HTTP or TCP connection options

# Terminology

* **Channel** - A strip of LEDs or a single LED that is RGB or RGBW
* **Controller** This application is the controller application
* **Fixture** - This terminology was taken out of DMX because it seemed like the perfect word for a device that has LED strips on it.
* **Universe** - A group of devices where their channels will recieve the same data if in the same universe. You are not required to make universes if you want to deal with individual devices. 

![Lighting Architecture](https://raw.githubusercontent.com/IoTPanic/pixelpusher/master/.github/Lighting%20Architecture.png)

# API

This API can be used by front ends, or applications. The architecture was though out so a stream could be created from a live performance application, or a wireless controller, while the frontend such as pixel crasher can continue to live monitor the LED strips.

![System Architecture](https://raw.githubusercontent.com/IoTPanic/pixelpusher/master/.github/Pixels%20System%20Architecture.png)

## Routes

`/api/status`

Responds with if the MQTT is connected, the number of connected devices, the number of registered devices, the current average frame rate, number of open client streams open, and the current configuration.

`/api/healthcheck`

Should always return HTTP code `200` to a GET response.

`/api/fixtures`

Lists all the devices registered, with their current status and channel configuration.

`/api/fixtures/add`

Allows for a `POST` request that contains a valid JSON representation of a device which will be added to storage. 

`/api/fixtures/remove`


`/api/fixtures/modify`

`/api/universes`

`/api/universes/add`

`/api/universes/remove`

`/api/universes/modify`

`/api/universes/fixtures`

## JSON Objects

### Fixture

The Go objects for a fixture are in `internal/pusher/comonents.go` and are `Fixture`, `Connection`, and `Channel` to construct an entire fixture object.

```
{
  "name": "fixture-name",
  "longID": "esp-ID",
  "pixelsID": 1,
  "connection": {
    "IP": "0.0.0.0",
    "port": 1234,
    "type": "UDP",
    "mqtt": true
  },
  "channels": [
    {
      "ID": 0,
      "pixelCnt": 144,
      "RGBW": false
    }
  ],
  "universe": "null"
}
```

### Universe

```

```
