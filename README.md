# PixelPusher
API for PixelCrasher and the UDPX project

# Routes

`/status`

Responds with if the MQTT is connected, the number of connected devices, the number of registered devices, the current average frame rate, number of open client streams open, and the current configuration.

`/health heck`

Should always return HTTP code `200` to a GET response.

`/devices`

Lists all the devices registered, with their current status and channel configuration.

`/devices/add`

Allows for a `POST` request that contains a valid JSON representation of a device which will be added to storage. 

`/devices/remove`



`/devices/modify`
