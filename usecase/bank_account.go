package usecase

import (
	"github.com/hiiamninna/market_place/collections"
	"github.com/hiiamninna/market_place/repository"
)

// Use case => to list possible case when do unit testing
type BankAccount struct {
	repo repository.BankAccount
}

func NewBankAccountUseCase(repo repository.BankAccount) BankAccount {
	return BankAccount{
		repo: repo,
	}
}

func (u BankAccount) Create(input collections.BankAccountInput) error {
	return u.repo.Create(input)
}

func (u BankAccount) Update(input collections.BankAccountInput) error {
	return u.repo.Update(input)
}

func (u BankAccount) Delete(id, userID string) error {
	return u.repo.Delete(id, userID)
}

func (u BankAccount) GetByID(id, userID string) (collections.BankAccount, error) {
	return u.repo.GetByID(id, userID)
}

func (u BankAccount) List(userID string) ([]collections.BankAccount, error) {
	return u.repo.List(userID)
}
