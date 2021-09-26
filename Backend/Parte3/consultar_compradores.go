package main

import (
	"fmt"
	"test3/Bd_Neo4j"
	"test3/Data_Endpoint"
)

type Historial_Compra struct {
	Name           string
	Id_Producto    string
	Name_Producto  string
	Price_Producto int64
}
type Buyer_same_ip struct {
	Ip   string
	Name string
	Dev  string
}
type Recomendacion_Producto struct {
	Name string
}

func main() {

	/*
		NOTA:
		-----
		Se demora en ejecutarse Aprox 2 minutos
	*/

	var data = map[string]interface{}{}
	var consulta = `MATCH (t:Transaccion),(c:Comprador)
	WHERE t.buyer_id=c.id RETURN DISTINCT c.name, c.id`

	list_name, err := Bd_Neo4j.Read_Bd_map(consulta, data)
	if err != nil {
		fmt.Println(err)
	}
	var dataId = map[string]interface{}{}
	for _, row := range list_name {
		var buyer_id = row[1].(string)
		resultado := consultar_comprador(buyer_id)
		dataId[buyer_id] = resultado
	}
	Data_Endpoint.Create_endpoint("/ConsultarCompradores", ":1000", dataId)

}

func consultar_comprador(buyer_id string) []interface{} {

	/*

		LISTAR PERSONAS QUE HAN COMPRADO EN LA PLATAFROMA
		--------------------------------------------------

	*/

	var data = map[string]interface{}{"Buyer_id": buyer_id}
	var Historial_Compras []Historial_Compra
	var consulta = `MATCH (n:Comprador),(t:Transaccion{buyer_id:$Buyer_id}),(p:Producto)
	WHERE n.id=t.buyer_id AND t.product_ids=p.id 
	RETURN n.name,p.id,p.name,p.price 
	ORDER BY t.buyer_id`

	list_name, err := Bd_Neo4j.Read_Bd_map(consulta, data)
	if err != nil {
		fmt.Println(err)
	}

	for _, data := range list_name {
		Historial_Compras = append(Historial_Compras, Historial_Compra{
			Name:           data[0].(string),
			Id_Producto:    data[1].(string),
			Name_Producto:  data[2].(string),
			Price_Producto: data[3].(int64),
		})
	}

	/*

		OTROS COMPRADORES USANDO LA MISMA IP
		------------------------------------

	*/

	var Buyers_same_ip []Buyer_same_ip
	consulta = `MATCH (n:Comprador),(t:Transaccion{buyer_id:$Buyer_id})
	WHERE n.id=t.buyer_id
	WITH n.name AS name, COLLECT(DISTINCT t.ip) AS ip
	CALL{
		WITH name, ip
		MATCH  (n:Comprador),(t:Transaccion)
		WHERE n.id=t.buyer_id AND t.ip IN ip AND n.name <> name
		//WITH t.ip AS ip1,t.device AS dev, n.name AS name
		RETURN  DISTINCT n.name AS name1,t.ip AS ip1,t.device AS dev
	}
	RETURN name1,ip1,dev`

	list_name, err = Bd_Neo4j.Read_Bd_map(consulta, data)
	if err != nil {
		fmt.Println(err)
	}

	for _, data := range list_name {
		Buyers_same_ip = append(Buyers_same_ip, Buyer_same_ip{
			Ip:   data[1].(string),
			Name: data[0].(string),
			Dev:  data[2].(string),
		})
	}

	/*

		RECOMENDACIONES OTROS PRODUCTOS
		--------------------------------

	*/

	var Recomendacion_Productos []Recomendacion_Producto
	consulta = `MATCH (n:Comprador),(t:Transaccion{buyer_id:"228a73a2"}),(p:Producto)
	WHERE n.id=t.buyer_id AND t.product_ids=p.id
	WITH COLLECT(DISTINCT p.name) AS pro
	CALL{
		WITH pro
		MATCH  (n:Comprador),(t:Transaccion),(p:Producto)
		WHERE n.id=t.buyer_id AND t.product_ids=p.id AND p.name IN
		pro
		RETURN  DISTINCT p.name AS name
	}
	RETURN name
	LIMIT 3`

	list_name, err = Bd_Neo4j.Read_Bd_map(consulta, data)
	if err != nil {
		fmt.Println(err)
	}

	for _, data := range list_name {
		Recomendacion_Productos = append(Recomendacion_Productos, Recomendacion_Producto{
			Name: data[0].(string),
		})
	}

	/*

		SALIDA
		----------------

	*/

	var result = []interface{}{Historial_Compras, Buyers_same_ip, Recomendacion_Productos}
	return result

}
