package sql

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"grapi/utils"
)

// Delete : delete entry from table
func (db *SQL) Delete(w http.ResponseWriter, r *http.Request) {
	tabName := mux.Vars(r)["table"]
	id := mux.Vars(r)["id"]
	idNum, errAtoi := strconv.Atoi(id)

	statement := fmt.Sprintf("DELETE FROM %s WHERE id=%d", tabName, idNum)
	if errAtoi != nil {
		fmt.Fprintf(w, "{\"Error\": \"Invalid id '%s'\"}", id)
	} else {
		_, err := db.DB.Query(statement)
		if err != nil {
			utils.SendError(w, err, http.StatusBadRequest)
		}
	}
}
