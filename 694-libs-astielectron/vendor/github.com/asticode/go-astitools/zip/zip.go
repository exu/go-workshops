package astizip

import (
	"archive/zip"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/asticode/go-astitools/io"
)

// Zip zips a src into a dst
// dstRoot can be used to create root directories so that the content is zipped in /path/to/zip.zip/root/path
func Zip(ctx context.Context, src, dst, dstRoot string) (err error) {
	// Create destination file
	var dstFile *os.File
	if dstFile, err = os.Create(dst); err != nil {
		return err
	}
	defer dstFile.Close()

	// Create zip writer
	var zw = zip.NewWriter(dstFile)
	defer zw.Close()

	// Walk
	filepath.Walk(src, func(path string, info os.FileInfo, e1 error) (e2 error) {
		// Process error
		if e1 != nil {
			return e1
		}

		// Init header
		var h *zip.FileHeader
		if h, e2 = zip.FileInfoHeader(info); e2 != nil {
			return
		}

		// Set header info
		h.Name = filepath.Join(dstRoot, strings.TrimPrefix(path, src))
		if info.IsDir() {
			h.Name += "/"
		} else {
			h.Method = zip.Deflate
		}

		// Create writer
		var w io.Writer
		if w, e2 = zw.CreateHeader(h); e2 != nil {
			return
		}

		// If path is dir, stop here
		if info.IsDir() {
			return
		}

		// Open path
		var walkFile *os.File
		if walkFile, e2 = os.Open(path); e2 != nil {
			return
		}
		defer walkFile.Close()

		// Copy
		if _, e2 = astiio.Copy(ctx, walkFile, w); e2 != nil {
			return
		}
		return
	})
	return
}
