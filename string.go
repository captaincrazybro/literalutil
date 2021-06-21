package lu

import "strings"

// String type of string to add class functions to
type String string

// ToS returns the raw string
func (s String) ToS() string {
	return string(s)
}

func (s String) Replace(old, new string, num int) String {
	return String(strings.Replace(s.ToS(), old, new, num))
}

func (s String) Split(sep string) []String {
	sz := strings.Split(s.ToS(), sep)
	var Sz []String

	for _, v := range sz {
		Sz = append(Sz, String(v))
	}
	return Sz
}

func (s String) Compare(b string) int {
	return strings.Compare(s.ToS(), b)
}

func (s String) Contains(substr string) bool {
	return strings.Contains(s.ToS(), substr)
}

func (s String) ContainsRune(r rune) bool {
	return strings.ContainsRune(s.ToS(), r)
}

func (s String) ContainsChar(chars string) bool {
	return strings.ContainsAny(s.ToS(), chars)
}

func (s String) Count(substr string) int {
	return strings.Count(s.ToS(), substr)
}

func (s String) EqualFold(t string) bool {
	return strings.EqualFold(s.ToS(), t)
}

func (s String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.ToS(), prefix)
}

func (s String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.ToS(), suffix)
}

func (s String) Index(substr string) int {
	return strings.Index(s.ToS(), substr)
}

func (s String) NewReader() *strings.Reader {
	return strings.NewReader(s.ToS())
}

func (s String) Repeat(count int) String {
	return String(strings.Repeat(s.ToS(), count))
}

func (s String) ReplaceAll(old, new string) String {
	return String(strings.ReplaceAll(s.ToS(), old, new))
}

func (s String) ToLower() String {
	return String(strings.ToLower(s.ToS()))
}

func (s String) ToUpper() String {
	return String(strings.ToUpper(s.ToS()))
}

func (s String) Trim(cutset string) String {
	return String(strings.Trim(s.ToS(), cutset))
}