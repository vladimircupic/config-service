package handlers

import (
	"config-service/models"
	"config-service/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// IMUTABILNOST - ne menjamo postojece, kreiramo nove sa izmenjenim vrednostima
func CreateConfiguration(c *gin.Context) {

	var config models.Configuration

	if err := c.ShouldBindJSON(&config); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	for _, existingConfig := range storage.Configurations {

		if existingConfig.Name == config.Name &&
			existingConfig.Version == config.Version {

			c.JSON(http.StatusConflict, gin.H{
				"error": "Configuration version already exists",
			})

			return
		}
	}

	config.ID = uuid.New().String()

	storage.Configurations[config.ID] = config

	c.JSON(http.StatusCreated, config)
}

func GetConfiguration(c *gin.Context) {

	name := c.Param("name")
	version := c.Param("version")

	for _, config := range storage.Configurations {

		if config.Name == name && config.Version == version {

			c.JSON(http.StatusOK, config)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Configuration not found",
	})
}

func DeleteConfiguration(c *gin.Context) {

	name := c.Param("name")
	version := c.Param("version")

	for id, config := range storage.Configurations {

		if config.Name == name && config.Version == version {

			delete(storage.Configurations, id)

			c.JSON(http.StatusOK, gin.H{
				"message": "Configuration deleted",
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Configuration not found",
	})
}
