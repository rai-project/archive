package archive

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/GeertJohan/go-sourcepath"
	"github.com/Unknwon/com"
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
	fixturesDir := suite.fixturesDir

	in, err := os.Open(fixturesDir)
	assert.NoError(t, err, "failed to open fixturesDir=%v", fixturesDir)
	defer in.Close()

	_, err = CompressStream(in)
	assert.NoError(t, err, "failed to compress stream")
}

func (suite *ArchiveTestSuite) TestZip() {
	t := suite.T()
	fixturesDir := suite.fixturesDir

	w, err := Zip(fixturesDir)
	assert.NoError(t, err, "failed to compress stream fixturesDir=%v", fixturesDir)

	dest := filepath.Join(suite.tmpDir, "testzip.tar.gz")
	tmpDir, err := os.Create(dest)
	assert.NoError(t, err, "failed to open dest=%v", dest)
	defer tmpDir.Close()

	_, err = io.Copy(tmpDir, w)
	assert.NoError(t, err, "failed to copy data to dest=%v", dest)

	assert.Equal(t, true, com.IsFile(dest), "dest exists")

	sz, err := com.FileSize(dest)
	assert.NoError(t, err, "failed to get file size for dest=%v", dest)
	assert.NotEqual(t, 0, sz, "file size must not be 0")
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
	tmpDir := Config.TempDir
	if Config.TempDir == "" {
		if t, err := ioutil.TempDir("", "raiarchive"); err == nil {
			tmpDir = t
		}
	}
	t.Logf("using %v as temp dir", tmpDir)
	suite.Run(
		t,
		&ArchiveTestSuite{
			fixturesDir: filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures"),
			tmpDir:      tmpDir,
		},
	)
}

func TestMain(m *testing.M) {
	os.Setenv("DEBUG", "TRUE")
	os.Setenv("VERBOSE", "TRUE")
	config.Init()
	os.Exit(m.Run())
}
