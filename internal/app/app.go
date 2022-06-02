package app

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"majezanu/capstone/config"
	"majezanu/capstone/infrastructure/router"
	"majezanu/capstone/registry"
)

func Run(conf *config.Config) {
	r := registry.NewRegistry(conf.CsvPath)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + conf.Port)
	if err := e.Start(":" + conf.Port); err != nil {
		log.Fatalln(err)
	}
}
