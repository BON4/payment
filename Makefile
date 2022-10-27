DB_URL=postgresql://root:secret@localhost:5432/payment?sslmode=disable
TEST_DB_URL=postgresql://root:secret@localhost:5432/payment_test?sslmode=disable

postgres:
	docker run --name psql_payments -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it psql_payments createdb --username=root --owner=root payment && docker exec -it psql_payments createdb --username=root --owner=root payment_test

dropdb:
	docker exec -it psql_payments dropdb payment && docker exec -it psql_payments dropdb payment_test

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up && migrate -path db/migration -database "$(TEST_DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down && migrate -path db/migration -database "$(TEST_DB_URL)" -verbose down

boil:
	sqlboiler --tag "csv" psql 

.PHONY: postgres createdb dropdb migrateup migratedown boil
