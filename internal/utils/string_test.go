package utils

import (
	"testing"
)

func TestIsEmptyString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Str Ok", args: args{"test string"}, want: false},
		{name: "Str Ok #2", args: args{"  test string"}, want: false},
		{name: "Empty Str", args: args{""}, want: true},
		{name: "White space Str", args: args{"     \t "}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyString(tt.args.str); got != tt.want {
				t.Errorf("IsEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StringContainsAlphaNumeric(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Str alpha Ok", args: args{"test string"}, want: true},
		{name: "Str alpha Ok #2", args: args{" \t test string  "}, want: true},
		{name: "Str digit Ok", args: args{"123"}, want: true},
		{name: "Str digit Ok #2", args: args{"   123  "}, want: true},
		{name: "Str digit Ok #2", args: args{""}, want: false},
		{name: "Str digit Ok #2", args: args{"   "}, want: false},
		{name: "Str digit Ok #2", args: args{";"}, want: false},
		{name: "Str digit Ok #2", args: args{"   ,  , "}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaNumericString(tt.args.str); got != tt.want {
				t.Errorf("stringConainsAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
