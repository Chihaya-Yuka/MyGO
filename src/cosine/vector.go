package cosine

import (
	"errors"
	"sync"
)

// CosineSimilarity calculates the cosine similarity between two vectors.
// If CUDA is available, it will use GPU for computation.
func CosineSimilarity(v1, v2 *Vector, useGPU bool) (float64, error) {
	if useGPU {
		// GPU implementation (assuming vectors are loaded into gorgonia Nodes).
		return calculateGPUSimilarity(v1, v2)
	} else {
		// Fallback to CPU implementation.
		dotProduct := v1.DotProduct(v2)
		magnitude1 := v1.Magnitude()
		magnitude2 := v2.Magnitude()

		if magnitude1 == 0 || magnitude2 == 0 {
			return 0, errors.New("magnitude of one or both vectors is zero")
		}

		return dotProduct / (magnitude1 * magnitude2), nil
	}
}

// TextSimilarity calculates the cosine similarity between two texts by converting them into vectors first.
func TextSimilarity(text1, text2 string) (float64, error) {
	useGPU := CheckCUDA() // Check if we should use GPU

	vector1 := NewVectorFromText(text1)
	vector2 := NewVectorFromText(text2)

	var wg sync.WaitGroup
	var similarity float64
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		similarity, err = CosineSimilarity(vector1, vector2, useGPU)
	}()

	wg.Wait()
	return similarity, err
}

// calculateGPUSimilarity calculates cosine similarity on GPU using gorgonia.
func calculateGPUSimilarity(v1, v2 *Vector) (float64, error) {
	graph := gorgonia.NewGraph()

	// Load vectors into Gorgonia nodes.
	node1 := gorgonia.NewTensor(graph, gorgonia.Float64, v1.RawMatrix())
	node2 := gorgonia.NewTensor(graph, gorgonia.Float64, v2.RawMatrix())

	// Perform dot product on GPU.
	dotProduct, err := GPUDotProduct(node1, node2)
	if err != nil {
		return 0, err
	}

	// Compute magnitudes.
	magnitude1 := gorgonia.Must(gorgonia.Norm(node1, 2))
	magnitude2 := gorgonia.Must(gorgonia.Norm(node2, 2))

	// Compute cosine similarity.
	similarity := gorgonia.Must(gorgonia.Div(dotProduct, gorgonia.Mul(magnitude1, magnitude2)))

	machine := gorgonia.NewTapeMachine(graph)
	defer machine.Close()

	if err := machine.RunAll(); err != nil {
		return 0, err
	}

	return similarity.Value().Data().(float64), nil
}
