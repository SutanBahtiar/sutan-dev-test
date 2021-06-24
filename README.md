# Simple tcp socket

go run .

#### Create docker image
`$docker build -t sutan-dev-test .`

#### Pull image
https://hub.docker.com/repository/docker/sutanirojim/sutan-dev-test

#### Create network
`$docker network create dev_network`

#### Running image
`$docker run -v /Users/sutan/workspace/dev-tester:/root/share --rm -it --network dev-network server-dev /app/start-server`
