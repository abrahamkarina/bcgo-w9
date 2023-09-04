package main

import (
	"errors"
	"testing"
)

func TestVerificarClienteExistente(t *testing.T) {
	clientesRegistrados = make(map[int]Cliente)
	err := verificarClienteExistente(1)
	if err != nil {
		t.Errorf("Esperado nil, pero obtuve un error: %v", err)
	}

	clienteExistente := Cliente{Legajo: 1, Nombre: "Carmen", DNI: "30.112.111", Telefono: "1234-1234", Domicilio: "Calle A"}
	clientesRegistrados[clienteExistente.Legajo] = clienteExistente
	err = verificarClienteExistente(clienteExistente.Legajo)
	if err == nil || !errors.Is(err, ErrClienteExistente) {
		t.Errorf("Esperado un error con mensaje específico, pero obtuve: %v", err)
	}
}

func TestValidarDatosCliente(t *testing.T) {
	clienteValido := Cliente{Legajo: 1, Nombre: "Juan", DNI: "12345678", Telefono: "555-1234", Domicilio: "Calle A"}
	err := validarDatosCliente(clienteValido)
	if err != nil {
		t.Errorf("Esperado nil, pero obtuve un error: %v", err)
	}

	clienteInvalido := Cliente{Legajo: 2, Nombre: "", DNI: "87654321", Telefono: "555-5678", Domicilio: "Calle B"}
	err = validarDatosCliente(clienteInvalido)
	if err == nil || !errors.Is(err, ErrValorCero) {
		t.Errorf("Esperado un error con mensaje específico, pero obtuve: %v", err)
	}

	clienteInvalido.Legajo = 0
	err = validarDatosCliente(clienteInvalido)
	if err == nil || !errors.Is(err, ErrValorCero) {
		t.Errorf("Esperado un error con mensaje específico, pero obtuve: %v", err)
	}
}
