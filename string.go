package lu

import (
	"fmt"
	"strings"
)

// String type of string to add class functions to
type String string

// Tos returns the raw string
func (Str String) Tos() string {
	return string(Str)
}

func (Str String) Replace(old, new string, num int) String {
	return String(strings.Replace(Str.Tos(), old, new, num))
}

func (Str String) Split(sep string) []String {
	sz := strings.Split(Str.Tos(), sep)
	var Sz []String

	for _, v := range sz {
		Sz = append(Sz, String(v))
	}
	return Sz
}

func (Str String) Compare(b string) int {
	return strings.Compare(Str.Tos(), b)
}

func (Str String) Contains(substr string) bool {
	return strings.Contains(Str.Tos(), substr)
}

func (Str String) ContainsRune(r rune) bool {
	return strings.ContainsRune(Str.Tos(), r)
}

func (Str String) ContainsChar(charStr String) bool {
	return strings.ContainsAny(Str.Tos(), charStr.Tos())
}

func (Str String) Count(substr string) int {
	return strings.Count(Str.Tos(), substr)
}

func (Str String) EqualFold(t string) bool {
	return strings.EqualFold(Str.Tos(), t)
}

func (Str String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(Str.Tos(), prefix)
}

func (Str String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(Str.Tos(), suffix)
}

func (Str String) Index(substr string) int {
	return strings.Index(Str.Tos(), substr)
}

func (Str String) NewReader() *strings.Reader {
	return strings.NewReader(Str.Tos())
}

func (Str String) Repeat(count int) String {
	return String(strings.Repeat(Str.Tos(), count))
}

func (Str String) ReplaceAll(old, new string) String {
	return String(strings.ReplaceAll(Str.Tos(), old, new))
}

func (Str String) ToLower() String {
	return String(strings.ToLower(Str.Tos()))
}

func (Str String) ToUpper() String {
	return String(strings.ToUpper(Str.Tos()))
}

func (Str String) Trim(cutset string) String {
	return String(strings.Trim(Str.Tos(), cutset))
}

func (Str String) TrimPrefix(p string) String {
	return String(strings.TrimPrefix(Str.Tos(), p))
}

func (Str String) TrimSuffix(S String) String {
	return String(strings.TrimSuffix(Str.Tos(), S.Tos()))
}

// F formats the string
func (Str String) F(a ...interface{}) String {
	return String(fmt.Sprintf(Str.Tos(), a...))
}