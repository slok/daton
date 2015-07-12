package configuration

// Configuration defaults
var (

	// Config configuration :P
	ConfigName  = "daton"
	ConfigType  = "json"
	ConfigPaths = []string{
		"/etc",
		"$HOME/.config",
		"./",
	}

	// Server configuration
	Port       = 3000
	ApiVersion = 1

	// Database configuration
	VoltdbName = ""

	// App configuration
	EnableAutomerge = false
	Debug           = true

	// Deploy configuration
	DefaultEnvironment = "production"
	DefaultTask        = "deploy"
)
