# Go Diff Algorithm
[![Build Status](https://travis-ci.com/ncalabro18/go_differ.svg?token=nA46L6ZRfhuqRxMq99Vv&branch=master)](https://travis-ci.com/ncalabro18/go_differ)
[![codecov](https://codecov.io/gh/ncalabro18/go_differ/branch/master/graph/badge.svg?token=1PY2D2KDY3)](https://codecov.io/gh/ncalabro18/go_differ)

An implementation of the Myers' Greedy Algorithm. In addition, there is a utility to quickly use edit scripts for recording deltas.
### Install
```
go get github.com/ncalabro18/go_differ.git
```
### Example

```go
import (
	"fmt"
	"github.com/ncalabro18/go_differ.git/godiff/v1"
)

func main() {
	f1 := []byte("ABCABBA")
	f2 := []byte("CBABAC")

	d := godiff.NewDelta(f1, f2) //creates the delta script to convert f1 + d -> f2
	scriptResult := d.F2()       //runs the script, f1 is stored in Delta structure
	fmt.Println(string(scriptResult))
}
```
Output:
```
CBABAC
```
