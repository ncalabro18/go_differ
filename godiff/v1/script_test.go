package godiff

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

var (
	dataScript = []byte("-A-B C+B A B-B A+C")
)

func TestGenScript(t *testing.T) {
	assert := assert.New(t)

	script := GenScript([]byte(dataF1), []byte(dataF2), dataBacktrack)
	scriptData := script[4:]
	//t.Logf("\"%v\"\n", string(scriptData))
	assert.Equal(scriptData, dataScript, "test script failed to match expected data")
}

func TestRunScript(t *testing.T) {
	assert := assert.New(t)

	result := RunScript(dataScript)
	//t.Logf("result:\"%v\"", string(result))

	//dataF2 is CBABAC
	assert.Equal([]byte(dataF2), result)
}


func TestRunScriptRandomized(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 20; i++ {
		t1 := randByteArray()
		t2 := randByteArray()

		script := GenScript(t1, t2, Backtrack(t1, t2))
		result := RunScript(script)

		assert.Equalf(t2, result, "random results failed t1: \"%s\", t2: \"%s\"", t1, t2)
	}


}

var (

)

func randByteArray() []byte {

	n := rand.Intn(128) + 10
	rString := make([]byte, n)
	rand.Read(rString)

	return rString
}