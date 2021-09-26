package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"test3/Bd_Neo4j"
)

type Transaccion struct {
	Id          string
	Buyer_id    string
	Ip          string
	Device      string
	Product_ids string
}

func main() {

	/*

		LECTURA DE ENDPOINT LOCAL TRANSACCIONES
		---------------------------------------

	*/

	var name = "../data/transactions"
	consulta := "CREATE (tra:Transaccion {id:$Id,buyer_id:$Buyer_id,ip:$Ip,device:$Device,product_ids:$Product_id}) RETURN tra.id"
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	arr_tran := strings.Split(strings.Replace(string(contents), "\n", "", -1), "#")
	for i := 1; i < len(arr_tran); i++ {
		if i == len(arr_tran) {
			break
		}

		ipRegexp := regexp.MustCompile("[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+") //\\s [a-z]")
		expresion := ipRegexp.FindString(arr_tran[i])
		indIni := strings.Index(arr_tran[i], expresion)
		indFin := len(expresion) + indIni
		s := strings.Index(arr_tran[i], "(")
		s1 := strings.Index(arr_tran[i], ")")

		var id = arr_tran[i][0:12]
		var buyer_id = arr_tran[i][13 : indIni-1]
		var ip = arr_tran[i][indIni:indFin]
		var device = arr_tran[i][indFin+1 : s-1]
		var product_ids = arr_tran[i][s+1 : s1]

		prolis_id := strings.Split(product_ids, ",")

		/*

			SAVE IN BD NEO4J
			----------------

		*/

		for _, Product_id := range prolis_id {
			var data = map[string]interface{}{"Id": id, "Buyer_id": buyer_id, "Ip": ip, "Device": device, "Product_id": Product_id}
			result, _ := Bd_Neo4j.Write_Bd(consulta, data)
			fmt.Println("···", result)
		}
	}

}
