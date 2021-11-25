FROM golang:1.17

WORKDIR /go/src/app

COPY go.mod ./
RUN go mod download

COPY *.go ./
#COPY assets/hello.html ./
COPY . ./

RUN go build -o /driver

EXPOSE 80

CMD [ "/driver" ]