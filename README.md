# passgen

A random password generator library and cli written in GO.

The api is quite simple and concise.

To simply generate passwords you can use the simplest form of this library

```go
package main

import (
	"fmt"
	"passgen"
)

func main() {
	pass := passgen.Generate()
	fmt.Println(pass)
}
````

The above code will generate a `base64` encoded string. The encoded string comes from `SHA521` hashing algorithm
provided
by go stdlib in the `crypto` package. The hashed original string is randomly generated through a package `randstr` which
handles providing a seed and allows for customization.

If you would like to customize the password generator you can do the following.

```go

package main

import (
	"crypto/md5"
	"fmt"
	"passgen"
)

func main() {
	pg := passgen.New(
		passgen.WithLength(20),
		passgen.WithHashingFunc(md5.New()),
	)

	pass := pg.Generate()
	fmt.Println(pass)
}
```

Lastly if you would like to customize the random string generator you can use the following

```go
package main

import (
	"crypto/md5"
	"fmt"
	"passgen"
	"passgen/randstr"
)

func main() {
	seed := 9810
	charset := randstr.CharRanges{
		'a':     'z',
		'A':     'Z',
		'-':     '-',
		rune(0): rune(250),
	}
	randStrGen, err := randstr.New(
		randstr.WithSize(100),
		randstr.WithSeed(seed),
		randstr.WithCharRanges(charset),
	)

	if err != nil {
		panic(err)
	}

	pg := passgen.New(
		passgen.WithLength(30),
		passgen.WithHashingFunc(md5.New()),
		passgen.WithRandStrGenerator(randStrGen),
	)

	pass := pg.Generate()
	fmt.Println(pass)
}
```

# CLI usage

You can simply use `./passgen` to create a unique password printed to your terminal. 
Moreover, you can use it like that to do administrative stuff. However, beware that this will show up in your terminal 
logs
```bash
cd ./cmd/passgen
go build .
./passgen <size> | read -s TestPass
echo $TestPass
```

## Future improvements for the lib
- [ ] Add interface to allow custom implementation of the random string generator

## Future improvements for the cli
- [ ] Add flags to specify the hashing algorithm to use
- [ ] Add flags to specify the length instead of depending on args

> ⚠️ I am not a security person! So open PR if you are a security ninja and see flaws in this lib or cli!