package models

//+ validacije - name, version su obavezni
type GroupConfiguration struct {
	Configuration Configuration     `json:"configuration"`
	Labels        map[string]string `json:"labels"`
}

type ConfigurationGroup struct {
	ID             string               `json:"id"`
	Name           string               `json:"name" binding:"required"`
	Version        string               `json:"version" binding:"required"`
	Configurations []GroupConfiguration `json:"configurations"`
}
