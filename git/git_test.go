package git

import (
	"testing"
)

func TestInstallPreCommitHook(t *testing.T) {
	err := InstallPreCommitHook(`
		#!/bin/bash
		./sanity.sh`, "../.git")

	if err != nil {
		t.Fatal(err)
	}
}
