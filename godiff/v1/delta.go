package godiff

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	headerBytes = []byte{0xfa, 0xca, 0xab, 0x3e}
)

type Delta struct {
	Script []byte
	F1     []byte
}

//NewDelta generates a delta script that can restore f2 givin f1
func NewDelta(f1, f2 []byte) Delta {
	d := Delta{
		Script: GenScript(f1, f2),
		F1:     f1,
	}
	return d
}

//Save stores the script data into an opened file at scriptPath. If f1 is not also saved, the data could be lost.
func (d *Delta) Save(scriptFilePath string) error {
	data := []byte{}
	data = append(data, headerBytes...)
	data = append(data, d.Script...)

	file, err := os.Create(scriptFilePath)

	if err != nil {
		return fmt.Errorf("godiff.SaveDelta: %v", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("godiff.SaveDelta: %v", err)
	}

	file.Close()

	return nil
}

//Load
func Load(f1 []byte, scriptFilePath string) (Delta, error) {
	d := Delta{}

	//loads the data from scriptFilePath and checks the headerBytes
	scriptFile, err := os.Open(scriptFilePath)
	if err != nil {
		return d, fmt.Errorf("godiff.LoadDelta: %v", err)
	}
	scriptData, err := ioutil.ReadAll(scriptFile)
	if err != nil {
		return d, fmt.Errorf("godiff.LoadDelta: %v", err)
	}
	if !isHeader(scriptData[:4]) {
		return d, fmt.Errorf("godiff.LoadDelta: header bytes not recognized")
	}

	d.F1 = f1
	d.Script = scriptData[4:]
	return d, nil
}

//F2 executes the script and returns the original f2 []byte.
func (d *Delta) F2() []byte {
	return RunScript(d.F1, d.Script)
}

func isHeader(hBytes []byte) bool {
	for i := 0; i < len(headerBytes); i++ {
		if headerBytes[i] != hBytes[i] {
			return false
		}
	}
	return true
}

