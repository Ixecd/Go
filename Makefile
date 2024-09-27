dockerPostgresUp:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qc -d postgres:17-alpine

dockerPostgresDown:


dockerPostgresLog:
	sudo docker logs postgres17

dockerMigrationUp:
	migrate create -ext sql -dir base/db/migration -seq init-schema


dockerRun:
	sudo docker exec -it postgres17 psql -U root