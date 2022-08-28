FROM node:latest as web
WORKDIR /src
COPY . .
RUN npm install --global yarn && cd web && yarn build

FROM golang:1.19 as base
WORKDIR /src
COPY --from=web . .
RUN CGO_ENABLED=0 go build -o app

FROM alpine:latest as final
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=base /src/app /usr/local/bin/tianyi
COPY tianyi.yml tianyi.yml
EXPOSE 8080
CMD ["tianyi", "server", "run"]