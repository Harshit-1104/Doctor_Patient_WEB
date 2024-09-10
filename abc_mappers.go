package main

type Customer struct {
	ID int
}

func (c *Customer) isEmpty() bool {
	return c == nil || c.ID == 0
}

type CustomerProto struct {
	ID int
}

func toProtoFromCustomer(cust *Customer) *CustomerProto {
	if cust.isEmpty() {
		return nil
	}

	return &CustomerProto{
		ID: cust.ID,
	}
}

type SubBank struct {
	ID   int
	stat string
}

type SubBankProto struct {
	ID   int
	stat string
}

type Bank struct {
	ID      int
	stat    string
	SubBank *SubBank
}

func (c *Bank) isEmpty() bool {
	return c == nil || (c.ID == 0 && c.stat == "")
}

type BankProto struct {
	ID      int
	stat    string
	SubBank *SubBankProto
}

func toProtoFromBank(bank *Bank) *BankProto {
	if bank == nil {
		return nil
	}

	subBankProto := &SubBankProto{
		ID:   bank.SubBank.ID,
		stat: bank.SubBank.stat,
	}

	return &BankProto{
		ID:      bank.ID,
		stat:    bank.stat,
		SubBank: subBankProto,
	}
}
