package numthree

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint   `json:"age"`
}

const file_path = "testdata/test_data.json"

func NumberThree() {
	data, err := os.ReadFile(file_path)
	if err != nil {
		log.Fatalln(err)
	}

	var p Person

	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Fatalln(err)
	}

	p.Email = "johndoe@example.com"
	p.Age = 1

	o, err := json.Marshal(&p)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.OpenFile(file_path, os.O_RDWR|os.O_TRUNC, 0555)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	_, err = file.Write(o)
	if err != nil {
		log.Fatalln(err)
	}

}
