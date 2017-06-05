package printbuf

import (
	"bytes"
	"fmt"
)

/*
Support buffered printing especially for error messages. This type
implements the "error" interface allowing types that consume this
one, as an anonymous field, to automatically project the public
interface below easing the creation of a custom error type.  A
custom type coupled with a type assertion enables a swift and direct
comparison mechanism to test for a specific error instead of inspecting
the text of an errormessage.

Example

	Define custom type called 'LoadFail':

	type LoadFail struct {
		printbuf.T
	}

	Declare variable 'lf' based on 'LoadFail' custom type:

	var lf LoadFail
	lf.Sprintf("message")
	...
	var err error = lf;
	...
	if err.(LoadFail) {
		fmt.Fprintf(err.Error())
		//	output: "message"
	}

Notes

* Currently not concurrency safe.
*/
type T struct {
	*bytes.Buffer
}

func (pb *T) Sprintf(format string, args ...interface{}) {
	pb.init()
	pb.WriteString(fmt.Sprintf(format, args...))
}
func (pb *T) Sprintln(a ...interface{}) {
	pb.init()
	pb.WriteString(fmt.Sprintln(a...))
}
func (pb T) Error() string {
	if pb.Buffer == nil {
		return String()
	}
	return pb.String()
}

func (pb *T) init() {
	if pb.Buffer == nil {
		pb.Buffer = bytes.NewBufferString("")
	}
}
