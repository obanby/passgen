package passgen_test

import (
	"crypto/md5"
	"crypto/sha512"
	"github.com/google/go-cmp/cmp"
	"hash"
	"passgen"
	"passgen/randstr"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	t.Parallel()

	pg := passgen.New()
	got := pg.Generate()

	if len(got) == 0 {
		t.Fatalf("wanted password to be not empty but got %d", len(got))
	}
}

func TestGenerateRandomPasswordOfLength(t *testing.T) {
	t.Parallel()

	want := 20
	pg := passgen.New(
		passgen.WithLength(20),
	)

	got := pg.Generate()

	if want != len(got) {
		t.Fatalf("wanted length %d got %s", want, got)
	}
}

func TestGenerateRandomPasswordWithHashingAlgo(t *testing.T) {
	t.Parallel()

	rsg, err := randstr.New(
		randstr.WithSeed(10),
		// this way we avoid map lack of ordering guarantees in go, so we use only letters a through z
		randstr.WithCharRanges(randstr.CharRanges{'a': 'z'}),
	)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		hasher hash.Hash
		want   string
	}{
		{
			hasher: sha512.New(),
			want:   "om7dYUSTNF8kDnmFEXdV",
		},
		{
			hasher: md5.New(),
			want:   "7CHorCN9cVbe/k5LAwHB",
		},
	}

	for _, tc := range testCases {

		t.Log(rsg.Generate())
		pg := passgen.New(
			passgen.WithLength(20),
			passgen.WithHashingFunc(tc.hasher),
			passgen.WithRandStrGenerator(rsg),
		)

		got := pg.Generate()
		if !cmp.Equal(tc.want, got) {
			t.Error(cmp.Diff(tc.want, got))
		}
	}
}

func TestGenerateWrapper(t *testing.T) {
	t.Parallel()

	got := passgen.Generate()
	if len(got) == 0 {
		t.Fatalf("wanted generate wrapper to use reasonable defaults")
	}
}
