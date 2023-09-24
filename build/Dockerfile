FROM golang:1.21 as builder
ENV APP_PATH=/go/src/app/
WORKDIR ${APP_PATH}
COPY . ${APP_PATH}
WORKDIR ${APP_PATH}/cmd/
RUN --mount=type=cache,target=/root/.cache CGO_ENABLED=0 go build -mod vendor -ldflags '-w -s' -o main -v .

FROM alpine:3.12
ENV APP_PATH=/go/src/app
WORKDIR /root/
COPY --from=0 ${APP_PATH}/cmd/main ./
EXPOSE 8080
CMD ["./main"]
