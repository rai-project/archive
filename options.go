package archive

import (
	"strings"

	"github.com/docker/docker/pkg/archive"
)

// Options ...
type Options struct {
	includeSourceDir bool
	format           archive.Compression
}

// Option ...
type Option func(*Options)

// Format ...
func Format(format string) Option {
	return func(o *Options) {
		o.format = parseFormat(format)
	}
}

// BZip2Format ...
func BZip2Format() Option {
	return Format("bzip2")
}

// XZFormat ...
func XZFormat() Option {
	return Format("xz")
}

// GZipFormat ...
func GZipFormat() Option {
	return Format("gzip")
}

// IncludeSourceDir ...
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
