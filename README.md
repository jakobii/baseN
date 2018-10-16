![alt text](https://raw.githubusercontent.com/jakobii/baseN/master/bN.svg "baseN Logo")




# baseN
baseN lets you encode any number into any base using your choice of utf8 characters.
character values are determined by the order you specify them.

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
    b4 := []rune("😁😡😪😋")
    e := baseN.Encode(255, b4)
    fmt.Println(e)
}
// prints: 😋😋😋😋
```

Base62 __[0-9a-zA-Z]__
```go
package main
import (
    "github.com/jakobii/baseN"
    "fmt"
)
func main (){
    b62 := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    e := baseN.Encode(65434198762432165, b62)
    fmt.Println(e)
}
// prints: 4PGJLkZKpD
```