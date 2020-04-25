# PixelPusher

PixelPusher is a backend for controlling RGB LED matrix installations which may contain multiple network or wired serial devices which is used by the PixelCrasher project. PixelPusher contains a REST like API for querying data, WebRTC for streaming data from a browser, state storage for frontends, command and control for dot controllers, and a UDP server for streaming PIXELs data to said dot controllers (or DMX hardware using a translation).

## Overview

*PixelPusher uses it's own lighting protocols and architecture instead of the popular and free to use artnet or DMX protocols.* This is because DMX512 has limitations to the protocol which do not make it suitable for massive arrays of RGB LEDs and their architecture gets confusing quickly if you have many LED matrix doing different things or have multiple large matrix on a single device. The first concept in PixelPusher is the matrix itself, a matrix can be a single dimensional array longer than one to represent a LED matrix, a two dimensional array to represent a true LED matrix, and a 1x1 "matrix" to represent a single off LED light. Each of these can be one of RGB, RGBW, or jjust white. A single white 1x1 matrix is assumed to be a strobe light.

The next concept to understand is a channel, each channel in PixelPusher has an ID (Which is not the human-readable ID presented to the user) and is assigned one or more matrixes. These channels can then be assigned to devices with one device able to use many channels. The concept of the channel is used in order to allow the state of the matrix to be updated without changing the channel and the frontend can assign their own lighting effects to channels. Multiple matrixes can be on the same channel.

All communication between the frontend pushing data and the PixelPusher backend is encoded using version three of Google's protobuffers, the schemas of which can be found in the /protobufs folder. Data sent from PixelPusher to devices is using the PIXELs open source protocol which also includes the ability to fragment packets for going over a limited network without much overhead, and a control protocol custom to PixelPusher

## Requirements

### Just Run it Already

_Will be posting binaries when there is something useful_

### Development / Contributing

In order to work with pixelpusher for development or to contribute to the project, the following tools will be needed in order to do everything-

* Golang (Developed with 1.13+)
* dep
* Protoc (Proto files can be found in the /protobufs folder)
* Git of course!

### Docker

### Building from Source

## Hardware

## API

#### /api/users

**GET**

A GET request will return an array of users in protobuf form that exist in the instance. This will also include user permissions and any other info required for the front end. Any user can query this endpoint regardless of if they are an admin or not.

**POST**

Using a POST request on this path will create a new user with the details provided by the client. Some fields will be ignored such as the ID filed by the server and the user will be returned back to the client with all the details such as the ID. 

#### /api/user/{ID}

**GET**

**POST**

**PUT**

**DELETE**

#### /api/projects

**GET**

**POST**

#### /api/project/{PID}

**GET**

**POST**

**PUT**

**DELETE**

#### /api/project/{PID}/devices

**GET**

Returns an array of devices that are related to the current project. 

**POST**

#### /api/project/{PID}/device/{ID}

**GET**

**POST**

**PUT**

**DELETE**

#### /api/project/{PID}/device/{ID}/channels

**GET**

**POST**

#### /api/project/{PID}/device/{ID}/matrixes

**GET**

**POST**

#### /api/project/{PID}/channels

**GET**

**POST**

#### /api/project/{PID}/channel/{ID}

**GET**

**POST**

**PUT**

**DELETE**

#### /api/project/{PID}/channel/{ID}/matrixes

**GET**

**POST**

#### /api/projects/{PID}/matrixes

**GET**

**POST**

#### /api/project/{PID}/matrix/{ID}

**GET**

**POST**

**PUT**

**DELETE**

## WebRTC

## Database
