FROM golang:1.20
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
WORKDIR /app
CMD [ "air", "-c", "/app/.air.toml" ]
