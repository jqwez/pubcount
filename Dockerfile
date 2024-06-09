FROM golang:latest

RUN apt-get update && apt-get install -y make curl

WORKDIR /app

COPY . .

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.4/tailwindcss-linux-arm64
RUN chmod +x tailwind-linux-arm64 && mv tailwind-linux-arm64 third_party/twind

CMD ["make", "build-run"]
