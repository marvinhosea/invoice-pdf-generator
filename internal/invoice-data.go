package internal

import "errors"

type InvoiceData struct {
	Title       string
	Quantity    int64
	Price       int64
	TotalAmount int64
}

func (d *InvoiceData) CalculateTotalAmount() int64 {
	totalAmount := d.Quantity * d.Price
	return totalAmount
}

func (d *InvoiceData) ReturnItemTotalAmount() float64 {
	totalAmount := d.CalculateTotalAmount()
	converted := float64(totalAmount) / 100
	return converted
}

func (d *InvoiceData) ReturnItemPrice() float64 {
	converted := float64(d.Price) / 100
	return converted
}

func NewInvoiceData(title string, qty int64, price interface{}) (*InvoiceData, error) {
	var convertedPrice int64

	switch priceValue := price.(type) {
	case int64:
		convertedPrice = priceValue * 100
	case int:
		convertedPrice = int64(priceValue * 100)
	case float32:
		convertedPrice = int64(priceValue * 100)
	case float64:
		convertedPrice = int64(priceValue * 100)
	default:
		return nil, errors.New("type not permitted")
	}

	return &InvoiceData{
		Title:    title,
		Quantity: qty,
		Price:    convertedPrice,
	}, nil
}
