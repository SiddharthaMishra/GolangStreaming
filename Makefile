install: 
	-go build -o server src/*.go

clean:
	-rm server