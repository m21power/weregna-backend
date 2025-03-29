sdocker:
	sudo sysctl -w kernel.apparmor_restrict_unprivileged_userns=0

startdocker: sdocker
	systemctl --user start docker-desktop

# Create migration file
startmigration:
	docker run -it --rm --volume "$(shell pwd)/db:/db" migrate/migrate:v4.17.0 create -ext sql -dir /db/migrations -seq messages

# Start PostgreSQL container
rundocker:
	docker run --name weregna-db -p 5432:5432 -e POSTGRES_PASSWORD=abate -d postgres:latest

# Create database in PostgreSQL container
createdb:
	docker exec -it weregna-db psql -U postgres -c "CREATE DATABASE weregna;"

# Migrate up for PostgreSQL
migrateup:
	docker run -it --rm --network host --volume "$(shell pwd)/db:/db" migrate/migrate:v4.17.0 -path /db/migrations -database "postgres://postgres:abate@localhost:5432/weregna?sslmode=disable" up

# Migrate down for PostgreSQL
migratedown:
	docker run -it --rm --network host --volume "$(shell pwd)/db:/db" migrate/migrate:v4.17.0 -path /db/migrations -database "postgres://postgres:abate@localhost:5432/weregna?sslmode=disable" down