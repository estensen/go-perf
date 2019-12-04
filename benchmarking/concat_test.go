package benchmarking

import (
	"math/rand"
	"testing"
)

const wordlength = 10
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func BenchmarkConcatString(b *testing.B) {
	words := make([]string, 100)
	for i := range words {
		words[i] = RandStringBytes(wordlength)
	}

	b.Run("concat string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ConcatString(words)
		}
	})

	b.Run("concat buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ConcatBuffer(words)
		}
	})

	b.Run("concat builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ConcatBuilder(words)
		}
	})
}

func RandStringBytes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
