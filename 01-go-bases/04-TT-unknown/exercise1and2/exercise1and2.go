package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	defer fmt.Println("Ejecución finalizada")
	fileName, err := filepath.Abs("01-go-bases/04-TT-unknown/customers.txt")
	if err != nil {
		panic(err)
	}

	err = readFile(fileName)
	if err != nil {
		//	fmt.Println("Error al leer el archivo:", err)
		panic(err)
	}

}

func readFile(fileName string) error {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		//return err
		return errors.New("“el archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println("Leyendo datos del archivo:")
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(string(buffer[:n]))
	}
	return nil
}
