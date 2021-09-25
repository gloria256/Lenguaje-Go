package Bd_Neo4j

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Write_Bd(consulta string, data map[string]interface{}) (string, error) {

	/*

		Escribie datos en la base local Neo4j
		-------------------------------------

	*/

	var uri = "bolt://localhost:7687"
	var username = "neo4j"
	var password = "truora"
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		fmt.Println("1: ", err)
		return "", err
	}

	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(consulta, data)
		if err != nil {
			fmt.Println("2 ", err)
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		fmt.Println("3 ", err)
		return "", err
	}
	return greeting.(string), nil
}

func Read_Bd_map(consulta string, data map[string]interface{}) ([][]interface{}, error) {

	/*

		Lee datos presentes en la base local Neo4j
		------------------------------------------
		y retorna los resultados como: [][]interface{}

	*/

	var uri = "bolt://localhost:7687"
	var username = "neo4j"
	var password = "truora"
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		fmt.Println("1: ", err)
		return nil, err
	}

	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	persons, err := neo4j.Collect(session.Run(consulta, data))
	if err != nil {
		fmt.Println("2: ", err)
		return nil, err
	}

	var result [][]interface{}
	for _, person := range persons {
		result = append(result, person.Values)
	}
	return result, nil
}

func Read_Bd(consulta string) ([]string, error) {

	/*

		Lee datos presentes en la base local Neo4j
		------------------------------------------
		y retorna los resultados como: []string, es decir
		sólo usar cuando en la consulta se encuentra presente
		el RETURN DE 1 sola variable tipo string

	*/

	var uri = "bolt://localhost:7687"
	var username = "neo4j"
	var password = "truora"
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		fmt.Println("1: ", err)
		return nil, err
	}

	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	greeting, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string
		result, err := tx.Run(consulta, nil)
		if err != nil {
			fmt.Println("1: ", err)
			return nil, err
		}
		for result.Next() {
			list = append(list, result.Record().Values[0].(string))
		}

		if err = result.Err(); err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		fmt.Println("2: ", err)
		return nil, err
	}
	return greeting.([]string), nil
}

/*
https://cloud.google.com/appengine/docs/standard/go111/datastore/transactions
*/
