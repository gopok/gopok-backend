package core

import (
	"container/list"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

/*
ControllersToRegister lists all controllers to be automatically registered at startup.
*/
var ControllersToRegister = list.New()

/*
Application is the main struct used to start the app and load all other parts.
*/
type Application struct {
	Config *CoreConfiguration
	Router *mux.Router
	Db     *mgo.Session
}

/*
Run starts the application.
*/
func (app *Application) Run() {
	log.Info("Starting application...")
	log.Info("Loading core configuration")
	configErr := app.loadConfig()
	if configErr != nil {
		log.Fatal("Failed to load core configuration: ", configErr)
		return
	}
	log.Info("Loaded core configuration")
	log.Info("Initializing router")
	app.initRouter()

	// register controllers

	for c := ControllersToRegister.Front(); c != nil; c = c.Next() {
		c.Value.(Controller).Register(app)
		log.Printf("Registered controller %s", reflect.TypeOf(c.Value).String())
	}
	app.Router.PathPrefix("/").HandlerFunc(app.notFoundHandler)
	app.Router.Use(loggingMiddleware)
	var dbError error
	log.Info("Connecting to database")
	app.Db, dbError = mgo.Dial(app.Config.MongoURL)
	if dbError != nil {
		log.Fatal("Failed to connect to database: ", dbError)
		return
	}

	log.Info("Starting to listen")
	httpErr := app.initHTTP()

	if httpErr != nil {
		log.Fatal("Failed to bind http listener: ", configErr)
		return
	}

}

func (app *Application) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(404)
	jsonData, _ := json.Marshal(map[string]interface{}{"message": "Not found", "code": 404})
	w.Write(jsonData)
}

func (app *Application) loadConfig() error {
	rawConfig, readErr := ioutil.ReadFile("config/core.json")
	if readErr != nil {
		return errors.Wrap(readErr, "failed to read config/core.json")
	}
	parseErr := json.Unmarshal(rawConfig, &app.Config)
	if parseErr != nil {
		return errors.Wrap(parseErr, "failed to parse config/core.json")
	}
	return nil
}

func (app *Application) initRouter() {
	app.Router = mux.NewRouter()
}
func (app *Application) initHTTP() error {
	err := http.ListenAndServe(":"+strconv.Itoa(int(app.Config.HTTPPort)), app.Router)
	if err != nil {
		return errors.Wrap(err, "failed to listen http")
	}

	return nil
}
