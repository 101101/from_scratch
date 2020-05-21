

# FROM golang:1.9 AS builder
# WORKDIR /go/src/github.com/erikasm/openweathermap_exporter
# ADD . .
# #RUN go get -u github.com/golang/dep/cmd/dep && dep ensure -v
# RUN go get -u github.com/prometheus/client_golang/prometheus
# RUN CGO_ENABLED=0 GOOS=linux go build -o openweathermap_exporter *.go

# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# EXPOSE 9520
# COPY --from=builder /go/src/github.com/erikasm/openweathermap_exporter/openweathermap_exporter openweathermap_exporter

# ENTRYPOINT ["/openweathermap_exporter"]

FROM golang:alpine as builder

RUN apk update && apk add git 
COPY . $GOPATH/src/dsmith73/openweathermap_exporter/
WORKDIR $GOPATH/src/dsmith73/openweathermap_exporter/
RUN go get -d -v


RUN go build -o /go/bin/openweathermap_exporter


FROM alpine
EXPOSE 2112
COPY --from=builder /go/bin/openweathermap_exporter /bin/openweathermap_exporter
ENTRYPOINT ["/bin/openweathermap_exporter"]

# Environment Variables
# OWM_PORT  
# OWM_API_KEY  
# OWM_LOCATION  
# OWM_DURATION  
