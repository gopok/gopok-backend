package core

/*
Controller allows for routes to be mounted.
*/
type Controller interface {
	Register(app *Application)
}
