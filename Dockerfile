FROM golang:alpine AS builder


LABEL maintainer="Handoko Wae <restapi.handoko@gmail.com>"
# Redundant, current golang images already include ca-certificates
RUN apk --no-cache add ca-certificates
ENV MYSQL_HOST=${MYSQL_HOST}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
ENV MYSQL_DATABASE=${MYSQL_DATABASE}
ENV MYSQL_USER=${MYSQL_USER}
ENV MYSQL_PORT=${MYSQL_PORT}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USERNAME=${DB_USERNAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}
ENV APP_PORT=${APP_PORT}
ENV USERNAME_DAPO=${USERNAME_DAPO}
ENV PASSWORD_DAPO=${PASSWORD_DAPO}
ENV API_TOKEN_DAPO=${API_TOKEN_DAPO}
ENV URL_DAPO=${URL_DAPO}

# Move to working directory (/build).
WORKDIR /app

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main .

# 2 staged build
FROM scratch
ENV MYSQL_HOST=${MYSQL_HOST}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
ENV MYSQL_DATABASE=${MYSQL_DATABASE}
ENV MYSQL_USER=${MYSQL_USER}
ENV MYSQL_PORT=${MYSQL_PORT}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USERNAME=${DB_USERNAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}
ENV APP_PORT=${APP_PORT}
ENV USERNAME_DAPO=${USERNAME_DAPO}
ENV PASSWORD_DAPO=${PASSWORD_DAPO}
ENV API_TOKEN_DAPO=${API_TOKEN_DAPO}
ENV URL_DAPO=${URL_DAPO}
# copy the ca-certificate.crt from the build stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary and config files from /build to root folder of scratch container.
# COPY --from=builder ["/build/apiserver", ".env", "/"]

COPY --from=builder /app/main .
EXPOSE 1992
CMD ["main"]







# FROM golang:alpine as builder
# ENV MYSQL_HOST=${MYSQL_HOST}
# ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
# ENV MYSQL_DATABASE=${MYSQL_DATABASE}
# ENV MYSQL_USER=${MYSQL_USER}
# ENV MYSQL_PORT=${MYSQL_PORT}
# ENV DB_HOST=${DB_HOST}
# ENV DB_PORT=${DB_PORT}
# ENV DB_USERNAME=${DB_USERNAME}
# ENV DB_PASSWORD=${DB_PASSWORD}
# ENV DB_NAME=${DB_NAME}
# ENV APP_PORT=${APP_PORT}
# ENV USERNAME_DAPO=${USERNAME_DAPO}
# ENV PASSWORD_DAPO=${PASSWORD_DAPO}
# ENV API_TOKEN_DAPO=${API_TOKEN_DAPO}
# ENV URL_DAPO=${URL_DAPO}
# WORKDIR /app

# # This will download all certificates (ca-certificates) and builds it in a
# # single file under /etc/ssl/certs/ca-certificates.crt (update-ca-certificates)
# # I also add git so that we can download with `go mod download` and
# # tzdata to configure timezone in final image
# RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
#     update-ca-certificates

# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN  GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o main ${MAIN_PATH}

# # Golang can run in a scratch image, so that, the only thing that your docker 
# # image contains is your executable
# FROM scratch
# ENV MYSQL_HOST=${MYSQL_HOST}
# ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
# ENV MYSQL_DATABASE=${MYSQL_DATABASE}
# ENV MYSQL_USER=${MYSQL_USER}
# ENV MYSQL_PORT=${MYSQL_PORT}
# ENV DB_HOST=${DB_HOST}
# ENV DB_PORT=${DB_PORT}
# ENV DB_USERNAME=${DB_USERNAME}
# ENV DB_PASSWORD=${DB_PASSWORD}
# ENV DB_NAME=${DB_NAME}
# ENV APP_PORT=${APP_PORT}
# ENV USERNAME_DAPO=${USERNAME_DAPO}
# ENV PASSWORD_DAPO=${PASSWORD_DAPO}
# ENV API_TOKEN_DAPO=${API_TOKEN_DAPO}
# ENV URL_DAPO=${URL_DAPO}
# LABEL maintainer="Handoko Wae <restapi.handoko@gmail.com>"
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# # This line will copy all certificates to final image
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# WORKDIR /app
# COPY --from=builder /app/main .
# EXPOSE 1992
# CMD ["/app/main"]







# FROM golang:1.17.9-alpine3.15 as builder
# LABEL maintainer="Handoko Wae <restapi.handoko@gmail.com>"

# RUN apk --no-cache add ca-certificates

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download all

# COPY . ./

# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# RUN go build -o main main.go

# # CMD ["go", "run", "/app/main.go"]
# FROM alpine:3.15
# ENV MYSQL_HOST=${MYSQL_HOST}
# ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
# ENV MYSQL_DATABASE=${MYSQL_DATABASE}
# ENV MYSQL_USER=${MYSQL_USER}
# ENV MYSQL_PORT=${MYSQL_PORT}
# ENV DB_HOST=${DB_HOST}
# ENV DB_PORT=${DB_PORT}
# ENV DB_USERNAME=${DB_USERNAME}
# ENV DB_PASSWORD=${DB_PASSWORD}
# ENV DB_NAME=${DB_NAME}
# ENV APP_PORT=${APP_PORT}
# ENV USERNAME_DAPO=${USERNAME_DAPO}
# ENV PASSWORD_DAPO=${PASSWORD_DAPO}
# ENV API_TOKEN_DAPO=${API_TOKEN_DAPO}
# ENV URL_DAPO=${URL_DAPO}
# WORKDIR /app
# COPY --from=builder /app/main .

# EXPOSE 1992
# CMD ["/app/main"]



# FROM golang:1.17.9-alpine

# RUN apk --no-cache add ca-certificates

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download all

# COPY . ./

# CMD ["go", "run", "/app/main.go"]