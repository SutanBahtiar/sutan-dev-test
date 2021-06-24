# build stage
FROM golang:1.16-alpine3.13 AS builder
LABEL maintainer="sutanirojim"
WORKDIR /app
COPY . .
# RUN go build -o start-server .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o start-server .

# run stage
# FROM alpine:3.13
# WORKDIR /app
# COPY --from=builder /app/start-server .

FROM scratch
WORKDIR /app
COPY --from=builder /app/start-server .

# ENV PORT 8010
EXPOSE 8010
CMD [ "/app/start-server" ]
# ENTRYPOINT [ "/app/start-server" ]