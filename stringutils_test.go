package stringutils

import "testing"

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"Hello, world", false},
		{" ", false},
		{"", true},
	}
	
	for _, c := range cases {
		got := IsEmpty(c.in)
		if got != c.want {
			t.Errorf("IsEmpty(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	for i:=0; i < b.N; i++ {
		IsEmpty("Some String")
	}
}

func TestIsNotEmpty(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"Hello, world", true},
		{" ", true},
		{"", false},
	}
	
	for _, c := range cases {
		got := IsNotEmpty(c.in)
		if got != c.want {
			t.Errorf("IsNotEmpty(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func BenchmarkIsNotEmpty(b *testing.B) {
	for i:=0; i < b.N; i++ {
		IsNotEmpty("Some String")
	}
}

func TestIsBlank(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"Hello, world", false},
		{" H", false},
		{" ", true},
		{"  ", true},
		{"", true},
	}
	
	for _, c := range cases {
		got := IsBlank(c.in)
		if got != c.want {
			t.Errorf("IsBlank(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func BenchmarkIsBlank(b *testing.B) {
	for i:=0; i < b.N; i++ {
		IsBlank(" Some")
	}
}

func TestIsNotBlank(t *testing.T) {
	cases := []struct {
		in string
		want bool
	}{
		{"Hello, world", true},
		{" H", true},
		{" ", false},
		{"  ", false},
		{"", false},
	}
	
	for _, c := range cases {
		got := IsNotBlank(c.in)
		if got != c.want {
			t.Errorf("IsNotBlank(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func BenchmarkIsNotBlank(b *testing.B) {
	for i:=0; i < b.N; i++ {
		IsNotBlank(" Some")
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"gnirts", "string"},
		{" H", "H "},
		{" \n", "\n "},
		{"", ""},
	}
	
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("IsNotBlank(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}