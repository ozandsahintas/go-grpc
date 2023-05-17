package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"odesch.com/odesch/grpc-app/internal/rocket"
	"os"
)

type Store struct {
	db *sqlx.DB
}

// New - returns a new store or error
func New() (Store, error) {

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUsername, dbTable, dbPassword, dbSSLMode)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}

	return Store{db: db}, nil
}

func (s Store) GetRocketById(id string) (rocket.Rocket, error) {
	var r rocket.Rocket
	row := s.db.QueryRow(`SELECT id, name, type FROM rockets WHERE id=$1;`, id)
	err := row.Scan(&r.Id, &r.Name, r.Type)
	if err != nil {
		log.Println(err.Error())
		return rocket.Rocket{}, nil
	}

	return r, nil
}

func (s Store) InsertRocket(r rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedQuery(`INSERT INTO rockets (id,name,type) VALUES (:id, :name, :type)`, r)
	if err != nil {
		log.Println(err.Error())
		return rocket.Rocket{}, nil
	}
	return rocket.Rocket{
		Id:   r.Id,
		Name: r.Name,
		Type: r.Type,
	}, nil
}

func (s Store) DeleteRocket(id string) error {
	/*uid, err := uuid.FromString(id)
	if err != nil {
		return err
	}*/

	_, err := s.db.Exec(`DELETE FROM rockets WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}
