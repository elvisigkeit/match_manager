FROM golang:alpine
ENV GOPATH ""

ADD . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT $HOME/go/bin/matchManager
EXPOSE 8000