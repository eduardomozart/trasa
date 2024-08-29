FROM golang:1.14 as gobuilder

WORKDIR /go/src/seknox/trasa

COPY ./../../../ /go/src/seknox/trasa

#COPY go.mod .
#COPY go.sum .

RUN go mod download

#COPY server server
WORKDIR /go/src/seknox/trasa/server

RUN go build


FROM node:14.17.0-alpine as dashbuilder

WORKDIR /trasa
ENV PATH /trasa/node_modules/.bin:$PATH

# install app dependencies
#COPY dashboard/package.json ./
#RUN yarn install --silent
COPY --from=gobuilder /go/src/seknox/trasa/dashboard ./
RUN npm install --legacy-peer-deps

#COPY dashboard ./


#RUN yarn run build
RUN export NODE_OPTIONS=--openssl-legacy-provider
RUN npm run build

FROM debian:buster-slim
USER root
WORKDIR /trasa
ENV GUACENC_INSTALLED=true
COPY --from=seknox/guacd:v0.0.1 /usr/local/guacamole/bin/guacenc /usr/local/guacamole/bin/guacenc
COPY --from=seknox/guacd:v0.0.1 /usr/local/guacamole/lib/ /usr/local/guacamole/lib/
COPY --from=seknox/guacd:v0.0.1 /usr/local/guacamole/DEPENDENCIES /usr/local/guacamole/

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    apt-get install -y  libavcodec-dev libavformat-dev libavutil-dev libswscale-dev ffmpeg && \
    apt-get install -y  $(cat /usr/local/guacamole/DEPENDENCIES) && \
    rm -rf /var/lib/apt/lists/*

RUN update-ca-certificates
COPY --from=gobuilder /go/src/seknox/trasa/server/server .
COPY --from=dashbuilder /trasa/build /var/trasa/dashboard
COPY --from=gobuilder /go/src/seknox/trasa/build/etc/trasa /etc/trasa
CMD ["/trasa/server"]
