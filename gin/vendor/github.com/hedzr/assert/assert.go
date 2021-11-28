package assert

import (
	"fmt"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

// Equal validates that 'actual' is equal to 'expect' and throws an error with line number
func Equal(t testing.TB, expect, actual interface{}) {
	EqualSkip(t, 2, expect, actual)
}

// EqualSkip validates that 'actual' is equal to 'expect' and throws an error with line number
// but the skip variable tells EqualSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func EqualSkip(t testing.TB, skip int, expect, actual interface{}) {
	if !isEqual(expect, actual) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d expecting %v (%v), but got %v (%v). DIFF is: %v\n",
			path.Base(file), line,
			expect, reflect.TypeOf(expect),
			actual, reflect.TypeOf(actual),
			DiffValues(expect, actual))
		t.FailNow()
	}
}

// NotEqual validates that val1 is not equal val2 and throws an error with line number
func NotEqual(t *testing.T, expect, actual interface{}) {
	NotEqualSkip(t, 2, expect, actual)
}

// NotEqualSkip validates that val1 is not equal to val2 and throws an error with line number
// but the skip variable tells NotEqualSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func NotEqualSkip(t *testing.T, skip int, expect, actual interface{}) {
	if isEqual(expect, actual) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d expecting differ with %v (%v), but got %v (%v).\n",
			path.Base(file), line,
			expect, reflect.TypeOf(expect),
			actual, reflect.TypeOf(actual))
		t.FailNow()
	}
}

// EqualTrue validates that 'actual' is true
func EqualTrue(t testing.TB, actual bool) {
	EqualSkip(t, 2, true, actual)
}

// EqualFalse validates that 'actual' is false
func EqualFalse(t testing.TB, actual bool) {
	EqualSkip(t, 2, false, actual)
}

// Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
func Nil(t testing.TB, value interface{}) {
	NilSkip(t, 2, value)
}

// NilSkip asserts that the specified object is nil.
//
//    assert.NilSkip(t, err)
func NilSkip(t testing.TB, skip int, value interface{}) {
	if reflect.ValueOf(value).IsNil() {
		return
	}
	_, file, line, _ := runtime.Caller(skip)
	fmt.Printf("%s:%d %v should be nil\n", path.Base(file), line, value)
	t.FailNow()
}

// NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, err)
func NotNil(t testing.TB, value interface{}) {
	NotNilSkip(t, 2, value)
}

// NotNilSkip asserts that the specified object is not nil.
//
//    assert.NotNilSkip(t, 1, err)
func NotNilSkip(t testing.TB, skip int, value interface{}) {
	if !reflect.ValueOf(value).IsNil() {
		return
	}
	_, file, line, _ := runtime.Caller(skip)
	fmt.Printf("%s:%d %v should NOT be nil\n", path.Base(file), line, value)
	t.FailNow()
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoError(t, err) {
//	   assert.Equal(t, expectedObj, actualObj)
//   }
func NoError(t testing.TB, err error) {
	if err != nil {
		t.Fatalf("expecting no error but got %v", err)
	}
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Error(t, err) {
//	   assert.Equal(t, expectedError, err)
//   }
func Error(t testing.TB, err error) {
	if err == nil {
		t.Fatal("expecting error occurs but got nil")
	}
}

// PanicMatches validates that the panic output of running fn matches the supplied string
func PanicMatches(t testing.TB, fn func(), matches interface{}) {
	PanicMatchesSkip(t, 2, fn, matches)
}

// PanicMatchesSkip validates that the panic output of running fn matches the supplied string
// but the skip variable tells PanicMatchesSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func PanicMatchesSkip(t testing.TB, skip int, fn func(), matches interface{}) {
	_, file, line, _ := runtime.Caller(skip)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("%s", r)

			if err != matches {
				fmt.Printf("%s:%d Panic...  expected [%s] received [%s]", path.Base(file), line, matches, err)
				t.FailNow()
			}
		} else {
			fmt.Printf("%s:%d Panic Expected, none found...  expected [%s]", path.Base(file), line, matches)
			t.FailNow()
		}
	}()

	fn()
}

// NotMatch validates that value matches the regex, either string or *regex
// and throws an error with line number
func NotMatch(t *testing.T, value string, regex interface{}) {
	NotMatchSkip(t, 2, value, regex)
}

// NotMatchSkip validates that value matches the regex, either string or *regex
// and throws an error with line number
// but the skip variable tells NotMatchRegexSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func NotMatchSkip(t *testing.T, skip int, value string, regex interface{}) {

	if r, ok, err := regexMatches(regex, value); ok || err != nil {
		_, file, line, _ := runtime.Caller(skip)

		if err != nil {
			fmt.Printf("%s:%d %v error compiling regex %v\n", path.Base(file), line, value, r.String())
		} else {
			fmt.Printf("%s:%d %v matches regex %v\n", path.Base(file), line, value, r.String())
		}

		t.FailNow()
	}
}

// Match validates that value matches the regex, either string or *regex
// and throws an error with line number
func Match(t *testing.T, value string, regex interface{}) {
	MatchSkip(t, 2, value, regex)
}

// MatchSkip validates that value matches the regex, either string or *regex
// and throws an error with line number
// but the skip variable tells MatchRegexSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func MatchSkip(t *testing.T, skip int, value string, regex interface{}) {

	if r, ok, err := regexMatches(regex, value); !ok {
		_, file, line, _ := runtime.Caller(skip)

		if err != nil {
			fmt.Printf("%s:%d %v error compiling regex %v\n", path.Base(file), line, value, r.String())
		} else {
			fmt.Printf("%s:%d %v does not match regex %v\n", path.Base(file), line, value, r.String())
		}

		t.FailNow()
	}
}

func regexMatches(regex interface{}, value string) (r *regexp.Regexp, matched bool, err error) {
	var ok bool
	r, ok = regex.(*regexp.Regexp)

	// must be a string
	if !ok {
		if r, err = regexp.Compile(regex.(string)); err != nil {
			return
		}
	}

	matched = r.MatchString(value)
	return
}
