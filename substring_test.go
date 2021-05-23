package main

import (
	"testing"
)

func TestExtendedString_IsSubStringOf(t *testing.T) {
	type args struct {
		longer string
	}
	tests := []struct {
		name string
		s    ExtendedString
		args args
		want bool
	}{
		{"test name", ExtendedString("small string"), args{"is a substring of not so small string"}, true},
		{"hello world", ExtendedString("hello"), args{"hello world"}, true},
		{"world hello", ExtendedString("world"), args{"hello world"}, true},
		{"hello world", ExtendedString("world"), args{"hello world hello"}, true},
		{"world hell", ExtendedString("world"), args{"hello worl"}, false},
		{"worldofwarcraft", ExtendedString("worldofwarcraft"), args{"world"}, false},
		{"worldofwarcraft returns", ExtendedString("world"), args{"worldofwarcraft"}, true},
		{"accents work fine", ExtendedString("évidemment"), args{"évidemment"}, true},
		{"accents work", ExtendedString("évidemment"), args{"evidemment"}, false},
		{"hellohello", ExtendedString("hello"), args{"hell o"}, false},
		{"hellohello", ExtendedString("hello"), args{"hell ohello"}, true},
		{"hellohello", ExtendedString("helloh"), args{"hellohello"}, true},
		{"empty short", ExtendedString(""), args{"empty short string should match"}, true},
		{"empty long", ExtendedString("empty long should not match"), args{""}, false},
		{"both empty", ExtendedString(""), args{""}, true},
		{"shorter long match", ExtendedString("notempty"), args{"not"}, false},
		{"shorter long no match", ExtendedString("notempty"), args{"yeah"}, false},
		{"aaaaaab didn't kill me", ExtendedString("aaaab"), args{"aaaaaaaaab"}, true},
		{"aaaaaab didn't kill me, see?", ExtendedString("aaaab"), args{"aaaaaaaaaab"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ContainedIn(tt.args.longer); got != tt.want {
				t.Errorf("ExtendedString.IsSubStringOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
