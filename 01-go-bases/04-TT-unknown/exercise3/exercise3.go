package main

import (
	"errors"
	"fmt"
)

var (
	ErrClienteExistente = errors.New("el cliente ya existe")
	ErrValorCero        = errors.New("todos los datos del cliente deben tener un valor distinto de cero")
)

type Cliente struct {
	Legajo    int
	Nombre    string
	DNI       string
	Telefono  string
	Domicilio string
}

var clientesRegistrados = make(map[int]Cliente)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
			fmt.Println(r)
		}
		fmt.Println("Fin de la ejecución")
	}()

	nuevoCliente := Cliente{Legajo: 1, Nombre: "Juan", DNI: "30.215.321", Telefono: "4444-1234", Domicilio: "Calle Falsa 123"}
	err := verificarClienteExistente(nuevoCliente.Legajo)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	err = validarDatosCliente(nuevoCliente)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	clientesRegistrados[nuevoCliente.Legajo] = nuevoCliente
	fmt.Println("Cliente registrado con éxito:", nuevoCliente)

	err = verificarClienteExistente(nuevoCliente.Legajo)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func verificarClienteExistente(legajo int) error {
	if _, existe := clientesRegistrados[legajo]; existe {
		return ErrClienteExistente
	}
	return nil
}

func validarDatosCliente(cliente Cliente) error {
	if cliente.Legajo == 0 || cliente.Nombre == "" || cliente.DNI == "" || cliente.Telefono == "" || cliente.Domicilio == "" {
		return ErrValorCero
	}
	return nil
}
