package archive

import (
	"io/ioutil"
	"strings"

	"os"

	"github.com/docker/docker/pkg/archive"
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type archiveConfig struct {
	TempDir                 string              `json:"temp_dir" config"app.temp_dir" default:"default"`
	CompressionFormatString string              `json:"compression_format" config:"archive.format" default:"bzip2"`
	CompressionFormat       archive.Compression `json:"-" config:"-"`
}

var (
	Config = &archiveConfig{}
)

func (archiveConfig) ConfigName() string {
	return "Archive"
}

func (archiveConfig) SetDefaults() {
}

func (a *archiveConfig) Read() {
	vipertags.Fill(a)
	if a.TempDir == "" || a.TempDir == "default" {
		pth, err := ioutil.TempDir("", config.App.Name)
		if err != nil {
			pth = os.TempDir()
		}
		a.TempDir = pth
	}
	switch strings.TrimRight(strings.ToLower(a.CompressionFormatString), "tar") {
	case "xz":
		a.CompressionFormat = archive.Xz
	case "gzip":
		a.CompressionFormat = archive.Gzip
	case "bzip2":
		a.CompressionFormat = archive.Bzip2
	default:
		a.CompressionFormat = archive.Bzip2
	}
}

func (c archiveConfig) String() string {
	return pp.Sprintln(c)
}

func (c archiveConfig) Debug() {
	Debug("Archive Config = ", c)
}

func init() {
	config.Register(Config)
}
