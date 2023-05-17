//go:generate mockgen -destination=rocket_mocks_test.go -package rocket odesch.com/odesch/grpc-app/internal/rocket Store
package rocket

import "context"

// Rocket - should contain the definition of our rocket
type Rocket struct {
	Id      string
	Name    string
	Type    string
	Flights int
}

// Store - defines the interface we expect
// our database implementation to follow
type Store interface {
	GetRocketById(id string) (Rocket, error)
	InsertRocket(r Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - responsible our rocket service updating rocket inventory
type Service struct {
	Store Store
}

func New(s Store) Service {
	return Service{
		Store: s,
	}
}

// GetRocketById - retrieves a rocket based on the ID from the store
func (s Service) GetRocketById(ctx context.Context, id string) (Rocket, error) {
	r, err := s.Store.GetRocketById(id)
	if err != nil {
		return Rocket{}, err
	}
	return r, nil
}

// InsertRocket - Insert new rocket into the store
func (s Service) InsertRocket(ctx context.Context, r Rocket) (Rocket, error) {
	r, err := s.Store.InsertRocket(r)
	if err != nil {
		return Rocket{}, err
	}
	return r, nil
}

// DeleteRocket - Deletes a rocket from our inventory
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
