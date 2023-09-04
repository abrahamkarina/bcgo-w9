package main

import "fmt"

/*Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y
los adicionales en caso que los tenga
*/

type Producto interface {
	Precio() float64
}
type producto struct {
	precio float64
}
type Pequeño struct {
	producto
}
type Mediano struct {
	producto
	porcentajeMantenimiento float64
}

type Grande struct {
	producto
	porcentajeMantenimiento float64
	costoEnvio              float64
}

func (p producto) Precio() float64 {
	return p.precio
}
func (p Mediano) Precio() float64 {
	return p.precio + p.precio*p.porcentajeMantenimiento
}
func (p Grande) Precio() float64 {
	return p.precio + p.precio*p.porcentajeMantenimiento + p.costoEnvio
}

const (
	typePequeno = iota
	typeMediano
	typeGrande
)

func Factory(productType int, precio float64) Producto {
	switch productType {
	case typePequeno:
		return newPequeno(precio)
	case typeMediano:
		return newMediano(precio)
	case typeGrande:
		return newGrande(precio)
	default:
		return &producto{precio: precio}
	}

}

func newGrande(precio float64) Producto {
	const (
		porcMantenimiento = 0.06
		costoEnvio        = 2500
	)
	product := new(Grande)
	product.precio = precio
	product.porcentajeMantenimiento = porcMantenimiento
	product.costoEnvio = costoEnvio
	return product
}

func newMediano(precio float64) Producto {
	const (
		porcMantenimiento = 0.03
	)
	product := new(Mediano)
	product.precio = precio
	product.porcentajeMantenimiento = porcMantenimiento
	return product
}

func newPequeno(precio float64) Producto {
	product := new(Pequeño)
	product.precio = precio
	return product
}

func main() {
	p := Factory(typePequeno, 200)
	m := Factory(typeMediano, 1000)
	g := Factory(typeGrande, 1000)

	fmt.Println(p.Precio())
	fmt.Println(m.Precio())
	fmt.Println(g.Precio())
}
