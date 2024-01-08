# Build stage
FROM --platform=amd64 golang:alpine3.18 AS build-stage
# FROM golang:alpine3.18 AS build-stage

WORKDIR /app4

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN go get

RUN go build -o ./main .

#Run stage
FROM --platform=amd64 golang:alpine3.18 AS run-stage
# FROM golang:alpine3.18 AS run-stage

WORKDIR /app4

COPY --from=build-stage /app4/. /app4/.
COPY app.env /app4/.

EXPOSE 8080

CMD [ "/app4/main"]

