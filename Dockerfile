FROM golang:latest

RUN apt-get update && apt-get install -y make

WORKDIR /app

COPY . .

RUN chmod 777 ./third_party/*

CMD ["make", "build-run"]
