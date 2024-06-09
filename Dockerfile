FROM golang:latest

RUN apt-get update && apt-get install -y make

WORKDIR /app

COPY . .

CMD ["make", "build-run"]
