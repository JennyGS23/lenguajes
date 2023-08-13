package main

import "fmt"

// estructura de calzados
type calzado struct {
	marca    string
	precio   int
	talla    int //hacer la restriccion de las tallas
	cantidad int
}

// arreglo/slice
type listaZapatos []calzado

var lZapatos listaZapatos

// función para validación y agregado de calzados (validando que la talla esté dentro del rango)
func (l *listaZapatos) agregarCalzado(marca string, precio int, talla int, cantidad int) {
	bandera := false
	for i := 0; i < len(lZapatos); i++ {
		if marca == (*l)[i].marca && talla >= 34 && talla <= 44 {
			cantidad += (*l)[i].cantidad
			precio = (*l)[i].precio
			bandera = true
		}

	}
	if !bandera && talla >= 34 && talla <= 44 {
		*l = append(*l, calzado{marca: marca, precio: precio, talla: talla, cantidad: cantidad})
	} else {
		print("El nuevo artículo de la marca: ", marca, ", no se registró,"+
			" porque no cumple con el rango de talla establecido\n")
	}
}

// función para llamar a agregarZapatos
func ingresarDatosCalzado() {
	lZapatos.agregarCalzado("Nike", 50000, 35, 2)
	lZapatos.agregarCalzado("Puma", 60000, 39, 1)
	lZapatos.agregarCalzado("Adidas", 40000, 40, 2)
	lZapatos.agregarCalzado("Vans", 35000, 41, 2)
	lZapatos.agregarCalzado("Converse", 38000, 42, 15)
	lZapatos.agregarCalzado("Timberland", 60000, 38, 7)
	lZapatos.agregarCalzado("Reebok", 50000, 35, 4)
	lZapatos.agregarCalzado("Fila", 60000, 39, 9)
	lZapatos.agregarCalzado("Ecco", 50000, 40, 7)
	lZapatos.agregarCalzado("Crocs", 10000, 50, 1)

}

// funcion para buscar el calzado por su marca, retornando si existe con una cantidad mayor a 0
func (l *listaZapatos) buscarCalzado(marca string) (*calzado, int) {
	var mensaje = -1
	var c *calzado = nil

	for i := 0; i < len(*l); i++ {
		if (*l)[i].marca == marca && (*l)[i].cantidad > 0 {
			c = &((*l)[i])
			mensaje = 0
		}
	}
	return c, mensaje
}

// función para vender el calzado en caso de que su existencia sea mayor a 0. En caso contrario se emite un mensaje
func (l *listaZapatos) venderCalzado(marca string) {
	var calz, mensaje = l.buscarCalzado(marca)
	if mensaje != -1 {
		(*calz).cantidad = (*calz).cantidad - 1
		print("\nVenta del calzado de marca: ", marca, ", realizada con éxito. ")
		print("La cantidad actual es de: ", (*calz).cantidad, "\n")
	} else {
		print("\nEste calzado de marca: ", marca+
			" no se puede vender, se registran 0 cantidades de existencia\n")

	}
}

// función menú
func main() {
	ingresarDatosCalzado()
	print("Lista de calzados ingresados: ")
	fmt.Println(lZapatos)

	lZapatos.venderCalzado("Puma")
	lZapatos.venderCalzado("Puma")
	lZapatos.venderCalzado("Nike")

	print("\nLa lista actualizada de calzados cumple el siguiente orden: Marca, precio, talla y cantidad \n")
	fmt.Println(lZapatos)

}
