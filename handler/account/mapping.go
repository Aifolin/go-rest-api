package account

import (
	"github.com/AlekSi/pointer"
	"github.com/Toshik1978/go-rest-api/handler"
	"github.com/Toshik1978/go-rest-api/repository"
)

const (
	outgoingPayment = "outgoing"
	incomingPayment = "incoming"
)

// mapRepositoryAccount maps repository account model to API
func mapRepositoryAccount(account repository.Account) *handler.Account {
	return &handler.Account{
		UID:       account.UID,
		Currency:  account.Currency,
		Balance:   float64(account.Balance) / 100,
		CreatedAt: account.CreatedAt,
	}
}

// mapRepositoryAccounts maps multiple repository account models to API
func mapRepositoryAccounts(accounts []repository.Account) []handler.Account {
	results := make([]handler.Account, 0, len(accounts))
	for _, account := range accounts {
		results = append(results, *mapRepositoryAccount(account))
	}
	return results
}

// mapRepositoryPayment maps repository payment model to API
func mapRepositoryPayment(payment repository.Payment) *handler.Payment {
	if payment.Amount >= 0 {
		return &handler.Payment{
			UID:       payment.PayerAccountUID,
			SourceUID: nil,
			TargetUID: pointer.ToString(payment.RecipientAccountUID),
			Direction: outgoingPayment,
			Amount:    float64(payment.Amount) / 100,
			CreatedAt: payment.CreatedAt,
		}
	}
	return &handler.Payment{
		UID:       payment.PayerAccountUID,
		SourceUID: pointer.ToString(payment.RecipientAccountUID),
		TargetUID: nil,
		Direction: incomingPayment,
		Amount:    -float64(payment.Amount) / 100,
		CreatedAt: payment.CreatedAt,
	}
}

// mapRepositoryPayments maps repository payment model to API
func mapRepositoryPayments(payments []repository.Payment) []handler.Payment {
	results := make([]handler.Payment, 0, len(payments))
	for _, payment := range payments {
		results = append(results, *mapRepositoryPayment(payment))
	}
	return results
}
