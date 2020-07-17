package main

import (
    "log"  
    "net/http"
    "goCRUD/config"
    "goCRUD/models/employees"
)
 
func main() {
    http.HandleFunc("/", employees.Index)
    http.HandleFunc("/delete", employees.Delete)
    http.HandleFunc("/new", employees.New)
    http.HandleFunc("/insert", employees.Insert)
    http.HandleFunc("/edit", employees.Edit)
    http.HandleFunc("/show", employees.Show)
    http.HandleFunc("/update", employees.Update)

    env := config.InitEnv()
    log.Println("Server started on: http://localhost:" + env.Port)
    http.ListenAndServe(":" + env.Port, nil)
}
