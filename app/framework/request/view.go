package request

import (
	"net/http"
	"html/template"
	"strings"
)

func View(w http.ResponseWriter, content string, data interface{}) {
	ViewRunner(w, "layouts/header", nil, content, data, "layouts/footer", nil)
}

func ViewRunner(w http.ResponseWriter, headerLayout string, headerData interface{},
	content string, contentData interface{},
	footerLayout string, footerData interface{}) {
	var allFiles []string
	content = content + ".html"

	allFiles = append(allFiles, "app/presenters/views/"+headerLayout+".html")
	allFiles = append(allFiles, "app/presenters/views/"+content)
	allFiles = append(allFiles, "app/presenters/views/"+footerLayout+".html")

	t, _ := template.ParseFiles(allFiles...)
	header := t.Lookup("header.html")
	header.ExecuteTemplate(w, "header", headerData)

	splitContent := strings.Split(content, "/")
	body := t.Lookup(splitContent[len(splitContent)-1])
	body.ExecuteTemplate(w, "content", contentData)

	footer := t.Lookup("footer.html")
	footer.ExecuteTemplate(w, "footer", footerData)
}
