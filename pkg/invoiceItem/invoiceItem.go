package invoiceItem

type Item struct {
	id      uint
	product string
	value   float64
}

// Funcion construtora para crear un nuevo item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Funcion constructora para crear un nuevo slices de items
// Usamos una funcion variatica
func NewItems(items ...Item) Items {
	//creamos una variable para almacenar nuestro slices
	var is Items
	// for range para recorrer los nuevos items
	for _, item := range items {
		is = append(is, item)
	}
	return is
}

/*Al crear el nuevo tipo ya no es necesario este getter. Porque podemos acceder al valor value en el mismo paquete
y no hay necesidad de ser exportado

func (i *Item) Value() float64 {
	return i.value
}
*/

// creando un nuevo tipo a base de un slices de []Item
type Items []Item

// Hacemos referencia al nuevo tipo Items que creamos
func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}
	return total
}
