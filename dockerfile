FROM golang:1.21

WORKDIR /jacc
COPY . /jacc/
RUN apt-get update && apt-get install wkhtmltopdf -y
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go build -o main .
EXPOSE 3000
ENTRYPOINT ["./main"]
