package cosine

import (
	"errors"
	"fmt"
)

// CosineSimilarity calculates the cosine similarity between two vectors.
func CosineSimilarity(v1, v2 Vector) (float64, error) {
	dotProduct := v1.DotProduct(v2)
	magnitude1 := v1.Magnitude()
	magnitude2 := v2.Magnitude()

	if magnitude1 == 0 || magnitude2 == 0 {
		return 0, errors.New("magnitude of one or both vectors is zero")
	}

	similarity := dotProduct / (magnitude1 * magnitude2)
	return similarity, nil
}

// TextSimilarity calculates the cosine similarity between two texts by converting them into vectors first.
func TextSimilarity(text1, text2 string) (float64, error) {
	vector1 := NewVectorFromText(text1)
	vector2 := NewVectorFromText(text2)

	return CosineSimilarity(vector1, vector2)
}
