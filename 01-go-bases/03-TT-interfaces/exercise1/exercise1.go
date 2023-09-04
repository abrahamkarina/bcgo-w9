package main

import "fmt"

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (e Estudiante) Detalle() {
	fmt.Println(e)
}

func Inscribir(Nombre, Apellido, Dni, Fecha string) Estudiante {
	return Estudiante{
		Nombre:   Nombre,
		Apellido: Apellido,
		DNI:      Dni,
		Fecha:    Fecha,
	}
}

func main() {
	nuevoEstudiante := Inscribir("Juan", "Perez", "40.502.256", "21/05/2005")
	nuevoEstudiante.Detalle()
}
