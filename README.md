# PixelPusher

PixelPusher is a backend for controlling RGB LED matrix installations which may contain multiple network or wired serial devices (controllers) which is used by the PixelCrasher project. PixelPusher contains a REST like API for querying data, WebRTC for streaming data from a browser, state storage for front-ends, command and control for dot controllers, and a UDP server for streaming PIXELs data to said dot controllers (or DMX hardware using a translation).

## Overview

*PixelPusher uses it's own lighting protocols and architecture instead of the popular and free to use artnet or DMX protocols.* This is because DMX512 has limitations to the protocol which do not make it suitable for massive arrays of RGB LEDs and their architecture gets confusing quickly if you have many LED matrix doing different things or have multiple large matrix on a single device. The first concept in PixelPusher is the matrix itself, a matrix can be a single dimensional array longer than one to represent a LED strip, a two dimensional array to represent a true LED matrix, and a 1x1 "matrix" to represent a single off LED light. Each of these can be one of RGB, RGBW, or just white. A single white 1x1 matrix is assumed to be a strobe light. If the matrix is a two-dimensional array, the LEDs are addresses as a single dimensional array, with each row increase the LED index with the index increasing from left to right similar to reading.

The next concept to understand is a channel, each channel in PixelPusher has an ID (Which is not the human-readable ID presented to the user) and is assigned one or more strips. These channels can then be assigned to devices with one device able to use many channels. The concept of the channel is used in order to allow the state of the strip to be updated without changing the channel and the frontend can assign their own lighting effects to channels. Multiple strips can be on the same channel.

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

Docker is pretty simple, all you need to do is clone this git repository and you can simply build the Dockerfile that is included.

### Building from Source

## Resources

Resources in pixelpusher which can be also referred to as "components", they have both a table in the database along with 

### Devices

The device which is similar to a DMX "controller" has many properties which are stored within the database that will change how it operates in the system and are a mutatable resource within the API. Each device represents a piece of hardware, or in some cases connector software which that uses the pixel pusher control protocol as well as the pixels lighting protocol in order to operate. All communication to these devices is handled over UDP.

`id` - The SQL ID that will be used to refrence the device numerically.

`long_id` - A long ID for a device.

`name` - Human readbale name.

`hostname` - IP or domain to reach the device.

`port` - Port number the device is listening on.

`state` - The current state of the device, this is not stored in the SQL database.

`connector` - Boolean

`key` - Security key.

`use_key` - Boolean if to use the key or to skip integrity checking.

`channel_cnt` - The number of channels available for the device.

### Channels

Channels are device positions to attach lighting, all of which have basic on or off or the ability to drive a strip of LEDs. Each channel can be assigned one or more strips, with multiple strips being able to receive the same data or even be slightly offset in accordance to it's setup. Every strip of a channel must however match the coloring information as it effects how the data is encoded and sent to the devices.

`id` - The SQL ID used to ID channels.

`device` - A device SQL ID to refrence which device a channel is assigned to.

`type` - A strips type detirmines what type of fixture is connected such as BLKWHT, RGB, RGBW.

`color` - String color information for use by the front end.

`max_length` - The maximum matrix length that can be used on the channel.

### Strips

Strips are attached to channels and represent a matrix that is either single dimensional single value, single dimensional multi-value, or a two dimensional matrix. No matter the size of the matrix, all data is represented as a single dimensional array.

`id` - The SQL ID for SQL for identifying stips.

`device` - Device the strip is conencted to.

`channel` - Channel ID used for IDing the channel the strip is attached to.

`width` - The width of the matrix, this is the value used for the legnth.

`height` - Height of the matrix.

`type` - Matrix type such as RGB, RGBW, or BLKWHT.

`coloring` - Coloring information for the front end.

`offset` - Offset from zero for the strip relative to the channel.

`brightness` - This is a signed value which can be set to change the average brightness for the individual strip.

## Hardware

For hardware, either a basic ESP32-WROOM module, Dot, or another generic board can be used to stream data, please view the pixels library to see if there is an example for your platform. Although, we encourage you to buy a dot board in order to support this project as well as the pixelcrasher web frontend.

## API

The exposed API for pixelpusher is a REST-like API that can be used to query or modify resources outlined in this documentation. This API is exposed on port `9040` which is by default exposed only on the localhost.

### Authentication

In order to authenticate with the client, a 256-bit token is assigned to a client which will also be used to check integrity of lighting packets. This token is valid for thirty days which is obtained through a basic login.

### Paths

`/api/devices` - Get an array of devices.

`/api/devices/{ID}` - Get a specific device by ID which can also be mutated.

`/api/devices/{ID}/channels` - Return an array of channels that are attached to the specific device.

`/api/devices/{ID}/strips` - Returns the strips attached to a device.

`/api/channels` - Get an array of channels available from the database.

`/api/channels/{ID}` - Return a channel from an ID which is mutable.

`/api/channels/{ID}/strips` - Get all the strips attached to a channel.

## WebRTC

WebRTC will be used to stream lighting data from the frontend to channels, with a REST API used to query and modify resources such as devices, strips, and channels along with any resource updates. WebRTC is supposed both by Google Chrome and Mozilla Firefox. 

## Database

The database is a basic SQLite database file which is used for it's portability, with each resource having it's own table. This database file can be used which can be moved to another instance or modified through a normal SQLite tool.
