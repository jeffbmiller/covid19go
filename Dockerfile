FROM golang:1.10

WORKDIR /
COPY . .
RUN go get -d

ENTRYPOINT ["go","run","main.go"]