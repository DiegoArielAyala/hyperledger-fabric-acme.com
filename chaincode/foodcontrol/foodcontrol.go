package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi" //En la version 1.4 de hyperledger tenia dos funciones init e invoke en el modulo gim. El modulo contractapi abstrae el modulo gim para tener esas funciones
)

// SmartContract provides functions for control the food
// Definimos una estructura
type SmartContract struct {
	contractapi.Contract
}

// Definimos el activo con otra estructura que va a tener 2 propiedades
// Food describes basic details of what majes up a food
type Food struct {
	// Los atributos van con mayuscula
	Farmer  string `json:"farmer"`
	Variety string `json:"variety"`

	// Como este elemento lo voy a guardar en la blockchain com un elemento json, en la estructura le damos un nombre concreto para los campos del json
}

// Funcion para guardar en la blockchain
func (s *SmartContract) Set(ctx contractapi.TransactionContextInterface, foodId string, farmer string, variety string) error { // Indico que la funcion sea parte de la estructura de SmartContract. Defino propiedades que van a ser parte de la data que envia el cliente al momento de ejecutar el contrato. La funcion solo va a devolver un error

	// Normalmente se realizan validaciones de sintaxis (que los elementos que se reciben no sean nulos por ejemplo), de negocio, etc.
	// food, err := s.Query(ctx, foodId) //ctx es del contractapi
	// if food != nil { //Si food no es nil, el asset ya existe
	// 	fmt.Printf("foodId already existe error: %s", err.Error())
	// 	return err
	// }

	food := Food{ // Creo una estructura tipo Food
		Farmer:  farmer,
		Variety: variety,
	} // Esta estructura es un objeto binario que se puede transformar en json
	// Con contractapi tenemos un objeto llamado GetStub y una funcion PutState, que permite guardar en la blockchain, y necesita una Clave y un Valor

	foodAsBytes, _ := json.Marshal(food) // La barra baja _ es una variable no existente. Transformo el elemento food a bytes

	return ctx.GetStub().PutState(foodId, foodAsBytes)

	// Todas las transacciones que se creen en la blockchain tienen un ID y una estampa de tiempo (timeStampt) que se asignan por defecto
}

// Funcion de consulta de la blockchain
func (s *SmartContract) Query(ctx contractapi.TransactionContextInterface, foodId string) (*Food, error) { //En Go puedo devolver cuantos elementos quiera, en este caso *Food (hace referencia a la estructura Food) y el error
	foodAsBytes, err := ctx.GetStub().GetState(foodId) //Consultaremos el estado actual
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state, %s", err.Error())
	}
	if foodAsBytes == nil { //Valido que no sea nulo
		return nil, fmt.Errorf("%s does not exist", foodId)
	}

	food := new(Food)

	err = json.Unmarshal(foodAsBytes, food) // Hago un unmarshal de foodAsBytes y el resultadolo guardo en food
	if err != nil {
		return nil, fmt.Errorf("unmarshal error, %s", err.Error())
	}
	return food, nil
}

func main() {
	// Levantamos un nuevo Chaincode y le enviamos una nueva estructura SmartContract
	chaincode, err := contractapi.NewChaincode(new(SmartContract)) // Esto devuelve 2 valores: el Chaincode y el Error, y asignamos los resultados a las variables chaincode y err, en solo una linea de código

	if err != nil { // Si hubo un error...
		fmt.Printf("Error creating foodcontrol chaincode: %s", err.Error()) // Usamos el método fmt para escribir en pantalla
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error create foodcontrol chaincode: %s", err.Error())
	} // Levantamos el chaincode y verificamos que no haya errores
}

// En Golang no es necesario declarar una variable para asignarle un valor ni hace falta definir el tipo de dato.
// Las funciones pueden devolver más de un resultado
