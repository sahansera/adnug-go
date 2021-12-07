# Build
FROM golang:latest as build

ARG TARGETOS
ARG TARGETARCH

WORKDIR /home
COPY main.go ./
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build main.go

# Run
FROM debian
WORKDIR /home
COPY --from=build /home/main ./main
ENTRYPOINT [ "./main" ]
EXPOSE 8080