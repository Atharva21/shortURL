# build the frontend
FROM node:16-alpine3.16 as WEB
WORKDIR /web
COPY web/ .
RUN npm install
RUN npm run build

# build backend
FROM golang:alpine3.10 as BACKEND
WORKDIR /app
COPY backend/go.* .
RUN go get ./...
COPY backend/ .
COPY --from=WEB /web/build/ ./public
RUN go build -o bin/app
CMD [ "./bin/app" ]
 