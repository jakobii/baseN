
# baseN
baseN lets you encode any number into any base using your choice of utf8 characters. character values are determined by the order you specify them. the base number is detirmined by the length of the character set

`go get github.com/jakobii/baseN`

## Examples

Base4 in emojis.
```go
package main
import (
    "github.com/jakobii/baseN"
    "fmt"
)
func main (){
    
    // create a base 4 in emojis 
    emojis := []rune("\xF0\x9F\x98\x81" + "\xF0\x9F\x98\xA1" + "\xF0\x9F\x98\xBC" + "\xF0\x9F\x98\x8B")
    b4,_:= baseN.NewBase(emojis)
    
    // encode 255 in our base 4 emojis
    e := b4.Encode(255)
    fmt.Println(string(e)) // prints: 😋😋😋😋
    
    //decode our emojis back to base 10
    d := b4.Decode(e)
    fmt.Println(d) // prints: 255

}

```

Base62 __[0-9a-zA-Z]__
```go
package main
import (
    "github.com/jakobii/baseN"
    "fmt"
)
func main (){
    chars := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b62,_:= baseN.NewBase(chars)
    
    // encode 255 in base 62
    e := b62.Encode(255)
    fmt.Println(string(e)) // prints: 47
    
    //decode back to base 10
    d = b62.Decode(e)
    fmt.Println(d) // prints: 255

}
```