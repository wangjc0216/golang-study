FROM alpine:3.10

ENV APP_PATH="/app/printwebhook"
WORKDIR "/app"


# 拷贝程序，如有必要另外拷贝其他文件
COPY  printwebhook  ${APP_PATH}
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# 运行程序
ENTRYPOINT ["/app/printwebhook"]
