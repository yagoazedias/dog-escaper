FROM balenalib/raspberry-pi-debian-golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build main.go

EXPOSE 8000

CMD ["./main"]