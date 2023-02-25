package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"solution/pkg/model"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/yaakovlew/Desktop/WB/L0/solution/tmp/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
	r.ParseForm()
	if len(r.Form["id"]) != 0 {
		idStr := r.Form["id"][0]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		element, exist := model.AllData.Load(id)
		if exist {
			fmt.Fprint(w, element)
		}
	}
}

func Handle() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
