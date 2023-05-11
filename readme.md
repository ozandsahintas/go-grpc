Go gRPC Services
========================

Unit Tests
go install github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
go get github.com/stretchr/testify/assert

//go:generate mockgen -destination=rocket_mocks_test.go -package rocket odesch.com/odesch/grpc-app/internal/rocket Store
go test ./internal/rocket -v
go test ./... -v

docker run --name <NAME_OF_THE_DB> -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

go get github.com/jmoiron/sqlx
go get github.com/golang-migrate/migrate/v4/database/postgres

docker-compose up --build

docker ps -a
docker exec -it <CONTAINER_ID> bash
bash-5.0# psql -U postgres
postgres=# \c postgres
postgres=# \dt
