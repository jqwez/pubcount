FROM node:latest AS nodebuild

WORKDIR /tailwind

COPY web web

RUN npm install tailwindcss

RUN cd web && npx tailwindcss



FROM golang:latest

RUN apt-get update && apt-get install -y make

WORKDIR /app

COPY . .
COPY --from=nodebuild /tailwind/web/static /app/web/static

RUN make templ
RUN go build -o pubcount main.go

CMD ["./pubcount"]
