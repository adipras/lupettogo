package generator

// Module generation templates

var moduleTemplates = map[string]string{
	"model.go.tmpl": `package models

import (
	"time"

	"gorm.io/gorm"
)

type __Module__ struct {
	ID        uint           ` + "`" + `json:"id" gorm:"primarykey"` + "`" + `
	Name      string         ` + "`" + `json:"name" gorm:"not null" validate:"required"` + "`" + `
	Status    string         ` + "`" + `json:"status" gorm:"default:active"` + "`" + `
	CreatedAt time.Time      ` + "`" + `json:"created_at"` + "`" + `
	UpdatedAt time.Time      ` + "`" + `json:"updated_at"` + "`" + `
	DeletedAt gorm.DeletedAt ` + "`" + `json:"-" gorm:"index"` + "`" + `
}

func (__Module__) TableName() string {
	return "__module__s"
}`,

	"repository.go.tmpl": `package repositories

import (
	"{{.ProjectName}}/internal/models"
	"gorm.io/gorm"
)

type __Module__Repository struct {
	db *gorm.DB
}

func New__Module__Repository(db *gorm.DB) *__Module__Repository {
	return &__Module__Repository{
		db: db,
	}
}

func (r *__Module__Repository) FindAll() ([]*models.__Module__, error) {
	var __module__s []*models.__Module__
	err := r.db.Find(&__module__s).Error
	return __module__s, err
}

func (r *__Module__Repository) FindByID(id uint) (*models.__Module__, error) {
	var __module__ models.__Module__
	err := r.db.First(&__module__, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &__module__, nil
}

func (r *__Module__Repository) Create(__module__ *models.__Module__) (*models.__Module__, error) {
	err := r.db.Create(__module__).Error
	return __module__, err
}

func (r *__Module__Repository) Update(__module__ *models.__Module__) (*models.__Module__, error) {
	err := r.db.Save(__module__).Error
	return __module__, err
}

func (r *__Module__Repository) Delete(id uint) error {
	return r.db.Delete(&models.__Module__{}, id).Error
}

func (r *__Module__Repository) FindByField(field string, value interface{}) ([]*models.__Module__, error) {
	var __module__s []*models.__Module__
	err := r.db.Where(field+" = ?", value).Find(&__module__s).Error
	return __module__s, err
}`,

	"service.go.tmpl": `package services

import (
	"errors"

	"{{.ProjectName}}/internal/models"
	"{{.ProjectName}}/internal/repositories"
)

type __Module__Service struct {
	__module__Repo *repositories.__Module__Repository
}

func New__Module__Service(__module__Repo *repositories.__Module__Repository) *__Module__Service {
	return &__Module__Service{
		__module__Repo: __module__Repo,
	}
}

func (s *__Module__Service) GetAll__Module__s() ([]*models.__Module__, error) {
	return s.__module__Repo.FindAll()
}

func (s *__Module__Service) Get__Module__ByID(id uint) (*models.__Module__, error) {
	return s.__module__Repo.FindByID(id)
}

func (s *__Module__Service) Create__Module__(__module__ *models.__Module__) (*models.__Module__, error) {
	// Add business logic validation here
	if err := s.validate__Module__(__module__); err != nil {
		return nil, err
	}

	return s.__module__Repo.Create(__module__)
}

func (s *__Module__Service) Update__Module__(__module__ *models.__Module__) (*models.__Module__, error) {
	// Check if __module__ exists
	existing, err := s.__module__Repo.FindByID(__module__.ID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("__module__ not found")
	}

	// Add business logic validation here
	if err := s.validate__Module__(__module__); err != nil {
		return nil, err
	}

	return s.__module__Repo.Update(__module__)
}

func (s *__Module__Service) Delete__Module__(id uint) error {
	// Check if __module__ exists
	existing, err := s.__module__Repo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("__module__ not found")
	}

	return s.__module__Repo.Delete(id)
}

func (s *__Module__Service) validate__Module__(__module__ *models.__Module__) error {
	// Add your business logic validation here
	// Example:
	// if __module__.Name == "" {
	//     return errors.New("name is required")
	// }
	return nil
}`,

	"handler.go.tmpl": `package handlers

import (
	"net/http"
	"strconv"

	"{{.ProjectName}}/internal/models"
	"{{.ProjectName}}/internal/services"
	"github.com/gin-gonic/gin"
)

type __Module__Handler struct {
	__module__Service *services.__Module__Service
}

func New__Module__Handler(__module__Service *services.__Module__Service) *__Module__Handler {
	return &__Module__Handler{
		__module__Service: __module__Service,
	}
}

// Get__Module__s godoc
// @Summary Get all __module__s
// @Description Get a list of all __module__s
// @Tags __module__s
// @Accept json
// @Produce json
// @Success 200 {array} models.__Module__
// @Router /__module__s [get]
func (h *__Module__Handler) Get__Module__s(c *gin.Context) {
	__module__s, err := h.__module__Service.GetAll__Module__s()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, __module__s)
}

// Get__Module__ godoc
// @Summary Get a __module__ by ID
// @Description Get a single __module__ by its ID
// @Tags __module__s
// @Accept json
// @Produce json
// @Param id path int true "__Module__ ID"
// @Success 200 {object} models.__Module__
// @Failure 404 {object} map[string]string
// @Router /__module__s/{id} [get]
func (h *__Module__Handler) Get__Module__(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	__module__, err := h.__module__Service.Get__Module__ByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "__Module__ not found"})
		return
	}

	c.JSON(http.StatusOK, __module__)
}

// Create__Module__ godoc
// @Summary Create a new __module__
// @Description Create a new __module__ with the given data
// @Tags __module__s
// @Accept json
// @Produce json
// @Param __module__ body models.__Module__ true "__Module__ object"
// @Success 201 {object} models.__Module__
// @Failure 400 {object} map[string]string
// @Router /__module__s [post]
func (h *__Module__Handler) Create__Module__(c *gin.Context) {
	var __module__ models.__Module__
	if err := c.ShouldBindJSON(&__module__); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created__Module__, err := h.__module__Service.Create__Module__(&__module__)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created__Module__)
}

// Update__Module__ godoc
// @Summary Update a __module__
// @Description Update a __module__ with the given data
// @Tags __module__s
// @Accept json
// @Produce json
// @Param id path int true "__Module__ ID"
// @Param __module__ body models.__Module__ true "__Module__ object"
// @Success 200 {object} models.__Module__
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /__module__s/{id} [put]
func (h *__Module__Handler) Update__Module__(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var __module__ models.__Module__
	if err := c.ShouldBindJSON(&__module__); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	__module__.ID = uint(id)
	updated__Module__, err := h.__module__Service.Update__Module__(&__module__)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated__Module__)
}

// Delete__Module__ godoc
// @Summary Delete a __module__
// @Description Delete a __module__ by ID
// @Tags __module__s
// @Accept json
// @Produce json
// @Param id path int true "__Module__ ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /__module__s/{id} [delete]
func (h *__Module__Handler) Delete__Module__(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.__module__Service.Delete__Module__(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "__Module__ not found"})
		return
	}

	c.Status(http.StatusNoContent)
}`,
}