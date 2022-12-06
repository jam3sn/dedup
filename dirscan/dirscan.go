package dirscan

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
	"time"
)

var ignoreList = [1]string{".DS_Store"}

type File struct {
	Name      string
	Extension string
	Path      string
	Size      int64
	Modified  time.Time
}

func (f File) ReadableSize() string {
	return byteCountDecimal(f.Size)
}

func (f File) SanitisedName() (name string) {
	name = strings.ToLower(f.Name)
	name = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(name, "")
	return
}

type Files map[string]File
type DuplicateFiles map[string]Files

func GetFiles(scanPath string) (Files, error) {
	fileSystem := os.DirFS(scanPath)
	files := make(map[string]File)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if excludeEntry(d) {
			return nil
		}

		nameParts := strings.Split(d.Name(), ".")
		name := strings.Join(nameParts[:1], "")
		fileInfo, _ := d.Info()

		files[path] = File{
			Name:      name,
			Extension: strings.Join(nameParts[len(nameParts)-1:], ""),
			Path:      path,
			Size:      fileInfo.Size(),
			Modified:  fileInfo.ModTime(),
		}

		return nil
	})

	return files, nil
}

func excludeEntry(d fs.DirEntry) bool {
	if !d.Type().IsRegular() {
		return true
	}

	for _, ignore := range ignoreList {
		if d.Name() == ignore {
			return true
		}
	}

	return false
}

// Human readable file sizes, copied.
// [Source]: https://programming.guide/go/formatting-byte-size-to-human-readable-format.html
func byteCountDecimal(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
