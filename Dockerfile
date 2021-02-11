FROM alpine

WORKDIR /app

# 设置时区为上海
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

COPY my_assistant_go_linux /app
# 通过docker run -e BOT_TOKEN=“xxx”执行
ENV BOT_TOKEN ""

ENTRYPOINT /app/my_assistant_go_linux