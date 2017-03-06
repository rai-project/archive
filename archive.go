package archive

import (
	"io"

	"github.com/docker/docker/pkg/archive"
)

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

func DecompressStream(reader io.Reader) (io.ReadCloser, error) {
	return archive.DecompressStream(reader)
}

func CompressStream(dest io.Writer, opts ...Option) (io.WriteCloser, error) {
	options := Options{
		format: Config.CompressionFormat,
	}
	for _, o := range opts {
		o(&options)
	}
	return archive.CompressStream(dest, options.format)
}

func Zip(path string, opts ...Option) (io.ReadCloser, error) {
	options := Options{
		format: Config.CompressionFormat,
	}
	for _, o := range opts {
		o(&options)
	}
	return archive.TarWithOptions(path, &archive.TarOptions{
		IncludeSourceDir: false,
		Compression:      options.format,
		ExcludePatterns: []string{
			"*.git",
		},
	})
}

func Unzip(tarArchive io.Reader, destPath string, opts ...Option) error {
	options := Options{
		format: Config.CompressionFormat,
	}
	for _, o := range opts {
		o(&options)
	}
	return archive.Untar(tarArchive,
		destPath,
		&archive.TarOptions{
			Compression:      options.format,
			IncludeSourceDir: true,
		},
	)
}

// reads the content of src into a temporary file, and returns the contents
// of that file as an archive. The archive can only be read once - as soon as reading completes,
// the file will be deleted.
func ZipToArchive(src io.Reader) (*archive.TempArchive, error) {
	return archive.NewTempArchive(src, Config.TempDir)
}
