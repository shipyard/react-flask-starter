VOLUME=$(shell basename $(PWD))

develop: clean build migrations.upgrade run

clean:
	docker compose rm -vf

build:
	docker compose build

run:
	docker compose up

frontend-shell:
	docker compose run frontend \
	  sh

backend-shell:
	docker compose run worker \
	  sh

backend-go-shell:
	docker compose run backend-go \
	  sh

python-shell:
	docker compose run worker \
	  poetry run flask shell

postgres-shell:
	docker compose run postgres \
		sh

postgres.data.delete: clean
	docker volume rm $(VOLUME)_postgres

postgres.start:
	docker compose up -d postgres
	docker compose exec postgres \
	sh -c 'sleep 1'
	# The 'nc' command is not available on the updated postgres image; Change to use our own version of the Dockerfile 
	# and install or find some alternate.
	  # sh -c 'while ! nc -z postgres 5432; do sleep 0.1; done'

migrations.blank: postgres.start
	docker compose run worker \
	  poetry run flask db revision

migrations.create: postgres.start
	docker compose run worker \
	  poetry run flask db migrate

migrations.upgrade: postgres.start
	docker compose run worker \
	  poetry run flask db upgrade

migrations.heads: postgres.start
	docker compose run worker \
	  poetry run flask db heads
