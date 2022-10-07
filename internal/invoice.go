package internal

type Invoice struct {
	Name         string
	Address      string
	InvoiceItems []*InvoiceData
}

func CreateInvoice(name string, address string, invoiceItems []*InvoiceData) *Invoice {
	return &Invoice{
		Name:         name,
		Address:      address,
		InvoiceItems: invoiceItems,
	}
}

func (i *Invoice) CalculateInvoiceTotalAmount() float64 {
	var invoiceTotalAmount int64 = 0
	for _, data := range i.InvoiceItems {
		amount := data.CalculateTotalAmount()
		invoiceTotalAmount += amount
	}

	totalAmount := float64(invoiceTotalAmount) / 100

	return totalAmount
}
