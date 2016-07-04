package env

import (
	"os"
	"testing"
)

func TestEnvService_Environment(t *testing.T) {
	// Defaults to dev
	os.Setenv(EnvironmentVariable, "Unknown")
	if env := Environment(); env != Dev {
		t.Fatalf("Unexpected default environment: %v", env)
	}

	// Dev
	os.Setenv(EnvironmentVariable, Dev.ID)
	if env := Environment(); env != Dev {
		t.Fatalf("Unexpected DEV environment: %v", env)
	}

	// Test
	os.Setenv(EnvironmentVariable, Test.ID)
	if env := Environment(); env != Test {
		t.Fatalf("Unexpected TEST environment: %v", env)
	}

	// Prod
	os.Setenv(EnvironmentVariable, Prod.ID)
	if env := Environment(); env != Prod {
		t.Fatalf("Unexpected Prod environment: %v", env)
	}
}
