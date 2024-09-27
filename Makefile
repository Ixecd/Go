# 创建simple_bank数据库
createdb:
	sudo docker exec -it postgres17 createdb --username=root --owner=root simple_bank

# 删除数据库
dropdb:
	sudo docker exec -it postgres17 dropdb simple_bank

# 启动postgres容器
postgres17:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qc -d postgres:17-alpine

# 向上迁移数据库
migrateup:
	migrate -path base/db/migration -database "postgresql://root:qc@localhost:5432/simple_bank?sslmode=disable" -verbose up

# 向下迁移数据库
migratedown:
	migrate -path base/db/migration -database "postgresql://root:qc@localhost:5432/simple_bank?sslmode=disable" -verbose down

dockerPostgresLog:
	sudo docker logs postgres17

dockerMigrationUp:
	migrate create -ext sql -dir base/db/migration -seq init-schema

dockerRun:
	sudo docker exec -it postgres17 psql -U root

.PHONY: createdb dropdb postgres17 migrateup migratedown dockerPostgresLog dockerMigrationUp dockerRun