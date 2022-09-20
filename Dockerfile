FROM golang:alpine as builder
WORKDIR /build
COPY ./go.mod ./
COPY ./go.sum ./
COPY ./src ./src
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o ./main ./src/cmd
RUN apk add upx
RUN upx --best --lzma main

FROM node:alpine as ui-builder
WORKDIR /build
COPY ./src/admin-ui ./
RUN npm install
RUN npm run build

FROM scratch
WORKDIR /app
COPY --from=builder /build/main ./main
COPY --from=ui-builder /build/dist /app/admin-ui
ADD src/internal/html /app/html
EXPOSE 80
ENTRYPOINT ["./main"]
