package structutil

import (
	"github.com/jokruger/got"
)

func ID[I any, T got.Identifiable[I]](s T) I {
	return s.ID()
}

func ProductID[I any, T got.ProductIDProvider[I]](s T) I {
	return s.ProductID()
}

func CustomerID[I any, T got.CustomerIDProvider[I]](s T) I {
	return s.CustomerID()
}

func EventID[I any, T got.EventIDProvider[I]](s T) I {
	return s.EventID()
}

func AccountID[I any, T got.AccountIDProvider[I]](s T) I {
	return s.AccountID()
}

func Name[I any, T got.Named[I]](s T) I {
	return s.Name()
}

func EventName[I any, T got.EventNameProvider[I]](s T) I {
	return s.EventName()
}

func ValueTime[I any, T got.ValueTimeProvider[I]](s T) I {
	return s.ValueTime()
}

func CreatedAt[I any, T got.CreatedAtProvider[I]](s T) I {
	return s.CreatedAt()
}

func UpdatedAt[I any, T got.UpdatedAtProvider[I]](s T) I {
	return s.UpdatedAt()
}

func Amount[I any, T got.AmountProvider[I]](s T) I {
	return s.Amount()
}

func TransactionID[I any, T got.TransactionIDProvider[I]](s T) I {
	return s.TransactionID()
}

func Priority[I any, T got.PriorityProvider[I]](s T) I {
	return s.Priority()
}
