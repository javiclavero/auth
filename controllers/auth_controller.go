// controllers/auth_controller.go
package controllers

import (
	"encoding/json"
	"github.com/javiclavero/go-auth-service/config"
	"github.com/javiclavero/go-auth-service/services"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error en los datos proporcionados", http.StatusBadRequest)
		return
	}

	// Obtener la conexión a la base de datos
	client := config.ConnectDB()
	defer client.Close()

	// Llamar al servicio de autenticación
	success, err := services.Login(client, req.Username, req.Password)
	if err != nil {
		json.NewEncoder(w).Encode(LoginResponse{Message: err.Error(), Success: false})
		return
	}

	json.NewEncoder(w).Encode(LoginResponse{Message: "Login exitoso", Success: success})
}
