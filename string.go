package lu

import (
	"fmt"
	"strings"
)

// String type of string to add class functions to
type String string

// s returns the raw string
func (Str String) s() string {
	return string(Str)
}

func (Str String) Replace(old, new string, num int) String {
	return String(strings.Replace(Str.s(), old, new, num))
}

func (Str String) Split(sep string) []String {
	sz := strings.Split(Str.s(), sep)
	var Sz []String

	for _, v := range sz {
		Sz = append(Sz, String(v))
	}
	return Sz
}

func (Str String) Compare(b string) int {
	return strings.Compare(Str.s(), b)
}

func (Str String) Contains(substr string) bool {
	return strings.Contains(Str.s(), substr)
}

func (Str String) ContainsRune(r rune) bool {
	return strings.ContainsRune(Str.s(), r)
}

func (Str String) ContainsChar(charStr String) bool {
	return strings.ContainsAny(Str.s(), charStr.s())
}

func (Str String) Count(substr string) int {
	return strings.Count(Str.s(), substr)
}

func (Str String) EqualFold(t string) bool {
	return strings.EqualFold(Str.s(), t)
}

func (Str String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(Str.s(), prefix)
}

func (Str String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(Str.s(), suffix)
}

func (Str String) Index(substr string) int {
	return strings.Index(Str.s(), substr)
}

func (Str String) NewReader() *strings.Reader {
	return strings.NewReader(Str.s())
}

func (Str String) Repeat(count int) String {
	return String(strings.Repeat(Str.s(), count))
}

func (Str String) ReplaceAll(old, new string) String {
	return String(strings.ReplaceAll(Str.s(), old, new))
}

func (Str String) ToLower() String {
	return String(strings.ToLower(Str.s()))
}

func (Str String) ToUpper() String {
	return String(strings.ToUpper(Str.s()))
}

func (Str String) Trim(cutset string) String {
	return String(strings.Trim(Str.s(), cutset))
}

func (Str String) TrimPrefix(p string) String {
	return String(strings.TrimPrefix(Str.s(), p))
}

func (Str String) TrimSuffix(S String) String {
	return String(strings.TrimSuffix(Str.s(), S.s()))
}

// F formats the string
func (S String) F(a ...interface{}) String {
	return String(fmt.Sprintf(S.s(), a...))
}