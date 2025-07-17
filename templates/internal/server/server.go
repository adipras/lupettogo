package server

import (
	"log"
	"net/http"

	"{{.ProjectName}}/internal/config"
	"{{.ProjectName}}/internal/database"
	"{{.ProjectName}}/internal/handlers"
	"{{.ProjectName}}/internal/middleware"
	"{{.ProjectName}}/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
	config *config.Config
}

func New(cfg *config.Config) *Server {
	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize database
	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Printf("Warning: Failed to connect to database: %v", err)
		db = nil
	}

	// Run migrations if database is connected
	if db != nil {
		if err := database.Migrate(db); err != nil {
			log.Printf("Warning: Failed to run migrations: %v", err)
		}
	}

	// Initialize services
	services := services.New(db)

	// Initialize handlers
	handlers := handlers.New(services)

	// Initialize router
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// Setup routes
	setupRoutes(router, handlers, cfg)

	return &Server{
		router: router,
		db:     db,
		config: cfg,
	}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func setupRoutes(r *gin.Engine, h *handlers.Handlers, cfg *config.Config) {
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "🐺 LupettoGo API is running",
			"version": cfg.API.Version,
		})
	})

	// API routes
	api := r.Group("/api/" + cfg.API.Version)
	{
		// Add your API routes here
		api.GET("/example", h.Example.GetExample)
	}
}
