build:
	go build

docker:
	docker build ./ --tag dumbstored
run:
	docker-compose up
