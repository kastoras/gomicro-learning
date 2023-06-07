package main

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	repo := data.NewPosgresTestRepository(nil)
	testApp.Repo = repo
	os.Exit(m.Run())
}