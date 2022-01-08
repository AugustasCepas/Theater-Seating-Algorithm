package main

import (
	"fmt"

	"github.com/AugustasCepas/Theater-Seating-Algorithm/api"
	"github.com/AugustasCepas/Theater-Seating-Algorithm/configs"
	_ "github.com/lib/pq"
)

func main() {

	a := api.App{}
	var err error

	a.DB, err = configs.GetDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	a.Initialize(a.DB)

	a.Run(":8000")

}
