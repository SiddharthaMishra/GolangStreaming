# A websocket based livestream server in Go

### Dependancies:

1. Go: https://golang.org/doc/install


2. Python dependancies (optional, for broadcasting only)

* Use the requirements.txt: ``` pip install -r requirements.txt ``` 


### Building

1. ``` make install ``` to create the binary

2. ``` make clean ``` to, remove the binary

### Running

1. ./server to run the server
2. python testBroadcaster.py to start streaming
3. localhost:3000/camera.html to watch the stream