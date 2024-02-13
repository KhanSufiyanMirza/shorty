package utils

import (
	"errors"
	"regexp"

	"github.com/microcosm-cc/bluemonday"
)

var ErrMalformedInput = errors.New("invalid input")

func Sanitize(v string) error {
	p := bluemonday.StrictPolicy()
	// remove carriage return before sanitizing and comparing
	v = string(regexp.MustCompile("\r\n").ReplaceAll([]byte(v), []byte("\n")))

	sv := p.Sanitize(v)
	if v != sv {
		return ErrMalformedInput
	}

	return nil
}
