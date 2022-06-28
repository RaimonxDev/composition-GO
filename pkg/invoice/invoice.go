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
	// hace referencia al nuevo tipo que creamos
	items invoiceItem.Items
}

// Establecer el total  de una factura
// Como vamos a modificar el valor de una factura tenemos que usar *

func (i *Invoice) SetTotal() {
	// i.total hace referencia a valor total de la estructura de invoice, Porque i es el receptor de invoice
	// luego accedemos al campo items. y podemos acceder al nuevo metodo que creamos
	i.total = i.items.Total()
}

// cambiamos [] items por el nuevo tipo que creamos: Items
func New(country, city string, customer customer.Customer, items invoiceItem.Items) Invoice {

	return Invoice{
		country: country,
		city:    city,
		client:  customer,
		items:   items,
	}

}
