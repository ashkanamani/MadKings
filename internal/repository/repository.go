package repository

import (
	"context"
	"errors"
	"github.com/ashkanamani/madkings/internal/entity"
)

var (
	ErrorNotFound = errors.New("entity not found")
)

type AccountRepository interface {
	CommonBehaviour[entity.Account]
}

type CommonBehaviour[T entity.Entity] interface {
	Get(context.Context, entity.ID) (T, error)
	Save(context.Context, T) error
}
