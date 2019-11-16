package golua

import (
	"testing"
)

func Test_newLuaFromFile(t *testing.T) {
	type args struct {
		fileName string
		libs     []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// Error Case
		{"Test1", args{"sakfj jhiuq3", []string{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just Check to cover error case
			defer func() {
				if r := recover(); r == nil {
					t.Error("Should Have Panicked!!")
				}
			}()
			newLuaFromFile(tt.args.fileName, tt.args.libs...)
		})
	}
}
