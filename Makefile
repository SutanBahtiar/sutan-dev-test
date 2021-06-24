buildserver
	docker build -t sutan-dev-test .
createnetwork
	docker network create dev_network
runserver
	docker run -v /Users/sutan/workspace/dev-tester:/root/share --rm -it --network dev-network server-dev /app/start-server
dockerfileclient
	mv Dockerfile Dockerfile-server
dockerfileserver
	mv Dockerfile-server Dockerfile
buildclient
	docker build -t client-dev-test .
runclient
	docker run -v /Users/sutan/workspace/dev-tester:/root/share --rm -it --network dev-network client-dev /app/start-client

.PHONY: buildserver createnetwork runserver dockerfileclient dockerfileserver buildclient runclient