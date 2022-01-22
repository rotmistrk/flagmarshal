package flagmarshal

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

type testStruct struct {
	Username string  `flag:"u" help:"user name"`
	Password string  `flag:"p" help:"user password"`
	IntVal   int     `flag:"i" help:"integer"`
	BoolVal  bool    `flag:"b" help:"boolean"`
	FloatVal float64 `flag:"f"`
}

func TestParseFlags(t *testing.T) {
	type args struct {
		target *testStruct
		args   []string
	}
	tests := []struct {
		name    string
		args    args
		parsed  interface{}
		want    []string
		wantErr bool
	}{
		//{"fails for slice", args{target: make([]string, 0)}, nil, nil, true},
		{"works for struct", args{target: &testStruct{}}, nil, make([]string, 0), false},
		{
			"string works",
			args{
				args:   []string{"cmd", "-u", "aUser", "-i", "0x10", "-b", "true", "-f", "0.345", "hello"},
				target: &testStruct{},
			},
			testStruct{Username: "aUser", IntVal: 16, BoolVal: true, FloatVal: 0.345},
			[]string{"hello"},
			false,
		},
	}
	saveArgs := os.Args
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if tt.args.args != nil {
				os.Args = tt.args.args
			} else {
				os.Args = []string{"cmd"}
			}
			got, err := ParseFlags(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFlags() got = %v, want %v", got, tt.want)
			}
			if tt.parsed != nil && !reflect.DeepEqual(tt.parsed, *tt.args.target) {
				t.Errorf("ParseFlags() result = %v, expected %v", tt.args.target, tt.parsed)
			}
		})
	}
	os.Args = saveArgs
}
