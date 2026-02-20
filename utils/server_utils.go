package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"strings"
	"web-blog/model"
)

func ParseTemplates(templateName string) *template.Template {
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s", templateName))
	if err != nil {
		panic(fmt.Sprintf("error parsing template %s, %v", templateName, err))
	}
	return tmpl
}

func GetMaxArticleID(folderPath string) int {
	files, _ := os.ReadDir(folderPath)
	maxID := 0

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			var id int
			fmt.Sscanf(file.Name(), "article%d.json", &id)
			if id > maxID {
				maxID = id
			}
		}
	}
	return maxID
}

func CreateArticleByFilePath(folderPath string, a model.Article) error {
	filePath := fmt.Sprintf("%s/article%d.json", folderPath, a.ID)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	json.NewEncoder(file).Encode(a)
	return nil
}
