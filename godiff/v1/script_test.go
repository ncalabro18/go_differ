package godiff

import "testing"

func TestGenScript(t *testing.T) {

	script := GenScript([]byte(dataF1), []byte(dataF2), dataBacktrack)
	t.Logf("\"%v\"\n", string(script[4:]))

}
