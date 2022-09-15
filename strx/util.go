package strx

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	BYTE = 1.0 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

// FormatBytes format bytes unit
func FormatBytes(bytes int64) string {
	unit := ""
	value := float32(bytes)

	switch {

	case bytes >= TERABYTE:
		unit = "TB"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		unit = "GB"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "MB"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "KB"
		value = value / KILOBYTE
	case bytes == 0:
		return "0"

	}

	result := fmt.Sprintf("%.2f", value)
	result = strings.TrimSuffix(result, ".00")
	return fmt.Sprintf("%s%s", result, unit)
}

const (
	Pass = 1
	K    = 1000 * Pass
	M    = K * 10
)

// FormatNumber format number unit
func FormatNumber(number int64) string {
	unit := ""
	value := float32(number)

	switch {

	case number >= M:
		unit = "m"
		value = value / M
	case number >= K:
		unit = "k"
		value = value / K
	case number == 0:
		return "0"

	}

	result := fmt.Sprintf("%.1f", value)
	result = strings.TrimSuffix(result, ".00")
	return fmt.Sprintf("%s%s", result, unit)
}

// Substring source[start:end)
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	if start == end {
		return string(r[start])
	}

	var substring = ""
	for i := 0; i < length; i++ {
		if i < start {
			continue
		}
		if i >= end {
			break
		}
		substring += string(r[i])
	}

	return substring
}

func FuzzyAgoAbbr(now time.Time, createdAt time.Time) string {
	ago := now.Sub(createdAt)

	if ago < time.Hour {
		return fmt.Sprintf("%d%s", int(ago.Minutes()), "m")
	}
	if ago < 24*time.Hour {
		return fmt.Sprintf("%d%s", int(ago.Hours()), "h")
	}
	if ago < 30*24*time.Hour {
		return fmt.Sprintf("%d%s", int(ago.Hours())/24, "d")
	}

	return createdAt.Format("Jan _2, 2006")
}

// Truncate truncate string
func Truncate(s string, size int) string {
	if len(s) < size {
		return s
	}

	return Substring(s, 0, size)
}

// ToInt cast string to int,default 0
func ToInt(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}

func RepeatSpace(times ...int) string {
	count := 1
	if len(times) > 0 {
		count = times[0]
	}
	return strings.Repeat(Space, count)
}

// WrapSpace " " + s + " "
func WrapSpace(s string) string {
	return Wrap(Space, Space, s)
}

// Wrap  left + s + right
func Wrap(left, right, s string) string {
	return left + s + right
}

func RemoveEmpty(str []string) []string {
	if len(str) == 0 {
		return []string{}
	}

	var result []string
	for _, s := range str {
		if len(s) > 0 {
			result = append(result, s)
		}
	}

	return result
}

func BoolMapYesOrNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

// Equal check, alias of strings.EqualFold
var Equal = strings.EqualFold

// NoCaseEq check two strings is equals and case-insensitivity
func NoCaseEq(s, t string) bool {
	return strings.EqualFold(s, t)
}

// IsNumChar returns true if the given character is a numeric, otherwise false.
func IsNumChar(c byte) bool {
	return c >= '0' && c <= '9'
}

var numReg = regexp.MustCompile("^\\d+$")

// IsNumeric returns true if the given string is a numeric, otherwise false.
func IsNumeric(s string) bool {
	return numReg.MatchString(s)
}

// IsAlphabet char
func IsAlphabet(char uint8) bool {
	// A 65 -> Z 90
	if char >= 'A' && char <= 'Z' {
		return true
	}

	// a 97 -> z 122
	if char >= 'a' && char <= 'z' {
		return true
	}
	return false
}

// IsAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func IsAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// StrPos alias of the strings.Index
func StrPos(s, sub string) int {
	return strings.Index(s, sub)
}

// BytePos alias of the strings.IndexByte
func BytePos(s string, bt byte) int {
	return strings.IndexByte(s, bt)
}

// HasOneSub substr in the given string.
func HasOneSub(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// HasAllSubs all substr in the given string.
func HasAllSubs(s string, subs []string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// IsStartsOf alias of the HasOnePrefix
func IsStartsOf(s string, prefixes []string) bool {
	return HasOnePrefix(s, prefixes)
}

// HasOnePrefix the string start withs one of the subs
func HasOnePrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasPrefix substr in the given string.
func HasPrefix(s string, prefix string) bool { return strings.HasPrefix(s, prefix) }

// IsStartOf alias of the strings.HasPrefix
func IsStartOf(s, prefix string) bool { return strings.HasPrefix(s, prefix) }

// HasSuffix substr in the given string.
func HasSuffix(s string, suffix string) bool { return strings.HasSuffix(s, suffix) }

// IsEndOf alias of the strings.HasSuffix
func IsEndOf(s, suffix string) bool { return strings.HasSuffix(s, suffix) }

// IsValidUtf8 valid utf8 string check
func IsValidUtf8(s string) bool { return utf8.ValidString(s) }

// ----- refer from github.com/yuin/goldmark/util

// refer from github.com/yuin/goldmark/util
var spaceTable = [256]int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// IsSpace returns true if the given character is a space, otherwise false.
func IsSpace(c byte) bool {
	return spaceTable[c] == 1
}

// IsEmpty returns true if the given string is empty.
func IsEmpty(s string) bool { return len(s) == 0 }

// IsBlank returns true if the given string is all space characters.
func IsBlank(s string) bool {
	return IsBlankBytes([]byte(s))
}

// IsNotBlank returns true if the given string is not blank.
func IsNotBlank(s string) bool {
	return !IsBlankBytes([]byte(s))
}

// IsBlankBytes returns true if the given []byte is all space characters.
func IsBlankBytes(bs []byte) bool {
	for _, b := range bs {
		if !IsSpace(b) {
			return false
		}
	}
	return true
}

// IsSymbol reports whether the rune is a symbolic character.
func IsSymbol(r rune) bool {
	return unicode.IsSymbol(r)
}

var verRegex = regexp.MustCompile(`^[0-9][\d.]+(-\w+)?$`)

// IsVersion number. eg: 1.2.0
func IsVersion(s string) bool {
	return verRegex.MatchString(s)
}

// Compare for two string.
func Compare(s1, s2, op string) bool {
	return VersionCompare(s1, s2, op)
}

// VersionCompare for two version string.
func VersionCompare(v1, v2, op string) bool {
	switch op {
	case ">":
		return v1 > v2
	case "<":
		return v1 < v2
	case ">=":
		return v1 >= v2
	case "<=":
		return v1 <= v2
	default:
		return v1 == v2
	}
}
