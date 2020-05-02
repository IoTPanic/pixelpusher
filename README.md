# PixelPusher

**WARNING THIS IS A UNFINISHED PROJECT THAT DOES NOT FUNCTIONALLY WORK, WORK IN PROGRESS.**

PixelPusher is a backend for controlling RGB LED matrix installations which may contain multiple network or wired serial devices (controllers) which is used by the PixelCrasher project. PixelPusher contains a REST like API for querying data, WebRTC for streaming data from a browser, state storage for front-ends, command and control for dot controllers, and a UDP server for streaming PIXELs data to said dot controllers (or DMX hardware using a translation).

## Overview

*PixelPusher uses it's own lighting protocols and architecture instead of the popular and free to use artnet or DMX protocols.* This is because DMX512 has limitations to the protocol which do not make it suitable for massive arrays of RGB LEDs and their architecture gets confusing quickly if you have many LED matrix doing different things or have multiple large matrix on a single device. The first concept in PixelPusher is the matrix itself, a matrix can be a single dimensional array longer than one to represent a LED matrix, a two dimensional array to represent a true LED matrix, and a 1x1 "matrix" to represent a single off LED light. Each of these can be one of RGB, RGBW, or just white. A single white 1x1 matrix is assumed to be a strobe light. If the matrix is a two-dimensional array, the LEDs are addresses as a single dimensional array, with each row increase the LED index with the index increasing from left to right similar to reading.

The next concept to understand is a channel, each channel in PixelPusher has an ID (Which is not the human-readable ID presented to the user) and is assigned one or more matrixes. These channels can then be assigned to devices with one device able to use many channels. The concept of the channel is used in order to allow the state of the matrix to be updated without changing the channel and the frontend can assign their own lighting effects to channels. Multiple matrixes can be on the same channel, even including offsets.

All communication between the frontend pushing data and the PixelPusher backend is encoded using version three of Google's protobuffers, the schemas of which can be found in the /protobufs folder. Data sent from PixelPusher to devices is using the PIXELs open source protocol which also includes the ability to fragment packets for going over a limited network without much overhead, and a control protocol custom to PixelPusher

To allow for multiple lighting setups to be saved in a single pixelpusher instance, there is a basic concept of projects which come with a state concept. 

## Requirements

### Just Run it Already

_Will be posting binaries when there is something useful_

### Development / Contributing

In order to work with pixelpusher for development or to contribute to the project, the following tools will be needed in order to do everything-

* Golang (Developed with 1.13+)
* dep
* Protoc (Proto files can be found in the /protobufs folder)
* protoc-gen_go
* Git of course!

#### Protobufs

To generate fresh .proto files, run the following statement while in the project root.

*Golang* - `protoc --go_out=internal/ protobufs/protobufs.proto`

*JavaScript* - `protoc --js_out=protobufs/js/ protobufs/protobufs.proto`

### Docker

Docker is pretty simple, all you need to do is clone this git repository and you can simply build the Dockerfile that is included.

### Building from Source

## Hardware

For hardware, either a basic ESP32-WROOM module, Dot, or another generic board can be used to stream data, please view the pixels library to see if there is an example for your platform. Although, we encourage you to buy a dot board in order to support this project as well as the pixelcrasher web frontend.

## Resources

Resources in pixelpusher which can be also referred to as "components", they have both a table in the database along with 

### Devices

The device which is similar to a DMX "controller" has many properties which are stored within the database that will change how it operates in the system and are a mutatable resource within the API. Each device represents a piece of hardware, or in some cases connector software which that uses the pixel pusher control protocol as well as the pixels lighting protocol in order to operate. All communication to these devices is handled over UDP.

`id` - The SQL ID that will be used to reference the device numerically.

`long_id` - A long ID for a device.

`name` - Human readable name.

`hostname` - IP or domain to reach the device.

`port` - Port number the device is listening on.

`state` - The current state of the device, this is not stored in the SQL database.

