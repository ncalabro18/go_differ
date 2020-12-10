package godiff

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	t.Logf("result:\"%v\"", string(result))
	//dataF2 is CBABAC
	assert.Equal([]byte(dataF2), result)


}