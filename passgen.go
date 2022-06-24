package passgen

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"os"
	"passgen/randstr"
	"strconv"
)

type passwordGenerator struct {
	length             int
	randomStrGenerator *randstr.RandString
	hasher             hash.Hash
}

type option func(*passwordGenerator) error

func WithLength(n int) option {
	return func(generator *passwordGenerator) error {
		if n < 0 {
			return errors.New("size must be greater than zero")
		}
		generator.length = n
		generator.randomStrGenerator, _ = randstr.New(
			randstr.WithSize(uint(n)),
		)
		return nil
	}
}

func WithRandStrGenerator(rsg *randstr.RandString) option {
	return func(generator *passwordGenerator) error {
		generator.randomStrGenerator = rsg
		return nil
	}
}

func WithHashingFunc(h hash.Hash) option {
	return func(generator *passwordGenerator) error {
		if h == nil {
			return errors.New("must use a hashing algorithm that implements hash.Hash")
		}
		generator.hasher = h
		return nil
	}
}

func New(opts ...option) *passwordGenerator {
	rsg, _ := randstr.New()
	pg := &passwordGenerator{
		hasher:             sha512.New(),
		length:             64,
		randomStrGenerator: rsg,
	}
	for _, opt := range opts {
		opt(pg)
	}
	return pg
}

func (pg *passwordGenerator) Generate() string {
	rs := pg.randomStrGenerator.Generate()
	pg.hasher.Write([]byte(rs))
	checksum := pg.hasher.Sum([]byte{})
	encodedHash := base64.StdEncoding.EncodeToString(checksum)
	if pg.length <= len(encodedHash) {
		return encodedHash[0:pg.length]
	}
	return encodedHash
}

// TODO: add flags instead of args to support the following:
// 	--length
//  --algo (sha1, sha256, sha512, md5)
func RunCLI() {
	rsg, err := randstr.New()
	if err != nil {
		panic(err)
	}

	desiredLength := 16
	args := os.Args
	if len(args) >= 2 {
		sizeStr := args[1]
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			panic(err)
		}
		desiredLength = size
	}

	pg := New(
		WithRandStrGenerator(rsg),
		WithLength(desiredLength),
	)

	fmt.Println(pg.Generate())
}

func Generate() string {
	return New().Generate()
}
