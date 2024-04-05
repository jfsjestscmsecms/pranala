package numfour

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Worker struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NumFour() {
	data := []Worker{}
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":

			rs, err := json.Marshal(&data)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}

			io.WriteString(w, string(rs))
			return
		case "POST":

			d := Worker{ID: r.PostFormValue("id"), Name: r.PostFormValue("name")}

			data = append(data, d)

			io.WriteString(w, "success added new data")
			return
		default:
			io.WriteString(w, "Bad Request")
			return
		}
	})

	log.Fatalln(http.ListenAndServe(":7777", nil))
}
