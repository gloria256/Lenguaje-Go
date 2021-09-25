package Data_Endpoint

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Create_endpoint(nameUrl string, port string, data interface{}) {

	/*

		Cear endpoint con datos, usando Chi
		-----------------------------------

	*/

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get(nameUrl, func(response http.ResponseWriter, r *http.Request) {
		response.Header().Set("Content-type", "application/json")
		response.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		json.NewEncoder(response).Encode(data)
	})
	http.ListenAndServe(port, r) //":3000"
}

func Read_enpoint_json(url string) ([]byte, error) {

	/*

		Lee un endpoint de datos [{},{}] .. (.JSON)
		y retorna un arreglo de bytes
		-------------------------------------------

	*/

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	response, err := client.Do(request) //hace petición
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer response.Body.Close() //cerrar el proceso
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, err
}

func Read_endpoint_csv(url string) ([][]string, error) {

	/*

		Lee un endpoint de datos .csv y retorna un arreglo
		de listas de strings
		--------------------------------------------------

	*/

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	response, err := client.Do(request) //hace petición
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close() //cerrar el proceso

	reader := csv.NewReader(response.Body)
	reader.Comma = '\''
	res, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Read_local_file(name string) ([]string, error) {

	/*

		Lee un endpoint de datos local y retorna un arreglo
		de strings
		---------------------------------------------------

	*/

	contents, err := ioutil.ReadFile(name)
	if err == nil {
		result := strings.Replace(string(contents), " ", "", -1)
		result1 := strings.Split(result, "#")
		return result1, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}
