package main

import (
	"fmt"
	"to-do-app/app/controllers"
	"to-do-app/app/models"
)

func main() {
	

	fmt.Println(models.Db)

	controllers.StarMainSerever()
}
