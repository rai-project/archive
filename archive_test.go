package archive

import (
	"testing"

	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/stretchr/testify/assert"
  "github.com/GeertJohan/go-sourcepath"
)

func TestZip(t *testing.T) {
	trgt := "/tmp/test.tar.bz2"
	toZip := filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures")

	assert.True(t, com.IsDir(toZip), "the directory ._fixtures must exist")
	zip, err := Zip(trgt, toZip)

	assert.NoError(t, err, "zipping should not return an error")
	assert.True(t, com.IsFile(zip), "the output file must exist")

	trgtDir, err := Unzip("/tmp/pp", zip)
	assert.NoError(t, err, "zipping should not return an error")
	assert.True(t, com.IsDir(trgtDir), "the target directory must exist")
}
