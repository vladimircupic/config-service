package handlers

import (
	"config-service/models"
	"config-service/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// IMUTABILNOST - ne menjamo postojece, kreiramo nove sa izmenjenim vrednostima
func CreateGroup(c *gin.Context) {

	var group models.ConfigurationGroup

	if err := c.ShouldBindJSON(&group); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	for _, existingGroup := range storage.Groups {

		if existingGroup.Name == group.Name &&
			existingGroup.Version == group.Version {

			c.JSON(http.StatusConflict, gin.H{
				"error": "Group version already exists",
			})

			return
		}
	}

	group.ID = uuid.New().String()

	storage.Groups[group.ID] = group

	c.JSON(http.StatusCreated, group)
}

func GetGroup(c *gin.Context) {

	name := c.Param("name")
	version := c.Param("version")

	for _, group := range storage.Groups {

		if group.Name == name && group.Version == version {

			c.JSON(http.StatusOK, group)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Group not found",
	})
}

func DeleteGroup(c *gin.Context) {

	name := c.Param("name")
	version := c.Param("version")

	for id, group := range storage.Groups {

		if group.Name == name && group.Version == version {

			delete(storage.Groups, id)

			c.JSON(http.StatusOK, gin.H{
				"message": "Group deleted",
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Group not found",
	})
}

// dodavanje u grupu
func AddConfigurationToGroup(c *gin.Context) {

	groupName := c.Param("name")
	groupVersion := c.Param("version")

	var newConfig models.GroupConfiguration

	if err := c.ShouldBindJSON(&newConfig); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	for id, group := range storage.Groups {

		if group.Name == groupName &&
			group.Version == groupVersion {

			for _, existing := range group.Configurations {

				if existing.Configuration.Name == newConfig.Configuration.Name &&
					existing.Configuration.Version == newConfig.Configuration.Version {

					c.JSON(http.StatusConflict, gin.H{
						"error": "Configuration already exists in group",
					})

					return
				}
			}

			group.Configurations = append(group.Configurations, newConfig)

			storage.Groups[id] = group

			c.JSON(http.StatusOK, group)

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Group not found",
	})
}

// uklanjanje iz grupe
func RemoveConfigurationFromGroup(c *gin.Context) {

	groupName := c.Param("name")
	groupVersion := c.Param("version")

	configName := c.Param("configName")
	configVersion := c.Param("configVersion")

	for id, group := range storage.Groups {

		if group.Name == groupName &&
			group.Version == groupVersion {

			var updatedConfigurations []models.GroupConfiguration

			for _, config := range group.Configurations {

				if !(config.Configuration.Name == configName &&
					config.Configuration.Version == configVersion) {

					updatedConfigurations = append(updatedConfigurations, config)
				}
			}

			group.Configurations = updatedConfigurations

			storage.Groups[id] = group

			c.JSON(http.StatusOK, group)

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Group not found",
	})
}

// Labeli pretraga
func FilterConfigurationsByLabels(c *gin.Context) {

	groupName := c.Param("name")
	groupVersion := c.Param("version")

	labelKey := c.Query("key")
	labelValue := c.Query("value")

	for _, group := range storage.Groups {

		if group.Name == groupName &&
			group.Version == groupVersion {

			var filtered []models.GroupConfiguration

			for _, config := range group.Configurations {

				if value, exists := config.Labels[labelKey]; exists && value == labelValue {

					filtered = append(filtered, config)
				}
			}

			c.JSON(http.StatusOK, filtered)

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Group not found",
	})
}
