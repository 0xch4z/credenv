package ui

import "runtime"

var supportsEmoji bool

func init() {
	supportsEmoji = runtime.GOOS == "darwin"
}

// EmojiOr returns emoji if supported or a fallback string
func EmojiOr(e, s string) string {
	if supportsEmoji {
		return e
	}
	return s
}
