package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amsipe/nr-span-example/mysql"
	"github.com/amsipe/nr-span-example/server"
	"github.com/jmoiron/sqlx"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	config, err := getEnvConfig()
	if err != nil {
		log.Fatalf("error getting config env: %v", err)
		return
	}

	nrApp := setupNewRelic()

	db, cleanup, err := setupMySQL(config)
	defer cleanup()

	store := mysql.NewStore(db)

	s := server.New(store)

	r := newRouter(&routerConfig{
		Server:   s,
		NewRelic: nrApp,
	})

	log.Printf("Listening on port :%s\n", config.HTTPPort)
	http.ListenAndServe(fmt.Sprintf(":%s", config.HTTPPort), r)
}

func setupNewRelic() *newrelic.Application {
	appName := "asipe-make-spans"

	app, err := newrelic.NewApplication(
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigAppName(appName),
	)
	if err != nil {
		log.Printf("Error connecting to new relic: %v", err)
	}

	return app
}

func setupMySQL(config *envConfig) (*sqlx.DB, func(), error) {
	db, err := sqlx.Connect(
		"nrmysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/?parseTime=true&timeout=15s&readTimeout=60s",
			config.DBUsername,
			config.DBPassword,
			config.DBHost,
			config.DBPort,
		),
	)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil

}
