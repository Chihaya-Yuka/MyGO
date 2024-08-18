package cosine

import (
	"gorgonia.org/cu"
	"gorgonia.org/gorgonia"
	"log"
)

// CheckCUDA checks if CUDA is available on the system.
func CheckCUDA() bool {
	dev, err := cu.Init(0)
	if err != nil {
		log.Println("CUDA not available:", err)
		return false
	}
	log.Printf("CUDA device found: %v", dev)
	return true
}

// GPUDotProduct calculates the dot product using GPU (if available).
func GPUDotProduct(v1, v2 *gorgonia.Node) (*gorgonia.Node, error) {
	// Perform the dot product using Gorgonia and GPU.
	graph := gorgonia.NewGraph()
	dot := gorgonia.Must(gorgonia.Dot(v1, v2))

	machine := gorgonia.NewTapeMachine(graph, gorgonia.WithCUDA(true))
	defer machine.Close()

	if err := machine.RunAll(); err != nil {
		return nil, err
	}

	return dot, nil
}
