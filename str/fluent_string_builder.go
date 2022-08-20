package str

import (
	"fmt"
	"strings"
)

type (

	// FluentStringBuilder is strings.Builder wrapper,
	// but its api is fluent.
	FluentStringBuilder struct {
		sb strings.Builder
	}

	WriteFunc func(fluent *FluentStringBuilder)
)

func (b *FluentStringBuilder) Write(p []byte) (n int, err error) {
	return b.sb.Write(p)
}

// NewFluent new fluent string builder
func NewFluent() *FluentStringBuilder {
	return &FluentStringBuilder{
		sb: strings.Builder{},
	}
}

// NewLine append NewLine
func (b *FluentStringBuilder) NewLine() *FluentStringBuilder {
	return b.Str(NewLine)
}

// Space append Space
func (b *FluentStringBuilder) Space(times ...int) *FluentStringBuilder {
	return b.Str(RepeatSpace(times...))
}

// Str append string
func (b *FluentStringBuilder) Str(s string) *FluentStringBuilder {
	_, _ = b.sb.WriteString(s)
	return b
}

// Brackets wrap ( s )
func (b *FluentStringBuilder) Brackets(s string) *FluentStringBuilder {
	b.Str("(").Str(s).Str(")")
	return b
}

// WrapSpace " " + s + " "
func (b *FluentStringBuilder) WrapSpace(s string) *FluentStringBuilder {
	b.Str(WrapSpace(s))
	return b
}

// WriteFunc call f get string and write into FluentStringBuilder.
func (b *FluentStringBuilder) WriteFunc(f WriteFunc) *FluentStringBuilder {
	f(b)
	return b
}

// WithSlice traverse slice and call mapper
func (b *FluentStringBuilder) WithSlice(slice []string, mapper func(idx int, item string) string) *FluentStringBuilder {
	if len(slice) == 0 {
		return nil
	}

	for i, s := range slice {
		b.Str(mapper(i, s))
	}

	return b
}

func (b *FluentStringBuilder) Join(str []string, seq string) *FluentStringBuilder {
	if len(str) == 0 {
		return b
	}
	return b.Str(strings.Join(str, seq))
}

// Joins concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func (b *FluentStringBuilder) Joins(elems []fmt.Stringer, sep string) *FluentStringBuilder {
	switch len(elems) {
	case 0:
		return b
	case 1:
		return b.Str(elems[0].String())
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i].String())
	}

	b.sb.Grow(n)
	b.sb.WriteString(elems[0].String())

	for _, s := range elems[1:] {
		b.sb.WriteString(sep)
		b.sb.WriteString(s.String())
	}

	return b
}

func (b *FluentStringBuilder) Bool(value bool) *FluentStringBuilder {
	if value {
		b.Str("true")
	} else {
		b.Str("false")
	}
	return b
}

// Len returns the number of accumulated bytes; b.Len() == len(b.String()).
func (b *FluentStringBuilder) Len() int {
	return b.sb.Len()
}

func (b *FluentStringBuilder) String() string {
	return b.sb.String()
}
