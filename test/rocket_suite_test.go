// +acceptance

package test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"odesch.com/odesch/grpc-app/protos/rocket/v1"
	"testing"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("Add new rocket successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				R: &rocket.Rocket{
					Id:   "494",
					Name: "v1",
					Type: "Falcon",
				},
			},
		)
		fmt.Println(resp.R)
		if assert.NoError(s.T(), err) {
			assert.Equal(s.T(), "494", resp.R.Id)
		}
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
