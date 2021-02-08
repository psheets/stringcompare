FROM golang:latest AS builder
WORKDIR /stringcompare/
COPY main.go .
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest 
WORKDIR /root/
COPY --from=builder /stringcompare/app .
EXPOSE 8080
CMD ["./app"]