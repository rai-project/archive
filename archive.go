package archive

import (
	"io"

	"github.com/docker/docker/pkg/archive"
)

func DecompressStream(reader io.Reader) (io.ReadCloser, error) {
	return archive.DecompressStream(reader)
}

func CompressStream(dest io.Writer) (io.WriteCloser, error) {
	return archive.CompressStream(dest, Config.CompressionFormat)
}

func Zip(path string) (io.ReadCloser, error) {
	return archive.TarWithOptions(path, &archive.TarOptions{
		IncludeSourceDir: true,
		Compression:      Config.CompressionFormat,
		ExcludePatterns: []string{
			"*.git",
		},
	})
}

func Unzip(tarArchive io.Reader, destPath string) error {
	return archive.Untar(tarArchive,
		destPath,
		&archive.TarOptions{
			Compression:      Config.CompressionFormat,
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
