FROM golang:1.10

WORKDIR /
COPY main.go .
RUN go get -d github.com/gorilla/mux \
    && go get -d github.com/mitchellh/mapstructure \
    && go get -d github.com/PuerkitoBio/goquery


CMD ["go","run","main.go"]