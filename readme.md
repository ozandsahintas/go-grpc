Go gRPC Services
=

Unit Tests
=
```sh
go install github.com/golang/mock/gomock
```
```sh
go get github.com/golang/mock/mockgen
```
```sh
go get github.com/stretchr/testify/assert
```

//go:generate mockgen -destination=rocket_mocks_test.go -package rocket odesch.com/odesch/grpc-app/internal/rocket Store <br />

```sh
go generate
```
```sh
go test ./internal/rocket -v
```
```sh
go test ./... -v 
```

DB
=
```sh
docker run --name <NAME_OF_THE_DB> -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```
```sh
go get github.com/jmoiron/sqlx 
```
```sh
go get github.com/golang-migrate/migrate/v4/database/postgres
```
Docker
=
```sh
docker-compose up --build
```
```sh
docker ps -a
```
```sh
docker exec -it <CONTAINER_ID> bash
```

bash-5.0#
```sh
psql -U postgres
```

postgres=# <br/>
```postgresql
\c postgres
```
```postgresql
\dt
```

Protobuf & gRPC
=
```sh
brew install protobuf | protoc --version
```

```sh
brew install protoc-gen-go-grpc
```
```sh
cd protos | make build
```

````sh
github.com/stretchr/testify/suite
````
```sh
go mod vendor
```

```sh
go test ./test -tags=acceptance -v
```

