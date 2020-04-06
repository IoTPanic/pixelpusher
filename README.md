# PixelPusher

PixelPusher is a backend for controlling RGB LED matrix installations which may contain multiple network or wired serial devices which is used by the PixelCrasher project. PixelPusher contains a REST like API for querying data, WebRTC for streaming data from a browser, state storage for frontends, command and control for dot controllers, and a UDP server for streaming PIXELs data to said dot controllers.

## Overview

PixelPusher uses it's own lighting protocols and architecture instead of the popular and free to use artnet or DMX protocols. This is because DMX512 has limitations to the protocol which do not make it suitable for massive arrays of RGB LEDs and their architecture gets confusing quickly if you have many LED matrix doing different things or have multiple large matrix on a single device. The first concept in PixelPusher is the matrix itself, a matrix can be a single dimensional array longer than one to represent a LED strip, a two dimensional array to represent a true LED matrix, and a 1x1 "matrix" to represent a single off LED light. Each of these can be one of RGB, RGBW, or jjust white. A single white 1x1 matrix is assumed to be a strobe light.

The next concept to understand is a channel, each channel in PixelPusher has an ID (Which is not the human-readable ID presented to the user) and is assigned one or more strips. These channels can then be assigned to devices with one device able to use many channels. The concept of the channel is used in order to allow the state of the strip to be updated without changing the channel and the frontend can assign their own lighting effects to channels. Multiple strips can be on the same channel.

## Requirements

### Just Run it Already

_Will be posting binaries when there is something useful_

### Development / Contributing

In order to work with pixelpusher for development or to contribute to the project, the following tools will be needed in order to do everything-

* Golang (Developed with 1.13+)
* dep
* Protoc
* Git of course!

### Docker

### Building from Source

## API

## WebRTC

## Database

## Connecting Devices and LEDs
