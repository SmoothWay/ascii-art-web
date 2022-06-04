FROM golang:latest

LABEL authors="timshowtime & insaneEra" release-date="2022-06-01" repo="https://01.alem.school/git/timshowtime/ascii-art-web-dockerize" 
LABEL version="0.1" defaultport="8080"

COPY . /app

WORKDIR /app

CMD go run ./cmd/