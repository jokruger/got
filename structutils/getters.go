package structutils

import (
	"github.com/jokruger/got"
)

func GetID[I any, T got.Identifiable[I]](s T) I {
	return s.GetID()
}

func GetProductID[I any, T got.ProductIDProvider[I]](s T) I {
	return s.GetProductID()
}

func GetCustomerID[I any, T got.CustomerIDProvider[I]](s T) I {
	return s.GetCustomerID()
}

func GetEventID[I any, T got.EventIDProvider[I]](s T) I {
	return s.GetEventID()
}

func GetAccountID[I any, T got.AccountIDProvider[I]](s T) I {
	return s.GetAccountID()
}

func GetName[I any, T got.Named[I]](s T) I {
	return s.GetName()
}

func GetEventName[I any, T got.EventNameProvider[I]](s T) I {
	return s.GetEventName()
}

func GetValueTime[I any, T got.ValueTimeProvider[I]](s T) I {
	return s.GetValueTime()
}

func GetCreatedAt[I any, T got.CreatedAtProvider[I]](s T) I {
	return s.GetCreatedAt()
}

func GetUpdatedAt[I any, T got.UpdatedAtProvider[I]](s T) I {
	return s.GetUpdatedAt()
}

func GetAmount[I any, T got.AmountProvider[I]](s T) I {
	return s.GetAmount()
}

func GetTransactionID[I any, T got.TransactionIDProvider[I]](s T) I {
	return s.GetTransactionID()
}

func GetPriority[I any, T got.PriorityProvider[I]](s T) I {
	return s.GetPriority()
}
