# golang-basic
migrate -path db/migration -database "postgres://ems-fullstack:123456a@@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate -path C:\Users\admin\Documents\improve-golang\golang-basic\web\db\migration -database postgres://simple_bank:123456a@@localhost:5433/simple_bank?sslmode=disable -verbose up

migrate -path db/migration -database "postgres://simple_bank:123456a@@localhost:5432/simple_bank?sslmode=disable" -verbose up