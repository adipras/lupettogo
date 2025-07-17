package generator

// Internal package templates

var internalTemplates = map[string]string{
	"internal/config/config.go": `package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   ` + "`" + `mapstructure:"server"` + "`" + `
	Database DatabaseConfig ` + "`" + `mapstructure:"database"` + "`" + `
	JWT      JWTConfig      ` + "`" + `mapstructure:"jwt"` + "`" + `
	API      APIConfig      ` + "`" + `mapstructure:"api"` + "`" + `
}

type ServerConfig struct {
	Port string ` + "`" + `mapstructure:"port"` + "`" + `
	Mode string ` + "`" + `mapstructure:"mode"` + "`" + `
}

type DatabaseConfig struct {
	Host     string ` + "`" + `mapstructure:"host"` + "`" + `
	Port     string ` + "`" + `mapstructure:"port"` + "`" + `
	User     string ` + "`" + `mapstructure:"user"` + "`" + `
	Password string ` + "`" + `mapstructure:"password"` + "`" + `
	Name     string ` + "`" + `mapstructure:"name"` + "`" + `
	Driver   string ` + "`" + `mapstructure:"driver"` + "`" + `
}

type JWTConfig struct {
	Secret    string ` + "`" + `mapstructure:"secret"` + "`" + `
	ExpiresIn string ` + "`" + `mapstructure:"expires_in"` + "`" + `
}

type APIConfig struct {
	Version string ` + "`" + `mapstructure:"version"` + "`" + `
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set environment variable prefix
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Set defaults
	setDefaults()

	// Read config file (optional)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func setDefaults() {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("jwt.expires_in", "24h")
	viper.SetDefault("api.version", "v1")
}`,

	"internal/database/database.go": `package database

import (
	"fmt"
	"log"

	"{{.ProjectName}}/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection(cfg *config.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch cfg.Database.Driver {
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.Database.Host,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Name,
			cfg.Database.Port,
		)
		dialector = postgres.Open(dsn)
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name,
		)
		dialector = mysql.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	// Add your models here for auto-migration
	// Example: db.AutoMigrate(&models.User{}, &models.Product{})
	log.Println("Database migration completed")
	return nil
}`,

	"internal/server/server.go": `package server

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
}`,

	"internal/middleware/cors.go": `package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}`,

	"internal/models/example.go": `package models

import (
	"time"

	"gorm.io/gorm"
)

type Example struct {
	ID        uint           ` + "`" + `json:"id" gorm:"primarykey"` + "`" + `
	Name      string         ` + "`" + `json:"name" gorm:"not null"` + "`" + `
	Email     string         ` + "`" + `json:"email" gorm:"uniqueIndex;not null"` + "`" + `
	Status    string         ` + "`" + `json:"status" gorm:"default:active"` + "`" + `
	CreatedAt time.Time      ` + "`" + `json:"created_at"` + "`" + `
	UpdatedAt time.Time      ` + "`" + `json:"updated_at"` + "`" + `
	DeletedAt gorm.DeletedAt ` + "`" + `json:"-" gorm:"index"` + "`" + `
}

func (Example) TableName() string {
	return "examples"
}`,

	"internal/repositories/repositories.go": `package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	Example *ExampleRepository
}

func New(db *gorm.DB) *Repositories {
	return &Repositories{
		Example: NewExampleRepository(db),
	}
}`,

	"internal/repositories/example_repository.go": `package repositories

import (
	"{{.ProjectName}}/internal/models"
	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *ExampleRepository {
	return &ExampleRepository{
		db: db,
	}
}

func (r *ExampleRepository) FindAll() ([]*models.Example, error) {
	var examples []*models.Example
	err := r.db.Find(&examples).Error
	return examples, err
}

func (r *ExampleRepository) FindByID(id uint) (*models.Example, error) {
	var example models.Example
	err := r.db.First(&example, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &example, nil
}

func (r *ExampleRepository) Create(example *models.Example) (*models.Example, error) {
	err := r.db.Create(example).Error
	return example, err
}

func (r *ExampleRepository) Update(example *models.Example) (*models.Example, error) {
	err := r.db.Save(example).Error
	return example, err
}

func (r *ExampleRepository) Delete(id uint) error {
	return r.db.Delete(&models.Example{}, id).Error
}`,

	"internal/services/services.go": `package services

import (
	"{{.ProjectName}}/internal/repositories"
	"gorm.io/gorm"
)

type Services struct {
	Example *ExampleService
}

func New(db *gorm.DB) *Services {
	repos := repositories.New(db)
	
	return &Services{
		Example: NewExampleService(repos.Example),
	}
}`,

	"internal/services/example_service.go": `package services

import (
	"{{.ProjectName}}/internal/models"
	"{{.ProjectName}}/internal/repositories"
)

type ExampleService struct {
	exampleRepo *repositories.ExampleRepository
}

func NewExampleService(exampleRepo *repositories.ExampleRepository) *ExampleService {
	return &ExampleService{
		exampleRepo: exampleRepo,
	}
}

func (s *ExampleService) GetExample() map[string]interface{} {
	return map[string]interface{}{
		"message": "Hello from LupettoGo! 🐺",
		"status":  "success",
		"data": map[string]interface{}{
			"example": "This is an example response from the service layer",
			"tips":    "Replace this service with your business logic",
		},
	}
}

func (s *ExampleService) GetAllExamples() ([]*models.Example, error) {
	return s.exampleRepo.FindAll()
}

func (s *ExampleService) GetExampleByID(id uint) (*models.Example, error) {
	return s.exampleRepo.FindByID(id)
}`,

	"internal/handlers/handlers.go": `package handlers

import (
	"{{.ProjectName}}/internal/services"
)

type Handlers struct {
	Example *ExampleHandler
}

func New(services *services.Services) *Handlers {
	return &Handlers{
		Example: NewExampleHandler(services.Example),
	}
}`,

	"internal/handlers/example_handler.go": `package handlers

import (
	"net/http"

	"{{.ProjectName}}/internal/services"
	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	exampleService *services.ExampleService
}

func NewExampleHandler(exampleService *services.ExampleService) *ExampleHandler {
	return &ExampleHandler{
		exampleService: exampleService,
	}
}

func (h *ExampleHandler) GetExample(c *gin.Context) {
	data := h.exampleService.GetExample()
	c.JSON(http.StatusOK, data)
}`,
}