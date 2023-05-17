Go gRPC Services
=

Unit Tests
=
```go install github.com/golang/mock/gomock
```
```go get github.com/golang/mock/mockgen
```
```go get github.com/stretchr/testify/assert
```

//go:generate mockgen -destination=rocket_mocks_test.go -package rocket odesch.com/odesch/grpc-app/internal/rocket Store <br />
go test ./internal/rocket -v <br />
go test ./... -v 

DB
=
docker run --name <NAME_OF_THE_DB> -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres

go get github.com/jmoiron/sqlx <br />
go get github.com/golang-migrate/migrate/v4/database/postgres

Docker
=
docker-compose up --build

docker ps -a <br />
docker exec -it <CONTAINER_ID> bash <br />
bash-5.0# psql -U postgres <br />
postgres=# \c postgres <br />
postgres=# \dt <br />

Protobuf & gRPC
=
brew install protobuf <br />
protoc --version

brew install protoc-gen-go-grpc <br />
cd protos <br/>
make build

