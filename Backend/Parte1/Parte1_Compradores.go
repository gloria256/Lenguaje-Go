package main

import (
	"encoding/json"
	"fmt"
	"test3/Bd_Neo4j"
	"test3/Data_Endpoint"
)

type Comprador struct {
	Id   string
	Name string
	Age  int
}

func main() {

	/*

		LECTURA DE ENDPOINT COMPRADORES
		-------------------------------
	*/

	var Compradores []Comprador
	var endpoint_buyers = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
	var consulta = "CREATE (comp:Comprador {id:$Id,name:$Name,age:$Age}) RETURN comp.id"

	arr_bytes, err := Data_Endpoint.Read_enpoint_json(endpoint_buyers)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(arr_bytes, &Compradores) // llega arreglo de bytes y la interfaz es una estructura
	fmt.Println(Compradores)

	/*

		SAVE IN BD NEO4J
		----------------
	*/

	for _, row := range Compradores {
		var data = map[string]interface{}{"Id": row.Id, "Name": row.Name, "Age": row.Age}
		result, _ := Bd_Neo4j.Write_Bd(consulta, data)
		fmt.Println("-->", result)

	}
}
