FROM alpine:latest

WORKDIR /

RUN apk --no-cache add bash curl

COPY . .

RUN chmod 777 /request.sh

CMD ["./request.sh"]
