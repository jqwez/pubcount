FROM golang:latest

RUN apt-get update && apt-get install -y make git

WORKDIR /app

RUN git clone https://github.com/jqwez/pubcount && cd pubcount

CMD ["make", "build-run"]
