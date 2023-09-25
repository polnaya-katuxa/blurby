FROM docker:24.0.2-dind-alpine3.18

RUN apk add go
RUN go version
RUN echo $DOCKER_HOST

RUN mkdir /test && chmod 777 /test
WORKDIR /test
