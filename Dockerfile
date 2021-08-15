FROM golang:1.16

WORKDIR /go-workspace/src/neoway

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /neoway

EXPOSE 4500

CMD [ "/neoway" ]