package util

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

func EnsureDir(elems ...string) error {
	fp := filepath.Join(elems...)
	fpInfo, err := os.Stat(fp)
	if err != nil {
		if err := os.MkdirAll(fp, 0755); err != nil {
			slog.Error("cannot create dir", slog.String("error", err.Error()))
			return err
		}
	} else {
		if !fpInfo.IsDir() {
			return fmt.Errorf("expecting directory")
		}
	}
	return nil
}
