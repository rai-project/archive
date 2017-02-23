package archive

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/GeertJohan/go-sourcepath"
	"github.com/rai-project/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArchiveTestSuite struct {
	fixturesDir string
	tmpDir      string
	suite.Suite
}

func (suite *ArchiveTestSuite) TestCompressStream() {
	t := suite.T()
	tmpDir := suite.tmpDir

	path := filepath.Join(tmpDir, "dest")
	dest, err := os.Create(path)
	if err != nil {
		t.Fatalf("Fail to create the destination file")
	}
	defer dest.Close()
	defer os.Remove(path)

	_, err = CompressStream(dest)
	assert.NoError(t, err, "failed to compress stream")
}

func (suite *ArchiveTestSuite) TestDecompressStream() {
	t := suite.T()
	fixturesDir := suite.fixturesDir

	zippedReader, err := Zip(fixturesDir)
	if !assert.NoError(t, err, "failed to create zip") {
		return
	}
	defer zippedReader.Close()

	decompressedReader, err := DecompressStream(zippedReader)
	if !assert.NoError(t, err, "failed to decompress stream") {
		return
	}
	defer decompressedReader.Close()
}

func TestArchive(t *testing.T) {
	suite.Run(
		t,
		&ArchiveTestSuite{
			fixturesDir: filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures"),
			tmpDir:      Config.TempDir,
		},
	)
}

func TestMain(m *testing.M) {
	os.Setenv("DEBUG", "TRUE")
	os.Setenv("VERBOSE", "TRUE")
	config.Init()
	os.Exit(m.Run())
}
