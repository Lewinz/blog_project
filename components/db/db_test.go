package db

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestRegister(t *testing.T) {
	type args struct {
		driver Driver
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.driver)
		})
	}
}
