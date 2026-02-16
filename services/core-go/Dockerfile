FROM golang:1.22-alpine AS build
WORKDIR /src
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -o /out/app
FROM alpine:3.19
WORKDIR /app
COPY --from=build /out/app /app/app
EXPOSE 8080
ENTRYPOINT ["/app/app"]
