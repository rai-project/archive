package archive

import (
	"strings"

	"github.com/docker/docker/pkg/archive"
)

type Options struct {
	includeSourceDir bool
	format           archive.Compression
}

type Option func(*Options)

func Format(format string) Option {
	return func(o *Options) {
		o.format = parseFormat(format)
	}
}

func BZip2Format() Option {
	return Format("bzip2")
}

func XZFormat() Option {
	return Format("xz")
}

func GZipFormat() Option {
	return Format("gzip")
}

func IncludeSourceDir(b bool) Option {
	return func(o *Options) {
		o.includeSourceDir = b
	}
}

func parseFormat(format string) archive.Compression {
	switch strings.TrimLeft(strings.ToLower(format), "tar.") {
	case "xz":
		return archive.Xz
	case "gzip", "gunzip":
		return archive.Gzip
	case "bzip", "bzip2":
		return archive.Bzip2
	default:
		return archive.Gzip
	}
}
