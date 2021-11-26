FROM golang:1.17

WORKDIR /go/src/app

COPY ./src/ .

RUN go get -d -v ./... 
RUN go install -v ./...
RUN go build -v ./...

ENTRYPOINT ["producer"]
