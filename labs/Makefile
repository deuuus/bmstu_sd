docker_postgres:
	docker run --name=policl_db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

docker_mysql:
	docker run --name=example2 -e MYSQL_PASSWORD='qwerty' -p 3306:3306 -d --rm mysql/mysql-server:latest

migrate_up:
	./migrate -source file://schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate_up_mysql:
	./migrate -source file://schema -database 'mysql://root:qwerty@tcp(localhost:3306)/mysql' up

migrate_down:
	./migrate -source file://schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

run:
	go run cmd/main.go