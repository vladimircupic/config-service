package models

// + validacije - name, version, params su obavezni
type Configuration struct {
	ID      string            `json:"id"`
	Name    string            `json:"name" binding:"required"`
	Version string            `json:"version" binding:"required"`
	Params  map[string]string `json:"params" binding:"required"`
}
