package main

import (
	"cv-builder-api/config"
	"cv-builder-api/internal/delivery"
	"cv-builder-api/internal/repository"
	"cv-builder-api/internal/usecase"
	customMw "cv-builder-api/pkg/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.LoadConfig()

	config.ConnectDatabase(cfg)

	userRepo := repository.NewUserRepository(config.DB)
	userUsecase := usecase.NewUserUsecase(userRepo, cfg.JWTSecret)
	userHandler := delivery.NewUserHandler(userUsecase)
	profileRepo := repository.NewProfileRepository(config.DB)
	profleUsecase := usecase.NewProfileUsecase(profileRepo)
	profileHandler := delivery.NewProfileHandler(profleUsecase)
	workExperienceRepo := repository.NewWorkExperienceRepository(config.DB)
	workExperienceUsecase := usecase.NewWorkExperienceUsecase(workExperienceRepo)
	workExperienceHandler := delivery.NewWorkExperienceHandler(workExperienceUsecase)
	skillRepo := repository.NewSkillsRepository(config.DB)
	skillUsecase := usecase.NewSkillsUsecase(skillRepo)
	skillHandler := delivery.NewskillHandler(skillUsecase)
	projectRepo := repository.NewProjectRepository(config.DB)
	projectUsecase := usecase.NewProjectUsecase(projectRepo)
	projectHandler := delivery.NewProjectHandler(projectUsecase)
	educationRepo := repository.NewEducationRepository(config.DB)
	educationUsecase := usecase.NewEducationUsecase(educationRepo)
	educationHandler := delivery.NewEducationHandler(educationUsecase)
	cvRepo := repository.NewCVRepository(config.DB)
	cvUsecase := usecase.NewCVUsecase(cvRepo)
	cvHandler := delivery.NewCVHandler(cvUsecase)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 5. Mendaftarkan Rute (Endpoints)
	r.Route("/api/v1", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Post("/register", userHandler.Register)
			r.Post("/login", userHandler.Login)
		})

		r.Group(func(r chi.Router) {

			r.Use(customMw.AuthMiddleware(cfg.JWTSecret))

			r.Get("/profile", profileHandler.GetProfile)
			r.Post("/profile", profileHandler.UpsertProfile)

			r.Get("/work-experience", workExperienceHandler.GetAllByUserID)
			r.Post("/work-experience", workExperienceHandler.Create)
			r.Put("/work-experience/{id}", workExperienceHandler.Update)
			r.Delete("/work-experience/{id}", workExperienceHandler.Delete)

			r.Get("/skill", skillHandler.GetAllByUserID)
			r.Post("/skill", skillHandler.Create)
			r.Put("/skill/{id}", skillHandler.Update)
			r.Delete("/skill/{id}", skillHandler.Delete)

			r.Get("/project", projectHandler.GetAllByUserID)
			r.Post("/project", projectHandler.Create)
			r.Put("/project/{id}", projectHandler.Update)
			r.Delete("/project/{id}", projectHandler.Delete)

			r.Get("/education", educationHandler.GetAllByUserID)
			r.Post("/education", educationHandler.Create)
			r.Put("/education/{id}", educationHandler.Update)
			r.Delete("/education/{id}", educationHandler.Delete)

			r.Get("/generate-cv", cvHandler.GenerateCV)

		})
	})

	http.ListenAndServe(":8080", r)
}
