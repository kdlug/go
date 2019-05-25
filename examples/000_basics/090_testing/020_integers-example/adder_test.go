// Go Examples
package integers

import (
	"fmt"
)

// Go examples are executed just like tests
// Please note that the example function will not be executed if you remove the comment "//Output: 6"
// By adding this code the example will appear in the documentation inside godoc
// godoc -http=:8010
// GOPATH/src/github.com/{your_id}
// http://localhost:8010/pkg/course/000_basics/090_testing/020_integers/
func ExampleAdd() {
	sum := Add(2, 2)
	// expected := 4

	// if sum != expected {
	// 	t.Errorf("Expected '%d' but got '%d'", expected, sum)
	// }

	fmt.Println(sum)
	// Output: 4
}
