package wallet

import (
	"github.com/google/uuid"
	"errors"
	"github.com/Shahlojon/wallet/pkg/types"
)

//ErrAccountNotFound - пользователь не найден
var ErrAccountNotFound = errors.New("account not found")
var ErrPaymentNotFound = errors.New("payment not found")
//ErrAmountMustBePositive - счёт не может быть пустым
var ErrAmountMustBePositive = errors.New("amount must be greater than zero")
//ErrPhoneRegistered - телефон уже регитрирован
var ErrPhoneRegistered = errors.New("phone already registred")
//ErrNotEnoughtBalance - на счете недостаточно средств
var ErrNotEnoughtBalance = errors.New("account not enough balance")
type Service struct{
	nextAccountID int64
	accounts []*types.Account
	payments []*types.Payment
}

// type Error string

// func (e Error) Error() string {
// 	return string(e)
// }
// type Messenger interface {
// 	Send(message string) (ok bool)
// 	Receive() (message string, ok bool)
// }

// type Telegram struct{

// }

// type (t *Telegram) Send(message string) bool {
// 	return true
// }

// type (t *Telegram) Receive() (message string, ok bool) {
// 	return " ",true
// }

// RegisterAccount - метод для регистрация нового прользователя.
func RegisterAccount(s *Service, phone types.Phone) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return
		}
	}
	s.nextAccountID++
	s.accounts = append(s.accounts, &types.Account{
		ID: s.nextAccountID,
		Phone: phone,
		Balance: 0,
	})
}

// RegisterAccount - метод для регистрация нового прользователя.
func (s *Service) RegisterAccount(phone types.Phone)  (*types.Account, error){
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountID++
	account := &types.Account {
		ID: s.nextAccountID,
		Phone: phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

//Deposit method
func (s *Service) Deposit(accountID int64, amount types.Money)  error{
	if amount<=0 {
		return ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID ==accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return ErrAccountNotFound
	}

	//зачисления средств пока не рассмотрим как платёж
	account.Balance +=amount
	return nil
}

//Pay метод для регистрации платежа
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error)  {
	if amount<=0 {
		return nil, ErrAmountMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account =acc
			break
		}
	}
	if account ==nil {
		return nil, ErrAccountNotFound
	}

	if account.Balance < amount {
		return nil, ErrNotEnoughtBalance
	}

	account.Balance-=1
	paymentID :=uuid.New().String()
	payment := &types.Payment{
		ID: paymentID,
		AccountID: accountID,
		Amount: amount,
		Category: category,
		Status: types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

func (s *Service) FindAccountByID(accountID int64) (*types.Account, error) {
	var account *types.Account
	for _, accounts := range s.accounts {
		if accounts.ID == accountID {
			account = accounts
			break
		} 
	}

	if account == nil {
		return nil, ErrAccountNotFound
	}

	return account, nil 
}

// Reject метод для отмены покупок
func (s *Service) Reject(paymentID string) error  {
	var targetPayment *types.Payment
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			targetPayment = payment
			break
		}
	}
	if targetPayment == nil {
		return ErrPaymentNotFound
	}

	var targetAccount *types.Account
	for _, account := range s.accounts {
		if account.ID == targetPayment.AccountID {
			targetAccount = account
			break
		}
	}
	if targetAccount == nil {
		return ErrAccountNotFound
	}

	targetPayment.Status = types.PaymentStatusFail
	targetAccount.Balance += targetPayment.Amount
	return nil
}

func (s *Service) FindPaymentByID(paymentID string) (*types.Payment, error) {
	for _, payment := range s.payments {
		if payment.ID == paymentID {
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound
}