/*
 * ecodata - OpenAPI 3.0
 *
 * API to store contamination events and data from different sources
 *
 * API version: 1.0.0
 * Contact: nestor.salvador@gdl.cinvestav.mx
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"flag"
	"log"

	ginsession "github.com/go-session/gin-session"

	"github.com/nestorneo/ecodata/config"
	"github.com/nestorneo/ecodata/middleware"
	sw "github.com/nestorneo/ecodata/models"
)

var (
	err         error
	userCfgFile string
	localConfig = config.DefaultConfig()
)

func init() {
	flag.StringVar(
		&userCfgFile, "config", "", "please provide server config")
}

func main() {
	log.Printf("Server started")
	flag.Parse()

	if userCfgFile != "" {
		log.Println("user provided config .... validating")
		localConfig, err = config.GetConfigFromFile(userCfgFile)
	}

	if err != nil {
		log.Panicln(err)
	}

	// middlewares
	// what is a middleware is an injector before reaching the actual endpoint it
	// pre-sets intended actions
	router := sw.NewRouter(
		ginsession.New(),
		middleware.GuidMiddleware(),
	)

	log.Fatal(router.Run(
		localConfig.Address))
}
