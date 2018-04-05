package core

/*
CoreConfiguration stores all configuration data required by the core app module (eg. http port, db creds).
*/
type CoreConfiguration struct {
	/*
		HTTPPort to know where to listen
	*/
	HTTPPort int16
}
