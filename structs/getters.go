package structs

import (
	. "github.com/jokruger/got/ifs"
)

func GetID[I any, T IDProvider[I]](s T) I {
	return s.GetID()
}

func GetProductID[I any, T ProductIDProvider[I]](s T) I {
	return s.GetProductID()
}

func GetCustomerID[I any, T CustomerIDProvider[I]](s T) I {
	return s.GetCustomerID()
}

func GetEventID[I any, T EventIDProvider[I]](s T) I {
	return s.GetEventID()
}

func GetAccountID[I any, T AccountIDProvider[I]](s T) I {
	return s.GetAccountID()
}

func GetName[I any, T NameProvider[I]](s T) I {
	return s.GetName()
}

func GetValueTime[I any, T ValueTimeProvider[I]](s T) I {
	return s.GetValueTime()
}

func GetCreatedAt[I any, T CreatedAtProvider[I]](s T) I {
	return s.GetCreatedAt()
}

func GetUpdatedAt[I any, T UpdatedAtProvider[I]](s T) I {
	return s.GetUpdatedAt()
}

func GetAmount[I any, T AmountProvider[I]](s T) I {
	return s.GetAmount()
}

func GetTransactionID[I any, T TransactionIDProvider[I]](s T) I {
	return s.GetTransactionID()
}