`connector` - Boolean

`key` - Security key.

`use_key` - Boolean if to use the key or to skip integrity checking.

`channel_cnt` - The number of channels available for the device.

### Channels

Channels are device positions to attach lighting, all of which have basic on or off or the ability to drive a matrix of LEDs. Each channel can be assigned one or more matrixes, with multiple matrixes being able to receive the same data or even be slightly offset in accordance to it's setup. Every matrix of a channel must however match the coloring information as it effects how the data is encoded and sent to the devices.

`id` - The SQL ID used to ID channels.

`name` - A name that is human readable and probably automatically created by the frontend software.

`device` - A device SQL ID to reference which device a channel is assigned to.

`type` - A matrixes type determines what type of fixture is connected such as BLKWHT, RGB, RGBW.

`color` - Matrix color information for use by the front end.

`max_length` - The maximum matrix length that can be used on the channel.

### Matrixes

Matrixes are attached to channels and represent a matrix that is either single dimensional single value, single dimensional multi-value, or a two dimensional matrix. No matter the size of the matrix, all data is represented as a single dimensional array.

`id` - The SQL ID for SQL for identifying the matrix.

`device` - Device the matrix is connected to.

`channel` - Channel ID used for IDing the channel the matrix is attached to.

`width` - The width of the matrix, this is the value used for the length.

`height` - Height of the matrix.

`type` - Matrix type such as RGB, RGBW, or BLKWHT.

`coloring` - Coloring information for the front end.

`offset` - Offset from zero for the matrix relative to the channel.

`brightness` - This is a signed value which can be set to change the average brightness for the individual matrix.

### Projects

In order to allow for multiple lighting setups to be saved and used, projects are created with all other resources being a child of the project. There can be as many projects saved at once that is reasonable but it is recommended to only be working on or using one project at a time.

`id` - SQL ID of the project

`name` - Human readable name for the project.

`created` - The datetime when the project was created.

`last_update` - Datetime of the last time state information was written.

`client` - Client name who uses the project.

`active` - Boolean if the project is currently active in the instance.

`frontend_state` - Byte array state information for the frontend.

### Users

A basic user system is established in order to allow for tokens to be created by client frontend software. In the future there may be the ability to establish a basic user role system with permissions.

`id` - The ID for the user

`name` - Nickname for the user that will be displayed in the GUI.

`username` - User's username to login with, this can be the same as the name but it should stay the 

`password` - Hash of the user password for a login, is excluded from any request, on a login this field can be used to submit a plaintext password value.

`created` - The time the user was created

`last_login` - Last time this user has logged in.

## API

The exposed API for pixelpusher is a REST-like API that can be used to query or modify resources outlined in this documentation. This API is exposed on port `9040` which is by default exposed only on the localhost.

### Authentication

In order to authenticate with the client, a 256-bit token is assigned to a client which will also be used to check integrity of lighting packets. This token is valid for thirty days which is obtained through a basic login.

Once a token is obtained by the client through the basic login mechanism. This token should be inserted into the Authorization header of the request following "bearer". Assuming the token is valid, the request will go through as it is meant to be, however if the token is invalid then a `401 Unauthorized` will be returned.

### Paths

Within the path, there is a consistent mention of ID and PID, with the ID being a SQL database ID that resource and the PID being a project ID. The project ID is included to ensure that any client state is at least on the same project state as the server to ensure correctness in queries and mutations.

#### /api/version

*GET*

Returns basic versioning info to ensure that the server is the same or similar version that the client expects.

#### /api/users

*GET*

A GET request will return an array of users in protobuf form that exist in the instance. This will also include user permissions and any other info required for the front end. Any user can query this endpoint regardless of if they are an admin or not.

*POST*

Using a POST request on this path will create a new user with the details provided by the client. Some fields will be ignored such as the ID filed by the server and the user will be returned back to the client with all the details such as the ID. 

#### /api/user/login

