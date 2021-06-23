package helpers

import (
	"blog/admin/log"
	_ "log"
	"path/filepath"
)

func Include(path string) []string {
	files, err := filepath.Glob("site/views/templates/*.html")
	if err != nil {
		log.LogJson("Error","Helpers","Include","Could not pull html files in templates folder!!",err.Error())
	}

	pathFiles,_ := filepath.Glob("site/views/"+path+"*.html")
	if err != nil {
		log.LogJson("Error","Helpers","Include","Could not pull html files in templates folder!!",err.Error())
	}

	for _, file := range pathFiles {
		files = append(files, file)
	}

	return files
}
