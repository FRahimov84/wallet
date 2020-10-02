package types

//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д)
type Money int64

//Status представляет собой статус платежа
type Status string

//PaymentCategory представляет собой категорию, в которой был совершен платеж (авто, аптеки, рестораны и т.д).
type PaymentCategory string

//PaymentStatus представляет собой статус платежа.
type PaymentStatus string

//Предопределенные статусы платежей
const (
	PaymentStatusOk Status ="OK"
	PaymentStatusFail Status = "FAIL"
	PaymentStatusInProgress Status = "INPROGRESS"
)

//Payment представляет информацию о платеже
type Payment struct{
	ID string
	Amount Money
	Category PaymentCategory
	Status PaymentStatus
}

type Phone string

//Account представляет информацию о счёте пользователя
type Account struct {
	ID int64
	Phone Phone
	Balance Money
}