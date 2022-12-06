package dedup

import "dedup/dirscan"

func FilterDuplicateFiles(files dirscan.Files) (dirscan.DuplicateFiles, int) {
	dedupped := make(dirscan.Files)
	duplicates := make(dirscan.DuplicateFiles)

	for _, file := range files {
		name := file.SanitisedName()
		deduppedFile, isDuplicate := dedupped[name]

		if !isDuplicate {
			dedupped[name] = file
			continue
		}

		if len(duplicates[file.Name]) == 0 {
			duplicates[file.Name] = make(dirscan.Files)
		}

		duplicates[file.Name][deduppedFile.Path] = deduppedFile
		duplicates[file.Name][file.Path] = file
	}

	return duplicates, len(duplicates)
}
