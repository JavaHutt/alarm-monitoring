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
	cd ./faker && $(MAKE) run $(duration)
else
	cd ./faker && $(MAKE) run
endif

# Monitor
run	:
	cd ./monitor && $(MAKE) run
