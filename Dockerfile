FROM golang:alpine as builder
WORKDIR /build
COPY ./go.mod ./
COPY ./go.sum ./
COPY ./src ./src
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./main ./src/cmd

FROM scratch
WORKDIR /app
COPY --from=builder /build/main ./main
ADD src/internal/html /app/html
EXPOSE 80
ENTRYPOINT ["./main"]
