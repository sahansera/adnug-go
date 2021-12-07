# Commands used for this tutorial

For this tutorial, we will be using our example in `11-http-service-json`

What's currently available?
```sh
go tool dist list
```

Targetting macOS and amd64
```sh
GOOS=darwin GOARCH=amd64 go build main.go
```

How about Windows?
```sh
GOOS=windows GOARCH=amd64 go build main.go
```

Building the Docker image
```sh
docker build -t sahan/hello-go:v1 -f Dockerfile .

# or you could target a specific OS and ARCH
docker build --build-arg TARGETOS=linux --build-arg TARGETARCH=amd64 -t sahan/hello-go:v1 -f Dockerfile .
```

Running the container
```sh
docker run --rm --publish 8080:8080 sahan/hello-go:v1
```