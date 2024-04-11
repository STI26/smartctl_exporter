##
## Build bin
##
FROM golang:1.22-alpine AS build

WORKDIR /app

ENV VERSION $VERSION
ENV BUILD_DATE $BUILD_DATE

COPY . ./
RUN go mod download
RUN go mod verify

RUN go build -o /go/bin/smartctl_exporter -ldflags "-w -s -X main.Version=${VERSION} -X main.BuildDate=${BUILD_DATE}"

##
## Build finish image
##
FROM alpine:latest

RUN apk add --no-cache smartmontools

WORKDIR /opt/app/

COPY --from=build /go/bin/smartctl_exporter smartctl_exporter

EXPOSE 9111

ENTRYPOINT ["/opt/app/smartctl_exporter"]
