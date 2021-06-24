# SERVER
# FROM alpine
FROM golang:1.16
 
WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o start-server .

EXPOSE 5678

ENTRYPOINT ["/app/start-server"]


# # CLIENT
# FROM alpine
# FROM golang:1.16
 
# WORKDIR /app

# COPY . .

# RUN go mod tidy

# RUN go build -o start-client .

# EXPOSE 5678

# ENTRYPOINT ["/app/start-client"]