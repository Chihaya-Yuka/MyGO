package cosine

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gorgonia.org/gorgonia"
	"strings"
	"sync"
	"unicode"
)

// Cache to store previously calculated vectors.
var vectorCache sync.Map

// Vector represents a vector in n-dimensional space.
type Vector struct {
	mat.Sparse
}

// NewVectorFromText creates a new vector from a given text by tokenizing it into words and counting their occurrences.
func NewVectorFromText(text string) *Vector {
	// Check cache
	if cachedVec, ok := vectorCache.Load(text); ok {
		return cachedVec.(*Vector)
	}

	// Tokenize the text into words.
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	// Create a sparse vector.
	vector := &Vector{
		Sparse: mat.NewSparse(0, len(words)),
	}
	for i, word := range words {
		word = strings.ToLower(word)
		vector.Set(i, 0, 1.0)
	}

	// Store the vector in the cache
	vectorCache.Store(text, vector)

	return vector
}

// DotProduct calculates the dot product of two vectors.
func (v *Vector) DotProduct(other *Vector) float64 {
	return mat.Dot(v, other)
}

// Magnitude calculates the magnitude (Euclidean norm) of a vector.
func (v *Vector) Magnitude() float64 {
	return mat.Norm(v, 2)
}

// String returns a string representation of the vector.
func (v *Vector) String() string {
	return fmt.Sprintf("%v", v.RawMatrix())
}
