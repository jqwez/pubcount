FROM golang:latest

RUN apt-get update && apt-get install -y make

RUN groupadd -r coolgroup && useradd -r -g coolgroup cool

WORKDIR /app

RUN chown -R cool:coolgroup /app

USER cool

COPY --chown=cool:coolgroup . .

CMD ["make", "build-run"]
