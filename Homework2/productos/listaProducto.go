package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos
var lProductosMinimos listaProductos

const existenciaMinima int = 10

// Método de agregar producto con puntero hacia la listaProductos
func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	bandera := false
	for n := 0; n < len(lProductos); n++ {
		if nombre == (*l)[n].nombre {
			(*l)[n].cantidad += cantidad
			bandera = true
			if precio != (*l)[n].precio {
				(*l)[n].precio = precio
			}
		}
	}
	if !bandera {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}

}

// Función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista
func (l *listaProductos) incrementarProducto(nombre string, cantidad int) {
	for i := 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			(*l)[i].cantidad += cantidad
		}
	}
}

// Función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
func (l *listaProductos) buscarProducto(nombre string) (*producto, int) {
	var err = -1
	var p *producto = nil
	for i := 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			//& asigna la dirección
			p = &((*l)[i])
			err = 0
		}
	}
	//retorna la dirección del producto y si existe error o no
	return p, err
}

func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	prod, err := l.buscarProducto(nombre)
	if err == 0 && prod != nil {
		prod.cantidad = prod.cantidad - cantidad

		if prod.cantidad == 0 {
			l.eliminarProducto(nombre)
			fmt.Println("El producto: ", nombre, " se eliminó de la lista por cantidad inexistente")
		}
	}
}

func (l *listaProductos) eliminarProducto(nombre string) {
	for i, p := range *l {
		if p.nombre == nombre {
			// Eliminar el producto de la lista usando el índice i
			*l = append((*l)[:i], (*l)[i+1:]...)
			break
		}
	}
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	for _, p := range *l {
		if p.cantidad < 10 {
			lProductosMinimos = append(lProductosMinimos, p)
		}
	}
	return lProductosMinimos
}

func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos) listaProductos {
	for i := 0; i < len(listaMinimos); i++ {
		listaMinimos[i].cantidad = existenciaMinima
	}
	return listaMinimos
}

func ordenarListaProductos(p listaProductos, keyFunc func(p *producto) int, forma bool) {
	funcOrden := func(i, j int) bool {
		if forma {
			return keyFunc(&p[i]) < keyFunc(&p[j])
		}
		return keyFunc(&p[i]) > keyFunc(&p[j])
	}

	sort.Slice(p, funcOrden)
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("Pan", 30, 7000)
	lProductos.agregarProducto("Tomate", 2, 2000)
	lProductos.agregarProducto("frijoles", 30, 2000)
	lProductos.agregarProducto("leche", 10, 1200)
	lProductos.agregarProducto("café", 10, 4500)
	lProductos.agregarProducto("café", 15, 5000)
	lProductos.incrementarProducto("arroz", 30)

}

func main() {
	llenarDatos()
	fmt.Println("\nlista original: ", lProductos)
	//venda productos
	lProductos.venderProducto("frijoles", 25)
	fmt.Println("Lista actualizada: ", lProductos)
	lProductos.listarProductosMínimos()
	fmt.Println("Lista de mínimos: ", lProductosMinimos)
	lProductos.aumentarInventarioDeMinimos(lProductosMinimos)
	fmt.Println("Lista de mínimos actualizada: ", lProductosMinimos)
	ordenarListaProductos(lProductos, func(p *producto) int { return p.precio }, true)
	fmt.Println("Lista ordenada de forma ascendente por el valor de precio: ", lProductos)
}
