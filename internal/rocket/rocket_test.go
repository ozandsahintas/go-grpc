package rocket

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

//White-box testing

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("Get Rocket By Id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().GetRocketById(id).Return(
			Rocket{
				Id: id,
			}, nil)
		rocketService := New(rocketStoreMock)
		r, err := rocketService.GetRocketById(context.Background(), id)

		if assert.NoError(t, err) {
			assert.Equal(t, "UUID-1", r.Id)
		}
	})

	t.Run("Insert Rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().InsertRocket(Rocket{
			Id: id,
		}).Return(
			Rocket{
				Id: id,
			}, nil)
		rocketService := New(rocketStoreMock)
		r, err := rocketService.InsertRocket(context.Background(), Rocket{
			Id: id,
		})

		if assert.NoError(t, err) {
			assert.Equal(t, "UUID-1", r.Id)
		}
	})

	t.Run("Delete Rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().DeleteRocket(id).Return(nil)
		rocketService := New(rocketStoreMock)
		err := rocketService.DeleteRocket(id)

		assert.NoError(t, err)
	})
}
