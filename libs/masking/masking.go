package masking

import (
	"reflect"
	"strings"
	"unicode/utf8"
)

const MaskChar = "*"

func Mask(v string) string {
	f, _ := utf8.DecodeRuneInString(v)
	l, _ := utf8.DecodeLastRuneInString(v)

	return string(f) + strings.Repeat(MaskChar, utf8.RuneCountInString(v)-2) + string(l)
}

func IsMasking(key string) bool {
	_, ok := MaskingList[key]
	return ok
}

func MaskMap(m map[string]any) {
	for k, v := range m {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			MaskMap(v.(map[string]any))
		case reflect.String:
			if IsMasking(k) {
				m[k] = Mask(v.(string))
			}
		}
	}
}
