c_m:
	# creates a new migration
	migrate create -ext sql -dir db/migrations -seq $(name)

p_up:
	# postgres up - creates postgres server
	docker-compose up -d

p_down:
	# postgres down - deletes postgres server
	docker-compose down

db_up:
	docker exec -it fingreat_postgres createdb --username=spicywords --owner=spicywords fingreat_db

db_down:
	docker exec -it fingreat_postgres dropdb --username=spicywords fingreat_db