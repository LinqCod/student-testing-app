postgres:
	docker run --name student-testing-postgres-15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15

create_db:
	docker exec -it student-testing-postgres-15 createdb --username=root --owner=root student-test

drop_db:
	docker exec -it student-testing-postgres-15 dropdb student-test

.PHONY: postgres create_db drop_db