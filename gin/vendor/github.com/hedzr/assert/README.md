# assert

![Go](https://github.com/hedzr/assert/workflows/Go/badge.svg)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/hedzr/assert.svg?label=release)](https://github.com/hedzr/assert/releases)
[![](https://img.shields.io/badge/go-dev-green)](https://pkg.go.dev/github.com/hedzr/assert)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/hedzr/assert)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fhedzr%2Fgo-ringbuf.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fhedzr%2Fassert?ref=badge_shield)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedzr/assert)](https://goreportcard.com/report/github.com/hedzr/assert)
[![Coverage Status](https://coveralls.io/repos/github/hedzr/assert/badge.svg?branch=master&.9)](https://coveralls.io/github/hedzr/assert?branch=master)
<!--
[![Build Status](https://travis-ci.org/hedzr/assert.svg?branch=master)](https://travis-ci.org/hedzr/assert)
[![codecov](https://codecov.io/gh/hedzr/assert/branch/master/graph/badge.svg)](https://codecov.io/gh/hedzr/assert) 
-->


`assert` provides a set of assertion helpers for unit/bench testing in golang.

`assert` is inspired by these projects:

- <https://github.com/alecthomas/assert>
- <https://github.com/go-playground/assert>
- <https://github.com/stretchr/testify> assert, mock, suite ...

### improvements

- Can be used with both unit test and bench test.
- Most of conventiional assertions: 
  - Equal, NotEqual, EqualTrue, EqualFalse
  - Nil, NotNil
  - Error, NoError
  - PanicMatches: test for the function which might throw a panic
  - Match, NotMatch: compares a value with regexp test 
- Fresh coding in go 1.13~1.15 and later.

### Short guide

```go
package some_test

import (
	"github.com/hedzr/assert"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestEqual(t *testing.T) {
	expected := []*Person{{"Alec", 20}, {"Bob", 21}, {"Sally", 22}}
	actual := []*Person{{"Alex", 20}, {"Bob", 22}, {"Sally", 22}}
	assert.NotEqual(t, expected, actual)

	assert.Equal(t, actual, actual)
}

func TestEqualTrue(t *testing.T) {
	assert.EqualTrue(t, true)
	assert.EqualFalse(t, false)
}
``` 

### LICENSE

MIT
