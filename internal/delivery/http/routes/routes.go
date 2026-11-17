package routes

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"portfolio/internal/delivery/http/handler"
)

// CustomHTMLRender implements Gin's render.HTMLRender interface
// to allow isolated template parsing for each page.
type CustomHTMLRender struct {
	templates map[string]*template.Template
}

// CustomHTMLInstance renders an HTML template with layout.
type CustomHTMLInstance struct {
	Template *template.Template
	Data     any
}

func (r CustomHTMLRender) Instance(name string, data any) render.Render {
	return CustomHTMLInstance{
		Template: r.templates[name],
		Data:     data,
	}
}

func (i CustomHTMLInstance) Render(w http.ResponseWriter) error {
	i.WriteContentType(w)
	// Execute layout template "base.html"
	return i.Template.ExecuteTemplate(w, "base.html", i.Data)
}

func (i CustomHTMLInstance) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"text/html; charset=utf-8"}
	}
}

// createTemplateRenderer builds isolated template sets for each page.
func createTemplateRenderer(webDir string) render.HTMLRender {
	funcMap := template.FuncMap{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.ReplaceAll(template.HTMLEscapeString(text), "\n", "<br>"))
		},
		"hasPrefix": strings.HasPrefix,
		"join":      strings.Join,
		"sub": func(a, b int) int {
			return a - b
		},
	}

	layoutPath := filepath.Join(webDir, "templates", "layouts", "base.html")
	navbarPath := filepath.Join(webDir, "templates", "components", "navbar.html")
	footerPath := filepath.Join(webDir, "templates", "components", "footer.html")
	cardPath := filepath.Join(webDir, "templates", "components", "project_card.html")

	commonFiles := []string{layoutPath, navbarPath, footerPath, cardPath}

	pages := []string{
		"index.html",
		"about.html",
		"experience.html",
		"projects.html",
		"project_detail.html",
		"system_modeling.html",
		"contact.html",
	}

	templates := make(map[string]*template.Template)

	for _, page := range pages {
		pagePath := filepath.Join(webDir, "templates", "pages", page)
		// Check if page file exists before parsing
		if _, err := os.Stat(pagePath); err == nil {
			files := append([]string{}, commonFiles...)
			files = append(files, pagePath)

			t := template.Must(template.New(page).Funcs(funcMap).ParseFiles(files...))
			templates[page] = t
		}
	}

	return CustomHTMLRender{templates: templates}
}

// SetupRouter initializes the Gin router, static assets, and route definitions.
func SetupRouter(h *handler.PortfolioHandler, webDir string) *gin.Engine {
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)

	// Set custom multi-template renderer
	r.HTMLRender = createTemplateRenderer(webDir)

	// Serve static files (CSS, JS, Images, Assets)
	staticDir := filepath.Join(webDir, "static")
	r.Static("/static", staticDir)

	// Setup routes according to Section 6 of Readme.md
	r.GET("/", h.Home)
	r.GET("/about", h.About)
	r.GET("/experience", h.Experience)
	r.GET("/projects", h.Projects)
	r.GET("/projects/:id", h.ProjectDetail)
	r.GET("/modeling", h.SystemModeling)
	r.GET("/system-modeling", h.SystemModeling)
	r.GET("/contact", h.Contact)
	r.POST("/contact", h.ContactSubmit)

	return r
}
