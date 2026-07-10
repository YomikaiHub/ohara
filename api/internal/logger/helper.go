package logger

import (
	"log/slog"
	"path/filepath"
	"strings"
)

func replaceSource(groups []string, a slog.Attr) slog.Attr {
	if a.Key != slog.SourceKey {
		return a
	}

	source, ok := a.Value.Any().(*slog.Source)
	if !ok {
		return a
	}

	idx := strings.Index(source.File, "internal")
	if idx != -1 {
		source.File = source.File[idx:]
	} else {
		source.File = filepath.Base(source.File)
	}

	return slog.Any(a.Key, source)
}
