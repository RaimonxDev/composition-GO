package invoice

import (
	"github.com/RaimonxDev/composition-GO/pkg/customer"
	"github.com/RaimonxDev/composition-GO/pkg/invoiceItem"
)

// Estructura de una factura
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	// Relacion 1 a muchos agregamos []
	items []invoiceItem.Item
}

// Establecer el total  de una factura
// Como vamos a modificar el valor de una factura tenemos que usar *

func (i *Invoice) SetTotal() {

	for _, item := range i.items {
		// item.Value es nuestro getter declarado en pkg invoiceItems
		i.total += item.Value()
	}

}

func New(country, city string, customer customer.Customer, items []invoiceItem.Item) Invoice {

	return Invoice{
		country: country,
		city:    city,
		client:  customer,
		items:   items,
	}

}
