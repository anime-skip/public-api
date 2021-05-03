package handlers

import (
	"net/http"

	"anime-skip.com/backend/internal/services/health_service"
	"anime-skip.com/backend/internal/utils/http_utils"
)

// Status is the handler that can be easily reached to tell if the application is running
func Status(w http.ResponseWriter, r *http.Request) {
	config := health_service.GetServerConfig()
	http_utils.JSON(w, http.StatusOK, config)
}
