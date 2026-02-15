package useragent

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()

	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

	for b.Loop() {
		Parse(ua)
	}
}

func BenchmarkParseBot(b *testing.B) {
	b.ReportAllocs()

	ua := "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

	for b.Loop() {
		Parse(ua)
	}
}

func BenchmarkParseMobile(b *testing.B) {
	b.ReportAllocs()

	ua := "Mozilla/5.0 (iPhone; CPU iPhone OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1"

	for b.Loop() {
		Parse(ua)
	}
}

func BenchmarkParseUnknown(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		Parse("")
	}
}
