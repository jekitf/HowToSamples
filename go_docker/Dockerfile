FROM golang:1.14.3-alpine AS buildContainer
WORKDIR /compile
COPY . .
RUN go build -o /out/go_service

FROM alpine:3.7 AS runtimeContainer
COPY --from=buildContainer /out/ /app
EXPOSE 8080
#ENTRYPOINT ["/bin/sh"]
ENTRYPOINT ["/app/go_service"]