FROM golang:1.19-alpine as dev

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . /app/

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11 as prod

COPY --from=dev go/bin/app /

EXPOSE 5000

CMD ["/app"]