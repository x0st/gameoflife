package solution3

import "testing"

func BenchmarkSolution_computeNextGeneration(b *testing.B) {
	solution := Solution{}
	solution.seed()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		solution.computeNextGeneration()
	}
}
