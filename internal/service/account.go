package service

import (
	"context"
	"github.com/ashkanamani/madkings/internal/entity"
	"github.com/ashkanamani/madkings/internal/repository"
)

type AccountService struct {
	accounts repository.AccountRepository
}

func NewAccountService(accounts repository.AccountRepository) *AccountService {
	return &AccountService{
		accounts: accounts,
	}
}
func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) error {
	return a.accounts.Save(ctx, account)
}
