package main

import (

	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi" //En la version 1.4 de hyperledger tenia dos funciones init e invoke en el modulo gim. El modulo contractapi abstrae el modulo gim para tener esas funciones

)

// SmartContract provides functions for control the food
// Definimos una estructura
type SmartContract struct{
	contractapi.Contract
}

func main() {
	// Levantamos un nuevo Chaincode y le enviamos una nueva estructura SmartContract
	chaincode, err := contractapi.NewChaincode(new(SmartContract)) // Esto devuelve 2 valores: el Chaincode y el Error, y asignamos los resultados a las variables chaincode y err, en solo una linea de código
	
	if err != nil{ // Si hubo un error...
		fmt.Printf("Error creating foodcontrol chaincode: %s", err.Error()) // Usamos el método fmt para escribir en pantalla
		return
	}
	if err := chaincode.Start() ; err != nil{
		fmt.Printf("Error create foodcontrol chaincode: %s", err.Error())
	} // Levantamos el chaincode y verificamos que no haya errores
}

// En Golang no es necesario declarar una variable para asignarle un valor ni hace falta definir el tipo de dato.
// Las funciones pueden devolver más de un resultado