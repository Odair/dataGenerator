package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var estados = [27]string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins", "Distrito Federal"}

type evento struct {
	id     string
	estado string
	valor  string
}

func main() {
	data := fillSlice()
	writeFile(data)
}

func fillSlice() []string {
	data := make([]string, 0, 190000)
	tamanho := 0
	for tamanho <= 190000 {
		ip := strconv.Itoa(rand.Intn(200)) + "." + strconv.Itoa(rand.Intn(200)) + "." + strconv.Itoa(rand.Intn(9)) + "." + strconv.Itoa(rand.Intn(99))
		estado := estados[rand.Intn(27)]
		valor := strconv.Itoa(rand.Intn(10))
		data = append(data, "{'ip':'"+ip+"','estado':'"+estado+"','valor':"+valor+"},")
		tamanho++
	}
	return data
}

func writeFile(data []string) {
	file, err := os.Create("result.json")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		v := []string{value}
		err := writer.Write(v)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
