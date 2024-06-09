FROM golang:latest

RUN apt-get update && apt-get install -y make curl

WORKDIR /app

COPY . .

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.4/tailwindcss-linux-arm64
RUN chmod +x tailwindcss-linux-arm64 && mv tailwindcss-linux-arm64 third_party/twind
RUN curl -SLO https://github.com/a-h/templ/releases/download/v0.2.707/templ_Linux_arm64.tar.gz
RUN tar -xzf templ_Linux_arm64.tar.gz && mv templ third_party/templ
RUN go mod download 

CMD ["make", "build-run"]
