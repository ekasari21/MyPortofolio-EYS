package main

import (
	"fmt"
	"log"
	"os"

	"portfolio/internal/delivery/http/handler"
	"portfolio/internal/delivery/http/routes"
	jsonPersistence "portfolio/internal/infrastructure/persistence/json"
	"portfolio/internal/usecase"
)

func main() {
	// Base directories
	dataDir := "data"
	webDir := "web"

	// 1. Initialize Infrastructure Layer (JSON Persistence)
	profileRepo := jsonPersistence.NewProfileRepository(dataDir)
	experienceRepo := jsonPersistence.NewExperienceRepository(dataDir)
	educationRepo := jsonPersistence.NewEducationRepository(dataDir)
	skillRepo := jsonPersistence.NewSkillRepository(dataDir)
	projectRepo := jsonPersistence.NewProjectRepository(dataDir)
	diagramRepo := jsonPersistence.NewDiagramRepository(dataDir)

	// 2. Initialize Use Case Layer
	profileUC := usecase.NewProfileUseCase(profileRepo)
	experienceUC := usecase.NewExperienceUseCase(experienceRepo)
	educationUC := usecase.NewEducationUseCase(educationRepo)
	skillUC := usecase.NewSkillUseCase(skillRepo)
	projectUC := usecase.NewProjectUseCase(projectRepo)
	diagramUC := usecase.NewDiagramUseCase(diagramRepo)

	// 3. Initialize Delivery Layer (HTTP Handler)
	portfolioHandler := handler.NewPortfolioHandler(
		profileUC,
		experienceUC,
		educationUC,
		skillUC,
		projectUC,
		diagramUC,
	)

	// 4. Setup Gin HTTP Router
	router := routes.SetupRouter(portfolioHandler, webDir)

	// Determine port (default 8080 per Readme.md)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	fmt.Println("====================================================")
	fmt.Printf("🚀 Portfolio Server is running at: http://localhost:%s\n", port)
	fmt.Println("Built with Golang, Gin Framework & Clean Architecture")
	fmt.Println("====================================================")

	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
