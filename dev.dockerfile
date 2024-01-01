FROM golang:1.21

WORKDIR /jacc
COPY . /jacc/
RUN apt-get update && apt-get install wkhtmltopdf -y
RUN go mod download
RUN go mod verify
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/cosmtrek/air@latest
CMD air -c .air.toml
