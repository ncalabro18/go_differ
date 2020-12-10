package godiff

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	dataScript = []byte("--_+B__-_+C")
	dataStrings = []string{
		"abcdefg", "bcdegf", "0001111", "0101111", "jqwk", "k1lj25",
		"alluq", "jkhsakjda,", "oijqqqq", "qqqqoij", "ba", "ab",

	}
)

func TestGenScript(t *testing.T) {
	assert := assert.New(t)

	Backtrack = func(f1, f2 []byte) [][]int {
		return dataBacktrack
	}
	script := GenScript([]byte(dataF1), []byte(dataF2))
	//t.Logf("\"%v\"\n", string(script))
	assert.Equal(dataScript, script, "test script failed to match expected data")
}

func TestRunScript(t *testing.T) {
	assert := assert.New(t)

	result := RunScript([]byte(dataF1), dataScript)
	//t.Logf("result:\"%v\"", string(result))

	//dataF2 is CBABAC
	assert.Equal([]byte(dataF2), result)
}

func TestScriptStringTable(t *testing.T) {
	assert := assert.New(t)

	for i := 0; i < len(dataStrings) - 1; i += 2 {
		f1 := []byte(dataStrings[i])
		f2 := []byte(dataStrings[i+1])

		Backtrack = backtrack
		script := GenScript(f1, f2)
		result := RunScript(f1, script)

		assert.Equalf(f2, result, "random data (t1: \"%s\", t2: \"%s\") failed, scriptProduced: \"%s\" scriptResult: \"%s\"", f1, f2, script, result)
	}

}

func TestScriptRandomized(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 3; i++ {
		testf1 := randByteArray()
		testf2 := randByteArray()

		Backtrack = backtrack
		script := GenScript(testf1, testf2)
		result := RunScript(testf1, script)

		assert.Equalf(testf2, result, "random generated data (t1: \"%s\", t2: \"%s\") failed, scriptProduced: \"%s\"", testf1, testf2, script)
	}

}

func randByteArray() []byte {

	n := rand.Intn(128) + 10
	rString := make([]byte, n)
	rand.Read(rString)

	return rString
}
