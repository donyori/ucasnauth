package main

import (
	"net/url"
	"strings"
)

// Equivalent to JavaScript function encodeURIComponent.
func EncodeUriComponent(s string) string {
	q := url.QueryEscape(s)
	var builder strings.Builder
	builder.Grow(len(q) + strings.Count(q, "+")*2)
	escapeStart, escapeCount := -1, 0
	for i, r := range q {
		switch r {
		case '%':
			escapeStart = i
			escapeCount = 0
		case '+':
			builder.WriteString("%20")
		default:
			if escapeStart < 0 {
				builder.WriteRune(r)
				break
			}
			escapeCount++
			if escapeCount == 1 {
				if r != '2' {
					builder.WriteRune('%')
					builder.WriteRune(r)
					escapeStart = -1
				}
				break
			}
			switch r {
			case '1':
				builder.WriteRune('!')
			case '7':
				builder.WriteRune('\'')
			case '8':
				builder.WriteRune('(')
			case '9':
				builder.WriteRune(')')
			case 'A':
				builder.WriteRune('*')
			default:
				builder.WriteString(q[escapeStart:i])
				builder.WriteRune(r)
			}
			escapeStart = -1
		}
	}
	return builder.String()
}

// This is corresponding to encodeURIComponent(encodeURIComponent(x))
//   in login_bch.js and AuthInterFace.js.
func EncodeUriComponentTwice(s string) string {
	return EncodeUriComponent(EncodeUriComponent(s))
}
