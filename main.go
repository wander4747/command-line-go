package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()

	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
	/*erro := application.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}*/

}
