##########################################################
# Prepare a build container with all dependencies inside #
##########################################################
FROM golang:alpine as builder

WORKDIR /build

COPY ./app/ ./
RUN go mod download
RUN go build -o /build main.go

###########################################
# Create clean container with binary only #
###########################################
FROM alpine as exec

RUN apk add --update bash ca-certificates

WORKDIR /app
COPY --from=builder /build /app

ENTRYPOINT ["/app/main"]