dev_up:
	docker-compose -f docker-compose-dev.yml run --name supchecker --rm supchecker

prod_up:
	docker-compose -f docker-compose-prod.yml up -d --build

prod_down:
	docker-compose -f docker-compose-prod.yml down