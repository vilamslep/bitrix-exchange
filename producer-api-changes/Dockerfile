FROM golang:1.17

COPY ./src  /opt/app

WORKDIR /opt/app

ENV HTTP_HOST=producer-api
ENV HTTP_PORT=15000
ENV HTTP_CHECK_INPUT=true
ENV KAFKA_HOST=kafka
ENV KAFKA_PORT=9092
ENV KAFKA_TOPIC=changes

EXPOSE 15000

RUN go get -d -v 
RUN go install -v ./...
RUN go build app.go
ENTRYPOINT [ "./app" ]



 