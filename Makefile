# Dependencies
vendor	:
	cd ./faker && go mod vendor

# Docker compose
up	:
	docker-compose up -d

down	:
	docker-compose down

# Faker
fake	:
ifdef duration
	cd ./faker/cmd && go run faker.go $(duration)
else
	cd ./faker/cmd && go run faker.go
endif
