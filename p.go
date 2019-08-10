
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
	isLabel := func(v interface{}) bool {
		return (reflect.TypeOf(v).Name() == "string" && strings.HasSuffix(v.(string), ":"))
	}

	isSameLineLabel := func(s string) bool {
		return strings.HasPrefix(s, " ")
	}

	newline := func(i int) string {
		if i > 0 {
			return "\n"
		}
		return ""
	}

	afterLabel := false

	for i, v := range args {
		var format string

		if afterLabel {
			// The value following the label.
			format = "%#v"
			afterLabel = false
		} else {
			// Value or label.
			if isLabel(v) {
				if isSameLineLabel(v.(string)) {
					format = "%s"
				} else {
					format = newline(i) + "%s"
				}

				afterLabel = true
			} else {
				// Regular value.
				format = newline(i) + "%#v"
			}
		}

		fmt.Printf(format, v)
	}

	fmt.Println()
}