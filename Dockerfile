FROM golang:1.22.5 AS build

WORKDIR /go/src/finance

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/finance ./cmd/finance/main.go

FROM gcr.io/distroless/static-debian11:nonroot

COPY --from=build /go/bin/finance /

EXPOSE 3000

CMD ["/finance"]