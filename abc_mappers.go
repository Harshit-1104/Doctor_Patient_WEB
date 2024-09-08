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

type Bank struct {
	ID int
}

func (c *Bank) isEmpty() bool {
	return c == nil || c.ID == 0
}

type BankProto struct {
	ID int
}

func toProtoFromBank(bank *Bank) *BankProto {
	if bank == nil {
		return nil
	}

	return &BankProto{
		ID: bank.ID,
	}
}
