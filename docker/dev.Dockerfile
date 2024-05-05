FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

RUN export GOPATH=$HOME/go && export PATH=$PATH:$GOROOT/bin:$GOPATH/bin && export PATH=$PATH:$(go env GOPATH)/bin

CMD sleep infinity


