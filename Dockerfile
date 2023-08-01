FROM golang:1.20.5-alpine3.18

WORKDIR /app
COPY . ./
RUN go mod download

COPY cmd/app/main.go ./

RUN go build -o app ./cmd/app/main.go

EXPOSE ${HTTP_PORT}

CMD [ "./app" ]
