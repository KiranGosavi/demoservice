FROM golang:latest as dev

WORKDIR /demoservice

COPY . .

RUN go mod download
RUN go build -o mydemoservice .

EXPOSE 8080
CMD [ "/demoservice/mydemoservice" ]