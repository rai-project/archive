package archive

import (
	"io"
  "os"
  "fmt"
  "runtime"
	"github.com/docker/docker/pkg/archive"
)

// MimeType ...
func MimeType(opts ...Option) string {
	options := Options{
		format: Config.CompressionFormat,
	}
	for _, o := range opts {
		o(&options)
	}
	switch options.format {
	case archive.Bzip2:
		return "application/bzip2"
	case archive.Gzip:
		return "application/x-gzip"
	case archive.Xz:
		return "application/x-xz"
	default:
		return "application/x-gzip"
	}
}

// Extension ...
func Extension(opts ...Option) string {
	options := Options{
		format: Config.CompressionFormat,
	}
	for _, o := range opts {
		o(&options)
	}
	return options.format.Extension()
}

// DecompressStream ...
func DecompressStream(reader io.Reader) (io.ReadCloser, error) {
	return archive.DecompressStream(reader)
}

// CompressStream ...
func CompressStream(dest io.Writer, opts ...Option) (io.WriteCloser, error) {
	options := Options{
		format: Config.CompressionFormat,
	}
	for _, o := range opts {
		o(&options)
	}
	return archive.CompressStream(dest, options.format)
}

// CanonicalTarNameForPath ...
func CanonicalTarNameForPath(path string) (string, error) {
  if runtime.GOOS != "windows" {
    return path, nil
  }
  // windows: convert windows style relative path with backslashes
	// into forward slashes. Since windows does not allow '/' or '\'
	// in file names, it is mostly safe to replace however we must
	// check just in case
	if strings.Contains(p, "/") {
		return "", fmt.Errorf("Windows path contains forward slash: %s", p)
	}
	return strings.Replace(p, string(os.PathSeparator), "/", -1), nil
}

// Zip ...
func Zip(path string, opts ...Option) (io.ReadCloser, error) {
	options := Options{
		includeSourceDir: false,
		format:           Config.CompressionFormat,
	}
	for _, o := range opts {
		o(&options)
	}
	return archive.TarWithOptions(path, &archive.TarOptions{
		Compression:      options.format,
		IncludeSourceDir: options.includeSourceDir,
		ExcludePatterns: []string{
			"*.git",
		},
	})
}

// Unzip ...
func Unzip(tarArchive io.Reader, destPath string, opts ...Option) error {
	options := Options{
		includeSourceDir: false,
		format:           Config.CompressionFormat,
	}
	for _, o := range opts {
		o(&options)
	}
	return archive.Untar(tarArchive,
		destPath,
		&archive.TarOptions{
			Compression:      options.format,
			IncludeSourceDir: options.includeSourceDir,
			NoLchown:         true,
		},
	)
}

// reads the content of src into a temporary file, and returns the contents
// of that file as an archive. The archive can only be read once - as soon as reading completes,
// the file will be deleted.
func ZipToArchive(src io.Reader) (*archive.TempArchive, error) {
	return archive.NewTempArchive(src, Config.TempDir)
}
