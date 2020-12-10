package godiff

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaveLoadDelta(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(time.Now().Unix())

	path := "../../test/data/save_load_test.delta"

	for i := 0; i < 20; i++ {
		f1 := randByteArray()
		f2 := randByteArray()

		Backtrack = backtrack
		d := NewDelta(f1, f2)
		f2Result1 := d.F2()
		assert.Equal(f2, f2Result1, "f2Result1 == f2")

		err := d.Save(path)
		assert.Nil(err, "d.Save error")

		d, err = Load(f1, path)
		assert.Nil(err, "Load error")

		f2Result2 := d.F2()
		assert.Equal(f2, f2Result2, "f2Result2 == f2")


	}

	os.Remove(path)
}
