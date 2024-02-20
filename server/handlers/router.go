package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/vavilen84/nft-project/constants"
	"log"
	"net/http"
	"os"
	"time"
)

func MakeHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	v1Router := BuildV1Paths()
	r.Mount("/api/v1", v1Router)

	return r
}

func BuildV1Paths() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/security", func(r chi.Router) {
		c := SecurityController{}

		r.Post("/two-fa-login-step-one", c.TwoFaLoginStepOne)
		r.Post("/two-fa-login-step-two", c.TwoFaLoginStepTwo)
		r.Post("/register", c.Register)
		r.Post("/forgot-password", c.ForgotPassword)
		r.Post("/reset-password", c.ResetPassword)
		r.Get("/verify-email", c.VerifyEmail)

		r.With(UserAuth).Post("/change-password", c.ChangePassword)
	})
	r.Route("/drop", func(r chi.Router) {
		c := DropController{}

		r.With(UserAuth).Post("/create", c.Create)
		r.With(UserAuth).Post("/update", c.Update)
		r.With(UserAuth).Post("/upload-preview-image", c.UploadPreviewImage)
	})
	return r
}

func InitHttpServer(handler http.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	return &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: constants.DefaultWriteTimout,
		ReadTimeout:  constants.DefaultReadTimeout,
	}
}
