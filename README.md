# Simple tcp socket

go run .

#### Create docker image
`$docker build -t sutan-dev-test .`

#### Running image
`$docker run -v /Users/sutan/workspace/dev-tester:/root/share --rm -it sutan-dev-test /root/share/socks.sock`
