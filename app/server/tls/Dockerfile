FROM golang:1.15
WORKDIR /go/src/app
COPY . .
RUN go build && mv app /app

# Setup our non root user
RUN groupadd -g 1000 -r application && useradd --no-log-init -u 1000 -g application application

USER 1000
CMD ["/app"]
