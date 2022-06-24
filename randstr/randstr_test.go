package randstr_test

import (
	"github.com/google/go-cmp/cmp"
	"testing"

	"passgen/randstr"
)

func TestRandomStrIsNotEmpty(t *testing.T) {
	t.Parallel()

	str := randstr.Generate()
	strLen := len(str)
	if strLen == 0 {
		t.Fatalf("want generated string not to be empty got %q", str)
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()
	
	seed := 1
	rs, err := randstr.New(
		randstr.WithSeed(seed),
		randstr.WithCharset("abcd"),
		randstr.WithSize(4),
	)
	if err != nil {
		t.Fatal(err)
	}

	want := "bddd"
	got := rs.Generate()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
