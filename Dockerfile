FROM golang:1.14-alpine

WORKDIR /fanficfare

RUN apk add --update py-pip git

RUN pip3 install FanFicFare

RUN go get -d -v gopkg.in/labstack/echo.v4 gopkg.in/labstack/echo.v4/middleware github.com/microcosm-cc/bluemonday

COPY . /fanficfare

RUN go build server.go

EXPOSE 80

CMD ["./server"]f