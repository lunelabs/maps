# Slice

Generic golang package to help work with slices

## Installation

```
go get github.com/lunelabs/slice
```

## Usage

### Apply

Applies function to every slice element

```go
package main

import (
	"fmt"
	"github.com/lunelabs/slice"
)

func main() {
	s := []int{1, 2, 3, 4}

	fn := func(i int) int {
		return i * 2
	}

	fmt.Println(slice.Apply(s, fn))

	// Output: [2 4 6 8]
}
```

### Reduce

Reduce slice to single value

```go
package main

import (
	"fmt"
	"github.com/lunelabs/slice"
)

func main() {
	s := []int{1, 2, 3, 4}

	fn := func(i, initVal int) int {
		return initVal + i
	}

	fmt.Println(slice.Reduce(s, fn, 0))

	// Output: [10]
}
```

### Filter

Applies filter to slice and returns new slice

```go
package main

import (
	"fmt"
	"github.com/lunelabs/slice"
)

func main() {
	s := []int{1, 2, 3, 4}

	fn := func(i int) bool {
		return i > 1
	}

	fmt.Println(slice.Filter(s, fn))

	// Output: [2, 3, 4]
}
```