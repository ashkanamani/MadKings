package integrationtest

import (
	"context"
	"fmt"
	"github.com/ashkanamani/madkings/internal/entity"
	"github.com/ashkanamani/madkings/internal/repository"
	"github.com/ashkanamani/madkings/internal/repository/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (t testType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviour_SetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)

	ctx := context.Background()
	cb := repository.NewRedisCommonBehaviour[testType](redisClient)
	err = cb.Save(ctx, testType{ID: "33", Name: "pelamar"})
	assert.NoError(t, err)

	val, err := cb.Get(ctx, entity.NewID("testType", "33"))
	assert.NoError(t, err)
	assert.Equal(t, "pelamar", val.Name)
	assert.Equal(t, "33", val.ID)

	val, err = cb.Get(ctx, entity.NewID("testType", "34"))
	assert.ErrorIs(t, repository.ErrorNotFound, err)
}
