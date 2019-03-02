FROM golang:1.11 as testrunner
LABEL maintainer="Sven Jeppsson <sven@jeppsson.de>"
ENV DBCON root:secret@tcp(mysql:3306)/test
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep
VOLUME /go/src/app
WORKDIR /go/src/app
CMD ./build.sh
