FROM golang:1.15.2

RUN mkdir $GOPATH/app
WORKDIR $GOPATH/app

COPY . .

CMD go run .