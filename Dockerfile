
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk update && apk add --no-cache git bash
RUN go build tinywww.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app /app
ENTRYPOINT /app/tinywww
LABEL Name=dcasati/aks-devops-demo Version=0.0.1
EXPOSE 8000
