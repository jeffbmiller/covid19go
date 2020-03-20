FROM golang:1.10

WORKDIR /
COPY . .
RUN go get -d

CMD ["go","run","main.go"]