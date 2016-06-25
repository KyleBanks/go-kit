package git

import "testing"

func TestInstallPreCommitHook(t *testing.T) {
	if err := InstallPreCommitHook(`
		#!/bin/bash
		./go-test.sh
	`, "../.git"); err != nil {
		t.Fatal(err)
	}
}
