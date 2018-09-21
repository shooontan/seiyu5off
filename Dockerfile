FROM golang:1.11

ENV TZ="Asia/Tokyo"

RUN echo 'alias ll="ls -lah"' >> ~/.bashrc

WORKDIR /go/src/github.com/shooontan/seiyu5off

COPY . /go/src/github.com/shooontan/seiyu5off/

RUN \
  go get -u github.com/golang/dep/cmd/dep \
  && dep ensure
