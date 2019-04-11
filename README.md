# A websocket based livestream server in Go

### Dependancies:

1. Go: https://golang.org/doc/install


2. Python dependancies (optional, for broadcasting only)

* Required python version >= 3.6
* Use the requirements.txt: ``` pip install -r requirements.txt ``` 


### Building

1. ``` make install ``` to create the server binary

2. ``` make clean ``` to, remove the server binary

### Running

1. ./server to run the server
2. ``` python testBroadcaster.py ``` to start streaming (optional), else connect a websocket to ws://localhost:3000/broadcaster and send base64 encoded jpg images
3. from a browser, go to localhost:3000/camera.html to watch the stream