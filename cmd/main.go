package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm" 

	"go-auth/internal/entities"
	"go-auth/internal/frameworks/database"
	"go-auth/internal/frameworks/jwt"
	"go-auth/internal/interfaces/handlers"
	"go-auth/internal/interfaces/repositories"
	"go-auth/internal/usecases"
)

func main() {
	// Initialize dependencies
	db := initDatabase()
	userRepo := repositories.NewUserRepository(db)
	jwtService := jwt.NewJWTService()
	authUseCase := usecases.NewAuthUseCase(userRepo, jwtService)
	authHandler := handlers.NewAuthHandler(authUseCase)


	r := setupRouter(authHandler)

	startServer(r)
}

// Initializing database connection
func initDatabase() *gorm.DB {
	db, err := database.NewPostgresDB() 
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&entities.User{})
	log.Println("✅ Database connected & migrated")
	return db
}

func setupRouter(authHandler *handlers.AuthHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/auth/google/login", authHandler.GoogleLogin).Methods("GET")
	r.HandleFunc("/auth/google/callback", authHandler.GoogleCallback).Methods("GET")
	r.HandleFunc("/protected", authHandler.Protected).Methods("GET")
	return r
}

func startServer(router *mux.Router) {
	serverAddr := getEnv("SERVER_PORT", ":8080")
	server := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		log.Printf("🚀 Server started at %s", serverAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server error: %v", err)
		}
	}()

	waitForShutdown(server)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func waitForShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("⚠️ Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced shutdown: %v", err)
	}

	log.Println("✅ Server stopped gracefully")
}
