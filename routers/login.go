package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BreCkver/Go-Twitter/bd"
	"github.com/BreCkver/Go-Twitter/jwt"
	"github.com/BreCkver/Go-Twitter/models"
)

/*Login realiza el login */
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("contest-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "No se logro serializar el request "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario/Contraseña inválidos", 400)
		return
	}

	jwtkey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}

	respo := models.RespuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respo)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
