build:
	go build

docker:
	docker build ./ --tag dumbstored

clean:
	go clean

run: up

up:
	docker-compose up

down:
	docker-compose down

register:
	curl -u "tester:tester123" -XPOST localhost:8080/register

check:
	curl -u "tester:tester123" localhost:8080/check

checkbad: 
	curl -u "tester:BADPASSWORD" localhost:8080/check

psql:
	docker-compose exec db /usr/bin/psql -U dumbstored

start-transfer:
	curl -u "tester:tester123" --header "Length: 1234" -XPOST localhost:8080/store
