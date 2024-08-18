package maths

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

// Vector represents a vector in n-dimensional space.
type Vector map[string]float64

// NewVectorFromText creates a new vector from a given text by tokenizing it into words and counting their occurrences.
func NewVectorFromText(text string) Vector {
	// Tokenize the text into words.
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	// Count the occurrences of each word.
	vector := make(Vector)
	for _, word := range words {
		word = strings.ToLower(word)
		vector[word]++
	}

	return vector
}

// DotProduct calculates the dot product of two vectors.
func (v Vector) DotProduct(other Vector) float64 {
	sum := 0.0
	for key, value := range v {
		sum += value * other[key]
	}
	return sum
}

// Magnitude calculates the magnitude (Euclidean norm) of a vector.
func (v Vector) Magnitude() float64 {
	sum := 0.0
	for _, value := range v {
		sum += value * value
	}
	return math.Sqrt(sum)
}

// String returns a string representation of the vector.
func (v Vector) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for key, value := range v {
		sb.WriteString(fmt.Sprintf("%s: %f, ", key, value))
	}
	sb.WriteString("}")
	return sb.String()
}
