FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# build go app
RUN go mod download
# RUN go build -o user-store ./cmd/apiserver/main.go
RUN make build

CMD ["./user-store"]