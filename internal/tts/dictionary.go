package tts

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func dictionaryTrans(text string) string {
	if t := wTrans(text); t != "" {
		return t
	}
	return text
}

func wTrans(text string) string {
	re := regexp.MustCompile(`[wWｗ]+`)
	indexs := re.FindAllStringIndex(text, -1)
	// 単語の一部として使われる場合を考慮して、前後が英字・数字・記号でない場合にのみ変換を行う
	if len(indexs) == 0 {
		return ""
	}

	var b strings.Builder
	last := 0
	for _, idx := range indexs {
		start, end := idx[0], idx[1]

		var prev rune
		var hasPrev bool
		if start > 0 {
			prev, _ = utf8.DecodeLastRuneInString(text[:start])
			hasPrev = true
		}

		var next rune
		var hasNext bool
		if end < len(text) {
			next, _ = utf8.DecodeRuneInString(text[end:])
			hasNext = true
		}

		shouldConvert := true
		if hasPrev && unicode.IsLower(unicode.ToLower(prev)) {
			shouldConvert = false
		}
		if hasNext && unicode.IsLower(unicode.ToLower(next)) {
			shouldConvert = false
		}

		if shouldConvert {
			b.WriteString(text[last:start])
			// 変換内容: 連続する w/W/ｗ を 1 つの「草」に置換
			b.WriteString("うふふ")
		} else {
			// 変換対象外: 元の文字列をそのまま出力
			b.WriteString(text[last:end])
		}
		last = end
	}
	b.WriteString(text[last:])
	return b.String()
}
