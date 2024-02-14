FROM node:12.18.3-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN npm install

CMD ["npm", "start"]

# Path: broker-service.dockerfile
FROM golang:1.20-alpine AS builder

