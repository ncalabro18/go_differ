package godiff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBacktrack(t *testing.T) {
	assert := assert.New(t)
	expected := [][]int{
		[]int{
			7, 5, 7, 6,
		},
		[]int{
			6, 4, 7, 5,
		},
		[]int{
			5, 4, 6, 4,
		},
		[]int{
			4, 3, 5, 4,
		},
		[]int{
			3, 2, 4, 3,
		},
		[]int{
			3, 1, 3, 2,
		},
		[]int{
			2, 0, 3, 1,
		},
		[]int{
			1, 0, 2, 0,
		},
		[]int{
			0, 0, 1, 0,
		},
	}

	steps := Backtrack([]byte("ABCABBA"), []byte("CBABAC"))
	//t.Logf("%v\n", steps)

	assert.Equal(expected, steps, "godiff.Backtrack results not correct")
}

func TestTrace(t *testing.T) {


	trace := Trace([]byte("ABCABBA"), []byte("CBABAC"))
	t.Logf("%v\n", trace)
	t.Logf("%v\n", len(trace[0]))
}

func TestRevAssignment(t *testing.T) {
	assert := assert.New(t)

	a := []int{0, 1, 2, 3, 4, 5, 6}
	revAssignment(a, -1, 1)
	assert.Equal(a[6], 1, "revAssignment value not assigned")

	revAssignment(a, -2, 2)
	assert.Equal(a[5], 2, "revAssignment value not assigned")

	revAssignment(a, -3, 3)
	assert.Equal(a[4], 3, "revAssignment value not assigned")

	revAssignment(a, -4, 4)
	assert.Equal(a[3], 4, "revAssignment value not assigned")

	revAssignment(a, -5, 5)
	assert.Equal(a[2], 5, "revAssignment value not assigned")

	revAssignment(a, -6, 6)
	assert.Equal(a[1], 6, "revAssignment value not assigned")

	revAssignment(a, -7, 7)
	assert.Equal(a[0], 7, "revAssignment value not assigned")
}

func TestRevValue(t *testing.T) {
	assert := assert.New(t)

	a := []int{0, 1, 2, 3, 4, 5, 6}

	assert.Equal(a[6], revValue(a, -1), "revValue returned value incorrect")
	assert.Equal(a[5], revValue(a, -2), "revValue returned value incorrect")
	assert.Equal(a[4], revValue(a, -3), "revValue returned value incorrect")
	assert.Equal(a[3], revValue(a, -4), "revValue returned value incorrect")
	assert.Equal(a[2], revValue(a, -5), "revValue returned value incorrect")
	assert.Equal(a[1], revValue(a, -6), "revValue returned value incorrect")
	assert.Equal(a[0], revValue(a, -7), "revValue returned value incorrect")

}
