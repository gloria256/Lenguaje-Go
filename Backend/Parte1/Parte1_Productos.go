package main

import (
	"fmt"
	"strconv"
	"test3/Bd_Neo4j"
	"test3/Data_Endpoint"
)

type Producto struct {
	Id    string
	Name  string
	Price int
}

func main() {

	/*

		LECTURA DE ENDPOINT PRODUCTOS Y SAVE IN BD
		------------------------------------------

	*/

	var Productos [62]Producto // arreglo variable de tipo comprador struct (lista de estructuras)
	var endpoint_productos = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
	var consulta = "CREATE (pro:Producto {id:$Id,name:$Name,price:$Price}) RETURN pro.id"

	arr_pro, err := Data_Endpoint.Read_endpoint_csv(endpoint_productos)
	if err != nil {
		fmt.Println(err)
	}

	for idx, row := range arr_pro {
		var id = row[0]
		var name = row[1]
		var price int

		price, _ = strconv.Atoi(row[2])
		if idx == 62 {
			break
		}
		Productos[idx] = Producto{
			Id:    id,
			Name:  name,
			Price: price,
		}

		/*

			SAVE IN BD NEO4J
			----------------

		*/

		var data = map[string]interface{}{"Id": id, "Name": name, "Price": price}
		result, _ := Bd_Neo4j.Write_Bd(consulta, data)
		fmt.Println("***", result)
	}
	fmt.Println("productos", Productos)
}
