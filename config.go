package archive

import (
	"io/ioutil"

	"os"

	"github.com/docker/docker/pkg/archive"
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type archiveConfig struct {
	TempDir                 string              `json:"temp_dir" config"app.temp_dir" default:"default"`
	CompressionFormatString string              `json:"compression_format" config:"archive.format" default:"gzip"`
	CompressionFormat       archive.Compression `json:"-" config:"-"`
	done                    chan struct{}       `json:"-" config:"-"`
}

// Config ...
var (
	Config = &archiveConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (archiveConfig) ConfigName() string {
	return "Archive"
}

// SetDefaults ...
func (a *archiveConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *archiveConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.TempDir == "" || a.TempDir == "default" {
		pth, err := ioutil.TempDir("", config.App.Name)
		if err != nil {
			pth = os.TempDir()
		}
		a.TempDir = pth
	}
	a.CompressionFormat = parseFormat(a.CompressionFormatString)
}

// Wait ...
func (c archiveConfig) Wait() {
	<-c.done
}

// String ...
func (c archiveConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c archiveConfig) Debug() {
	pp.Println("Archive Config = ", c)
}

func init() {
	config.Register(Config)
}
