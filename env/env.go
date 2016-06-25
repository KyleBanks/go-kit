// Package env provides application environment detection, and support for a Dev/Test/Prod environment system.
package env

import "os"

const (
	EnvironmentVariable = "GO_ENV"
)

var (
	Dev  = environment{"DEV"}
	Test = environment{"TEST"}
	Prod = environment{"PROD"}
)

type environment struct {
	ID string
}

// Environment returns the current environment that the go application is running in,
// based on the environment variable. If no environment variable is found, or
// it is not one of Dev/Test/Prod, the default (Dev) will be returned.
func Environment() environment {
	envVar := os.Getenv(EnvironmentVariable)
	if len(envVar) > 0 {
		switch envVar {
		case Test.ID:
			return Test
		case Prod.ID:
			return Prod
		}
	}

	return Dev
}
