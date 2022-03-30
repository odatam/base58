# Very simple Base58 Encode/Decode

Usage Example :

```
package main

import (
    "github.com/odatam/base58" 
    "fmt"
    "log"
)

func main() {
	enc, err := base58.Base58EncodeToString([]byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}		
	
	fmt.Println(enc)

	dec, err := base58.Base58DecodeToString(enc)
	if err != nil {
		log.Fatal(err)
	}
			
	fmt.Println(string(dec))
}

```

Output :

```
$ ./example

StV1DL6CwTryKyV
hello world
```
