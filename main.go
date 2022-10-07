package main

import (
	"fmt"
	internal "marvinhosea/invoices/internal"
)

func main() {
	// Generate sample invoice data
	ecommerceInvoiceData, err := internal.NewInvoiceData("Ecommerce application", 1, 3000.50)
	if err != nil {
		panic(err)
	}
	laptopInvoiceData, err := internal.NewInvoiceData("Macbook Pro", 1, 200.70)
	if err != nil {
		panic(err)
	}
	// Invoice Items collection
	invoiceItems := []*internal.InvoiceData{ecommerceInvoiceData, laptopInvoiceData}

	// Create single invoice
	invoice := internal.CreateInvoice("Example Shop1", "Example address", invoiceItems)
	err = internal.GenerateInvoicePdf(*invoice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The Total Invoice Amount is: %f", invoice.CalculateInvoiceTotalAmount())
}
