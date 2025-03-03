package controllers

import (
	"PA/models"
	"PA/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type ProjectInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func GetProjectsController(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("user_id").(uint)

	projects, err := services.GetAllProjectsService(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

func GetProjectByIDController(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("user_id").(uint)
	projectIDStr := c.Param("project_id")

	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := services.GetProjectByIDService(db, uint(projectID), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

func AddProjectController(c *gin.Context) {
	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("user_id").(uint)

	project := models.Project{
		Name:        input.Name,
		Description: input.Description,
		OwnerID:     userID,
	}

	if err := services.CreateProjectService(db, &project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat project"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": project})
}

func EditProjectController(c *gin.Context) {
	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("user_id").(uint)
	projectIDStr := c.Param("project_id")

	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := services.GetProjectByIDService(db, uint(projectID), userID)
	if err != nil {
        errorMsg := err.Error()
        statusCode := http.StatusNotFound
        if errorMsg == "anda tidak memiliki akses ke project ini" {
            statusCode = http.StatusForbidden
        }
        c.JSON(statusCode, gin.H{"error": errorMsg})
        return
    }

	project.Name = input.Name
	project.Description = input.Description

	if err := services.UpdateProjectService(db, &project, userID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

func DeleteProjectController(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("user_id").(uint)

	projectID, err := strconv.ParseUint(c.Param("project_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	if err := services.DeleteProjectService(db, uint(projectID), userID); err != nil {
		errorMsg := gin.H{"error": err.Error()}
		status := http.StatusInternalServerError

		if strings.Contains(err.Error(), "unauthorized") {
			status = http.StatusForbidden
		} else if strings.Contains(err.Error(), "tidak ditemukan") {
			status = http.StatusNotFound
		}

		c.JSON(status, errorMsg)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project berhasil dihapus"})
}

func AddCollaboratorController(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	ownerID := c.MustGet("user_id").(uint)

	projectIDStr := c.Param("project_id")
	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var input struct {
		UserID uint `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddCollaboratorService(db, uint(projectID), input.UserID, ownerID); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collaborator berhasil ditambahkan"})
}

func RemoveCollaboratorController(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	ownerID := c.MustGet("user_id").(uint)

	projectIDStr := c.Param("project_id")
	projectID, err := strconv.ParseUint(projectIDStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var input struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.RemoveCollaboratorService(db, uint(projectID), input.UserID, ownerID); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Collaborator berhasil dihapus"})
}
