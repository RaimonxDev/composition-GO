package main

import (
	"fmt"
	"github.com/RaimonxDev/composition-GO/pkg/customer"
	"github.com/RaimonxDev/composition-GO/pkg/invoice"
	"github.com/RaimonxDev/composition-GO/pkg/invoiceItem"
)

func main() {

	i := invoice.New(
		"chile", "osorno",
		//usamos la funcion constructor
		customer.New("Ramon", "arturo pratt", "9815144452"),
		invoiceItem.NewItems(
			// Usamos la funcion constructora
			invoiceItem.New(1, "Curso de go", 60.50),
			invoiceItem.New(2, "Curso de avanzado con go", 70.34),
			invoiceItem.New(3, "testing de go", 50.34),
		),
	)

	// Calculamos el total de la factura

	i.SetTotal()

	fmt.Printf("%+v", i)
}
