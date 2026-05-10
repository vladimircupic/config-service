package storage

import "config-service/models"

var Configurations = make(map[string]models.Configuration)

var Groups = make(map[string]models.ConfigurationGroup)
