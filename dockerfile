FROM golang:1.21

WORKDIR /jacc
COPY . /jacc/
RUN go build -o main .
EXPOSE 3000
ENTRYPOINT ["./main"]
