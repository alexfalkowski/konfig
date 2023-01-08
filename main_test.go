//go:build features

package main

import (
	"testing"
)

func TestFeatures(t *testing.T) {
	if err := command().Run(); err != nil {
		t.Fatal(err.Error())
	}
}
