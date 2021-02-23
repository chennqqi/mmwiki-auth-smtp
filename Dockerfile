FROM golang:1.14.7-alpine3.12
RUN go env -w GOPROXY=https://goproxy.cn,direct
COPY . /go/src/auth
WORKDIR /go/src/auth
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /bin/auth

FROM alpine:3.12
WORKDIR /app
COPY --from=0 /bin/auth app

RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.12/main" > /etc/apk/repositories && \
	apk add --no-cache -U tzdata && \
	addgroup -S app && \
	adduser app -S -G app -h /app && \
	chown -R app:app /app

EXPOSE 8080
ENV TZ=Asia/Shanghai
ENTRYPOINT [ "/app" ]
