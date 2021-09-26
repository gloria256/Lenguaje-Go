package main

import (
	"fmt"
	"test3/Bd_Neo4j"
	"test3/Data_Endpoint"
)

type NombreComprador struct {
	Id   string
	Name string
}

func main() {

	/*

		CONSULTA: LISTAR PERSONAS QUE HAN COMPRADO EN LA PLATAFROMA
		-----------------------------------------------------------

	*/

	var data = map[string]interface{}{}
	var NombreCompradores []NombreComprador
	var consulta = `MATCH (n:Transaccion),(c:Comprador)
	WHERE n.buyer_id=c.id RETURN DISTINCT c.name, c.id`

	list_name, err := Bd_Neo4j.Read_Bd_map(consulta, data)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("···", result)

	for _, row := range list_name {
		fmt.Println("....", row)
		NombreCompradores = append(NombreCompradores, NombreComprador{
			Id:   row[1].(string),
			Name: row[0].(string),
		})
	}

	/*

		ENDPOINT DE DATOS
		-----------------

	*/

	Data_Endpoint.Create_endpoint("/listaNombreCompradores", ":3000", NombreCompradores)
}

/*Visualizar data en: http://localhost:3000/listaNombreCompradores*/