*POST*

The POST request made here is a basic login request as specified in [RFC7617](https://tools.ietf.org/html/rfc7617). On a 200 success a bearer token will be returned.

#### /api/user/{ID}

*GET*

Receiving a user protobuf can be obtained here.

*PUT*

Update a user record using a protobuf.

*DELETE*

Delete a user from the pixelpusher instance.

Delete a specified user.

#### /api/state

*GET*

Making a GET request to the state path will allow to get information about the current project and information on metrics.

#### /api/state/devices

*GET*

Returns all currently active devices and their state.

#### /api/devices

*GET*

Get all of the devices that have been registered with the API from all projects.

#### /api/projects

*GET*

Return an array of projects that are registered in the instance.

*POST*

Add a new project to the table, this will return the completed projects protobuf with the database ID and etc.

#### /api/project/{PID}

*GET*

Return a complete protobuf for the project with the given PID.

*PUT*

A PUT request allows for a resource to be updated, any project filed excluding the ID can be updated through the PUT request as long as the client is authenticated.

*DELETE*

Remove a project from an instance.

#### /api/project/{PID}/devices

*GET*

Returns an array of devices that are related to the current project. 

*POST*

Create a new device that starts life attached to a project.

#### /api/project/{PID}/device/{ID}

*GET*

Return a specified device's information.

*PUT*

Update the fields that were present in the protobuf. Some fields will be ignored.

*DELETE*

Delete a device from a project, if this device is attached in multiple projects it is only removed from the one specified.

#### /api/project/{PID}/device/{ID}/channels

*GET*

Return an array of channels that are attached to a device in the scope of the project specified.

Channels queried here can be mutated in the `/api/project/{PID}/channel/{ID}` path with the returned ID.

*POST*

Create a new channel for the device in the specified project.

#### /api/project/{PID}/device/{ID}/matrixes

*GET*

Get matrixes for the given device which will have channel information in the protobufs as an array.

*POST*

Create a new matrix for the device in the specified project.

#### /api/project/{PID}/channels

*GET*

Get the channels for the specified project as a protobuf array.

*POST*

Create a new channel in the project. Channels can only be in a single project at a time.

#### /api/project/{PID}/channel/{ID}

*GET*

Query a single channel by ID that exists within the project specified by the PID.

*PUT*

Mutate the channel record specified by the ID and PID.

*DELETE*

Delete a specified channel from the project.

#### /api/project/{PID}/channel/{ID}/matrixes

*GET*

Query matrixes by channel, returned as a protobuf array.

*POST*

Create a new matrix on the specified channel and project.

#### /api/projects/{PID}/matrixes

*GET*

Return an array of all matrixes that are on the project.

*POST*

Create a new matrix for the project with the provided protobuf.

#### /api/project/{PID}/matrix/{ID}

*GET*

*PUT*

*DELETE*

## Device Communication

In order to communicate with the device, pixelpusher uses UDP to connect, manage, and send application data to devices. UDP connections are created with the devices being the client and pixelpusher being the server.

### Joining



### Discovery

Discovery is handled by mDNS with the network's local pixelpusher instance being `pixelpusher.local` by default which clients can connect to. 

### Heartbeat

Every ten seconds a heartbeat must be sent through the UDP connection from the device to pixelpusher, missing two heartbeat messages at a time then it is assumed to be disconnected.

### PIXELS Application Data

## WebRTC

WebRTC will be used to stream lighting data from the frontend to channels, with a REST API used to query and modify resources such as devices, matrixes, and channels along with any resource updates. WebRTC is supposed both by Google Chrome and Mozilla Firefox. 

## Database

The database is a basic SQLite database file which is used for it's portability, with each resource having it's own table. This database file can be used which can be moved to another instance or modified through a normal SQLite tool.

### Tables

Below each table in the SQLite database are described, the internal database library must reflect this README.

#### Projects 

#### Devices

#### Channels

#### Matrixes