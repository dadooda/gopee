
//
// Originally from: https://github.com/dadooda/gopee.
//

package gopee

import (
	"fmt"
	"reflect"
	"strings"
)

// P prints labels and values to stdout. Strings ending with ":" are treated as labels, the value following
// is printed on the same line. Labels starting with " " are printed on the same line, producing more compact
// and readable output.
func P(args ...interface{}) {
	afterLabel := false

	for i, v := range args {
		var format string

		if afterLabel {
			// The value following the label.
			format = "%#v"
			afterLabel = false
		} else {
			// Value or label.

			var nl string

			if i > 0 {
				nl = "\n"
			} else {
				nl = ""
			}

			if reflect.TypeOf(v).String() == "string" && strings.HasSuffix(v.(string), ":") {
				// Label.
				if strings.HasPrefix(v.(string), " ") {
					// Same-line.
					format = "%s"
				} else {
					// On a new line.
					format = nl + "%s"
				}

				afterLabel = true
			} else {
				// Regular value.
				format = nl + "%#v"
			}
		}

		fmt.Printf(format, v)
	}

	fmt.Println()
}