FROM golang:1.14

WORKDIR /go/src/

COPY . .

RUN go mod download

RUN go build -o /go/src/app/application


CMD ["/go/src/app/application"]