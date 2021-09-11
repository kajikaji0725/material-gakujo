FROM golang:1.17 AS build

WORKDIR /work

COPY . .

RUN \
    CGO_ENABLED=0 \
    GOOS=linux \
    go build -mod=vendor -o app ./api

FROM ubuntu:20.04

WORKDIR /work

#RUN apk --no-cache add tzdata && \
#    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
#    apk del tzdata

RUN \
    apt-get update && \
    apt install -y tzdata && \
    apt-get clean && rm -rf /var/lib/apt/lists/*
    
ENV TZ Asia/Tokyo

RUN \
    apt-get update && \
    apt install -y ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=build /work/app app

ENTRYPOINT [ "./app" ]
