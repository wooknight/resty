#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
# Copy go mod and sum files
COPY ./go.mod  .
RUN go mod download -x
ADD . .
COPY . .

RUN go build -o /go/bin/app -v .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN ls
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
EXPOSE 8080
