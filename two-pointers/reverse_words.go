package twopointers

import "strings"

func reverseWords(sentence string) string {
	if len(sentence) == 0 {
		return ""
	}
	var words []string
	sentenceChars := []byte(sentence)
	reverse(sentenceChars, 0, len(sentenceChars)-1)
	for startIdx := 0; startIdx < len(sentenceChars); startIdx++ {
		for spaceIdx := startIdx; spaceIdx < len(sentenceChars); spaceIdx++ {
			if sentenceChars[spaceIdx] == ' ' {
				startIdx++
			} else {
				break
			}
		}
		wordEndIdx := startIdx
		for wordEndIdx < len(sentenceChars) && sentenceChars[wordEndIdx] != ' ' {
			wordEndIdx++
		}
		reverse(sentenceChars, startIdx, wordEndIdx-1)
		words = append(words, string(sentenceChars[startIdx:wordEndIdx]))
		startIdx = wordEndIdx
	}
	return strings.Join(words, " ")
}

func reverse(str []byte, start, end int) {
	for start < end {
		str[start], str[end] = str[end], str[start]
		start++
		end--
	}
}
