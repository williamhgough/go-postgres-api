run-dev:
	docker-compose -f docker-compose-dev.yml up --build
run-prod:
	docker-compose up --build
default:
	make run-prod