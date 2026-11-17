package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"portfolio/internal/usecase"
)

// PortfolioHandler handles HTTP requests for the portfolio website.
type PortfolioHandler struct {
	profileUC    usecase.ProfileUseCase
	experienceUC usecase.ExperienceUseCase
	educationUC  usecase.EducationUseCase
	skillUC      usecase.SkillUseCase
	projectUC    usecase.ProjectUseCase
	diagramUC    usecase.DiagramUseCase
}

// NewPortfolioHandler creates a new PortfolioHandler.
func NewPortfolioHandler(
	profileUC usecase.ProfileUseCase,
	experienceUC usecase.ExperienceUseCase,
	educationUC usecase.EducationUseCase,
	skillUC usecase.SkillUseCase,
	projectUC usecase.ProjectUseCase,
	diagramUC usecase.DiagramUseCase,
) *PortfolioHandler {
	return &PortfolioHandler{
		profileUC:    profileUC,
		experienceUC: experienceUC,
		educationUC:  educationUC,
		skillUC:      skillUC,
		projectUC:    projectUC,
		diagramUC:    diagramUC,
	}
}

// Home handles the root page (/).
func (h *PortfolioHandler) Home(c *gin.Context) {
	profile, _ := h.profileUC.GetProfile()
	experiences, _ := h.experienceUC.GetExperiences()
	skills, _ := h.skillUC.GetSkills()
	projects, _ := h.projectUC.GetFeaturedProjects()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"ActivePage":   "home",
		"Title":        "Home - Portfolio",
		"Profile":      profile,
		"Experiences":  experiences,
		"Skills":       skills,
		"Projects":     projects,
	})
}

// About handles the about page (/about).
func (h *PortfolioHandler) About(c *gin.Context) {
	profile, _ := h.profileUC.GetProfile()
	education, _ := h.educationUC.GetEducation()
	certifications, _ := h.educationUC.GetCertifications()
	skills, _ := h.skillUC.GetSkills()

	c.HTML(http.StatusOK, "about.html", gin.H{
		"ActivePage":     "about",
		"Title":          "About Me - Portfolio",
		"Profile":        profile,
		"Education":      education,
		"Certifications": certifications,
		"Skills":         skills,
	})
}

// Experience handles the experience and career page (/experience).
func (h *PortfolioHandler) Experience(c *gin.Context) {
	profile, _ := h.profileUC.GetProfile()
	experiences, _ := h.experienceUC.GetExperiences()
	education, _ := h.educationUC.GetEducation()
	certifications, _ := h.educationUC.GetCertifications()

	c.HTML(http.StatusOK, "experience.html", gin.H{
		"ActivePage":     "experience",
		"Title":          "Experience & Career - Portfolio",
		"Profile":        profile,
		"Experiences":    experiences,
		"Education":      education,
		"Certifications": certifications,
	})
}

// Projects handles the projects showcase page (/projects).
func (h *PortfolioHandler) Projects(c *gin.Context) {
	profile, _ := h.profileUC.GetProfile()
	projects, _ := h.projectUC.GetProjects()

	// Extract unique categories for filter tabs
	categoryMap := make(map[string]bool)
	var categories []string
	for _, p := range projects {
		if !categoryMap[p.Category] && p.Category != "" {
			categoryMap[p.Category] = true
			categories = append(categories, p.Category)
		}
	}

	c.HTML(http.StatusOK, "projects.html", gin.H{
		"ActivePage": "projects",
		"Title":      "Projects & Case Studies - Portfolio",
		"Profile":    profile,
		"Projects":   projects,
		"Categories": categories,
	})
}

// ProjectDetail handles viewing a single project (/projects/:id).
func (h *PortfolioHandler) ProjectDetail(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Redirect(http.StatusFound, "/projects")
		return
	}

	project, err := h.projectUC.GetProjectByID(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/projects")
		return
	}

	profile, _ := h.profileUC.GetProfile()

	c.HTML(http.StatusOK, "project_detail.html", gin.H{
		"ActivePage": "projects",
		"Title":      project.Name + " - Portfolio",
		"Profile":    profile,
		"Project":    project,
	})
}

// SystemModeling handles the system modeling & architecture diagrams page (/modeling).
func (h *PortfolioHandler) SystemModeling(c *gin.Context) {
	profile, _ := h.profileUC.GetProfile()
	experience, _ := h.diagramUC.GetModelingExperience()

	// Extract unique categories for filtering
	categoryMap := make(map[string]bool)
	var categories []string
	if experience != nil {
		for _, d := range experience.Diagrams {
			if !categoryMap[d.Category] && d.Category != "" {
				categoryMap[d.Category] = true
				categories = append(categories, d.Category)
			}
		}
	}

	c.HTML(http.StatusOK, "system_modeling.html", gin.H{
		"ActivePage": "modeling",
		"Title":      "System Modeling & Diagrams - Portfolio",
		"Profile":    profile,
		"Experience": experience,
		"Categories": categories,
	})
}

// Contact handles the contact page (/contact).
func (h *PortfolioHandler) Contact(c *gin.Context) {
	profile, _ := h.profileUC.GetProfile()

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"ActivePage": "contact",
		"Title":      "Contact - Portfolio",
		"Profile":    profile,
	})
}

// ContactSubmit processes contact form submission (POST /contact).
func (h *PortfolioHandler) ContactSubmit(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	subject := c.PostForm("subject")
	message := c.PostForm("message")

	profile, _ := h.profileUC.GetProfile()

	c.HTML(http.StatusOK, "contact.html", gin.H{
		"ActivePage":   "contact",
		"Title":        "Contact - Portfolio",
		"Profile":      profile,
		"Success":      true,
		"SenderName":   name,
		"SenderEmail":  email,
		"Subject":      subject,
		"SentMessage":  message,
	})
}
