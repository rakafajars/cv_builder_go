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

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 5. Mendaftarkan Rute (Endpoints)
	r.Route("/api/v1", func(r chi.Router) {

		// --- AREA PUBLIK (Tidak butuh token) ---
		r.Group(func(r chi.Router) {
			r.Post("/register", userHandler.Register)
			r.Post("/login", userHandler.Login)
		})

		// --- AREA TERLARANG (Wajib bawa token JWT) ---
		r.Group(func(r chi.Router) {
			// Pasang "Satpam" buatan kita khusus untuk grup rute ini
			r.Use(customMw.AuthMiddleware(cfg.JWTSecret))
			// Endpoint sementara untuk mengetes apakah token dan middleware berfungsi

			r.Get("/profile", profileHandler.GetProfile)
			r.Post("/profile", profileHandler.UpsertProfile)

			r.Get("/work-experience", workExperienceHandler.GetAllByUserID)
			r.Post("/work-experience", workExperienceHandler.Create)
			r.Put("/work-experience/{id}", workExperienceHandler.Update)
			r.Delete("/work-experience/{id}", workExperienceHandler.Delete)

		})
	})

	http.ListenAndServe(":8080", r)
}
