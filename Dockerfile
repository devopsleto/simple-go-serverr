FROM golang:1.15.2

RUN mkdir $GOPATH/app
RUN echo a && echo $(cat /etc/*-release)
WORKDIR $GOPATH/app

COPY . .

CMD go run .