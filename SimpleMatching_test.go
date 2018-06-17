package goexpr

import (
	"testing"
)

func TestSimpleMatching(t *testing.T) {
	type args struct {
		rex   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantB   bool
		wantErr bool
	}{
		{"Test Simple1", args{"Hello", "Hello Regular Expression"}, true, false},
		{"Test Simple2", args{"Hxllo", "Hello Regular Expression"}, false, false},
		{"Test Simple3", args{`H\wllo`, "Hello Regular Expression"}, true, false},
		{"Test Simple4", args{`\d`, "This word don't has number"}, false, false},
		{"Test Simple5", args{`\d`, "This word has number 49"}, false, false},
		{"Test Simple6", args{`\s`, "ThisWordDoesNotHasSpace"}, false, false},
		{"Test Simple7", args{`\s`, "This Word has Space"}, true, false},
		{"Test Simple8", args{`\S`, "ThisWordDoesNotHasSpace"}, true, false},
		{"Test Simple9", args{`w.rd`, "any wxrd"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := CompileMatchString(tt.args.rex, tt.args.value)
			t.Log("gotB", gotB, " err:", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpleMatching() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("SimpleMatching() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestCompileMustMatchString(t *testing.T) {
	type args struct {
		rex   string
		value string
	}
	tests := []struct {
		name  string
		args  args
		wantB bool
	}{
		{"TestMust1", args{`Hello`, "Hello Regular Express"}, true},
		{"TestMust2", args{`Hxllo`, "Hello Regular Express"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args)
			if gotB := CompileMustMatchString(tt.args.rex, tt.args.value); gotB != tt.wantB {
				t.Errorf("CompileMustMatchString() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestCompilePosixMustMatchString(t *testing.T) {
	type args struct {
		rex   string
		value string
	}
	tests := []struct {
		name  string
		args  args
		wantB bool
	}{
		{"TestCompilePosix", args{`AABCDE{2}|ABCDE{4}`, "ABCDEEEEE"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := CompilePosixMustMatchString(tt.args.rex, tt.args.value); gotB != tt.wantB {
				t.Errorf("CompilePosixMustMatchString() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestCompileFindString(t *testing.T) {
	type args struct {
		rex   string
		value string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{"Test Simple1", args{"Hello", "Hello Regular Expression"}, "Hello", false},
		{"Test Simple2", args{"Hxllo", "Hello Regular Expression"}, "", false},
		{"Test Simple3", args{`H\wllo`, "Hello Regular Expression"}, "Hello", false},
		{"Test Simple4", args{`\d`, "This word don't has number"}, "", false},
		{"Test Simple5", args{`\d`, "This word has number 49"}, "4", false},
		{"Test Simple6", args{`\s`, "ThisWordDoesNotHasSpace"}, "", false},
		{"Test Simple7", args{`\s`, "This Word has Space"}, " ", false},
		{"Test Simple8", args{`\S`, "ThisWordDoesNotHasSpace"}, "T", false},
		{"Test Simple9", args{`w.rd`, "any wxrd"}, "wxrd", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := CompileFindString(tt.args.rex, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompileFindString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("CompileFindString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
