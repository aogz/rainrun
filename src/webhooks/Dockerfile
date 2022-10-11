FROM golang:1.19

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN mkdir /app
COPY . /app
WORKDIR /app
CMD ["air", "-d"]
