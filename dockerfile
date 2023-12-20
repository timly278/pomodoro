# Build stage
FROM golang:1.21.2-alpine3.18 AS build-stage

WORKDIR /app4

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get

RUN go build -o ./main .

#Run stage
FROM alpine:3.18 AS run-stage

WORKDIR /app4

COPY --from=build-stage /app4/. /app4/.
COPY app.env /app4/.

EXPOSE 8080

CMD [ "/app4/main"]

