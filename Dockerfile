FROM golang:1.14.6-alpine3.12 as builder
RUN mkdir /build 
ADD main.go /build/
WORKDIR /build 
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o hpa .

FROM scratch
COPY --from=builder /build/hpa /app/
WORKDIR /app
CMD ["./hpa"]