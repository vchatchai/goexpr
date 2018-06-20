package goexpr

import (
	"reflect"
	"testing"
)

func TestCompilereFindAllStringSubmatch(t *testing.T) {
	type args struct {
		reg   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]string
		wantErr bool
	}{

		{"Test1", args{"(.)at", "The cat sat on the mat."}, [][]string{{"cat"}, {"sat"}, {"mat"}}, false},
		{"Test2", args{".(at)", "The cat sat on the mat."}, [][]string{{"cat"}, {"sat"}, {"mat"}}, false},
		{"Test3", args{"(e)(.)", "Nobody expects the Spanish inquisition."}, [][]string{{"ex", "e", "x"}, {"ec", "e", "c"}, {"e ", "e", " "}}, false},
		{"Test3", args{`(expects (...) Spanish)`, "Nobody expects the Spanish inquisition."}, [][]string{{"exp", "e", "x", "p"}, {"ect", "e", "c", "t"}, {"e S", "e", " ", "S"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := CompilereFindAllStringSubmatch(tt.args.reg, tt.args.value)
			t.Log("gotRes", gotRes, "err", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompilereFindAllStringSubmatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CompilereFindAllStringSubmatch() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestCompileStringSubmatch(t *testing.T) {
	type args struct {
		reg   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
		wantErr bool
	}{
		{"Test1", args{`(Mr)(s)?\. (\w+) (\w+)`, "Mr. Leonard Spock"}, []string{"", ""}, false},
		{"Test3", args{`(expects (...) Spanish)`, "Nobody expects the Spanish inquisition."}, []string{"exp", "e", "x", "p"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := CompileStringSubmatch(tt.args.reg, tt.args.value)
			t.Log("gotRes", gotRes, "err", err)
			for _, s := range gotRes {
				t.Log(s)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("CompileStringSubmatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("CompileStringSubmatch() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
