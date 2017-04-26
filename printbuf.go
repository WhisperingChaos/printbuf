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
	lf.Init("Preamle of message:")
	...
	lf.Sprintf("tail")
	...
	var err error = lf;
	...
	if err.(LoadFail) {
		fmt.Fprintf(err.Error())
		//	output: "Preamble of message:tail"
	}

Notes

* Defined methods as value receivers:
    Why?
		- To reinforce convention that a custom type should be declared as
		  a value type: "var <varName> <CustomErrorType>".  Therefore its
		  type assertion can be encoded as: "err.(<CustomErrorType>)" which seems
		  more readable than declaring the variable as a pointer type.  A pointer
		  type would requre: "err.(*<CustomErrorType>)".

* Use Init() method after declaring a custom type to properly instantiate custom type.
	Why?
		- Eliminates need to code custom new/init method for every custome type.  However
		  these custom methods can be enoded if necessary.
*/
type T struct {
	*bytes.Buffer
}

func New(preamble string) (pb *T) {
	pb = new(T)
	return pb.Init(preamble)
}
func (pb *T) Init(preamble string) *T {
	pb.Buffer = bytes.NewBufferString(preamble)
	return pb
}
func (pb T) Sprintf(format string, args ...interface{}) T {
	pb.WriteString(fmt.Sprintf(format, args...))
	return pb
}
func (pb T) Sprintln(a ...interface{}) T {
	pb.WriteString(fmt.Sprintln(a...))
	return pb
}
func (pb T) Error() string {
	return string(pb.Bytes())
}
