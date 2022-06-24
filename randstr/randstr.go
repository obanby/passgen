package randstr

import (
	"math/rand"
	"strings"
	"time"
)

type CharRanges map[rune]rune

type RandString struct {
	charset string
	size    int
	seed    int
}

func New(opts ...option) (*RandString, error) {
	randStr := &RandString{
		charset: generateCharset(CharRanges{
			'a': 'z',
			'A': 'Z',
			'1': '9',
		}),
		size: 10,
		seed: int(time.Now().UnixNano()),
	}
	for _, opt := range opts {
		err := opt(randStr)
		if err != nil {
			return &RandString{}, err
		}
	}

	return randStr, nil
}

func (r *RandString) Generate() string {
	strBuilder := strings.Builder{}
	randomizer := rand.New(
		rand.NewSource(int64(r.seed)),
	)
	for i := 0; i < r.size; i++ {
		idx := randomizer.Intn(len(r.charset))
		b := r.charset[idx]
		strBuilder.Write([]byte{b})
	}
	return strBuilder.String()
}

type option func(*RandString) error

func WithSeed(seed int) option {
	return func(rs *RandString) error {
		rs.seed = seed
		return nil
	}
}

func WithSize(n uint) option {
	return func(rs *RandString) error {
		rs.size = int(n)
		return nil
	}
}

func WithCharset(str string) option {
	return func(rs *RandString) error {
		rs.charset = str
		return nil
	}
}

func WithCharRanges(ch CharRanges) option {
	return func(rs *RandString) error {
		// TODO: validate charset values
		rs.charset = generateCharset(ch)
		return nil
	}
}

func generateCharset(cr CharRanges) string {
	charset := strings.Builder{}
	for from, to := range cr {
		for i := from; i <= to; i++ {
			c := byte(i)
			charset.Write([]byte{c})
		}
	}
	return charset.String()
}

func Generate() string {
	alphaNumericCharSet := CharRanges{
		'a': 'z',
		'A': 'Z',
		'1': '9',
	}

	randStr, _ := New(
		WithCharRanges(alphaNumericCharSet),
		WithSize(15),
	)
	return randStr.Generate()
}
