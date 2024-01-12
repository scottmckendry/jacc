FROM golang:1.21

WORKDIR /jacc
COPY . /jacc/
RUN apt-get update --fix-missing && apt-get install wkhtmltopdf -y
RUN go mod download
RUN go mod verify
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN git clone -b feat-live-proxy https://github.com/ndajr/air ~/.air && cd ~/.air && go install . && cd -
CMD air -c .air.toml
