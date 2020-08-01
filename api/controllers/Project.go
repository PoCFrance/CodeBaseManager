package controllers

import (
	"cbm-api/database"
	"cbm-api/models"
)

func FindProject(db database.Database, name string) *models.Project {
	project := models.Project{
		Name: name,
	}
	result := db.DB.First(&project)
	if result.Error != nil {
		return nil
	}
	return &project
}

func ListProjects(db database.Database) []models.Project {
	var projects []models.Project
	result := db.DB.Find(&projects)
	if result.Error != nil {
		return nil
	}
	return projects
}
