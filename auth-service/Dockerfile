FROM golang

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-service .

ENV GO_PORT 8181

CMD ["./auth-service"]
