package main

type ErrMinimoImponible struct {
	msg string
}

func NewMinimoNoImponibleError() *ErrMinimoImponible {
	return &ErrMinimoImponible{msg: "el salario ingresado no alcanza el minimo imponible"}
}
func (e *ErrMinimoImponible) Error() string {
	return e.msg
}
