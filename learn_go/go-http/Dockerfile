# stage 1 of multi-stage build
FROM golang:alpine AS builder
# can we just say FROM golang  ?  

WORKDIR /go/src/app
COPY main.go .

RUN go build -o webserver .



# stage 2
FROM alpine
# could reduce size more by building FROM scratch, but then you'd need to change the compile method in stage 1  
WORKDIR /app
#EXPOSE 8080
COPY --from=builder /go/src/app/ /app/
ARG VCS_REF

LABEL org.label-schema.vcs-ref=$VCS_REF
#  \
        # org.label-schema.vcs-url="e.g. https://github.com/microscaling/microscaling"

CMD ["./webserver"]