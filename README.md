# zhtool
traditional chinese word / simplified chinese word convert tool for golang

# Install
```
go get github/sunghau/zhtool
```

# Example
```
package main

import (
	"fmt"
	zhtool "github.com/sunghau/zhtool"
)

func main() {
  s := "遇到一个诡异Bug，每逢周三就崩溃"
  fmt.Println(zhtool.Chs2Cht(s))
}
```

# Author
* SungHau <sunghau@gmail.com>
