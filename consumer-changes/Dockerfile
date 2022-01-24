FROM golang:alpine

RUN wget -O src/main.zip https://github.com/vilamslep/bx24sync/archive/refs/heads/main.zip \
  && unzip src/main.zip -d src/ && cd src/bx24sync-main/ \
  && go build -o /go/bin/ main/consumer-changes/app.go

WORKDIR /go/bin

ENTRYPOINT [ "./app" ]