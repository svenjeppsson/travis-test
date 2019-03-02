FROM golang:1.11 as testrunner
LABEL maintainer="Sven Jeppsson <sven@jeppsson.de>"
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep
VOLUME /go/src/app
WORKDIR /go/src/app
CMD source ./build.sh
