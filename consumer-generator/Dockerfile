FROM golang:1.17

COPY ./src /opt/app

WORKDIR /opt/app

RUN go get -d -v
RUN go install -v ./...

CMD ["app"]