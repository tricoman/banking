package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Arnau", "Sant Feliu de codines", "08182", "1981-10-10", "1"},
		{"2", "Perico", "Barcelona", "08182", "2000-12-05", "1"},
		{"3", "Palotes", "Vic", "08182", "1994-08-01", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
