FROM golang:1.11

ENV TZ="Asia/Tokyo"

RUN echo 'alias ll="ls -lah"' >> ~/.bashrc

WORKDIR /go/src/seiyu5off

COPY . /go/src/seiyu5off

RUN \
  go get \
    github.com/PuerkitoBio/goquery
