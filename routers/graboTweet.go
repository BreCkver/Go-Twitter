package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BreCkver/Go-Twitter/bd"
	"github.com/BreCkver/Go-Twitter/models"
)

/*GraboTweet  */
func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.GraboTweet
	err := json.NewDecoder(r.Body).Decode((&mensaje))

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el registro, intente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
