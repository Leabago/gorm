package main

import (
	"encoding/json"
	"go-orm/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (a *App) getNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	n := model.Note{}
	n.ID = uint(id)

	if err := n.GetNote(a.DB).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Note not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, n)
}

func (a *App) getNotes(w http.ResponseWriter, r *http.Request) {
	limitName := "limit"
	offsetName := "offset"
	limit := r.URL.Query().Get(limitName)
	offset := r.URL.Query().Get(offsetName)
	startInt, err := parseIntParameter(limit, limitName, w)
	if err != nil {
		return
	}
	countInt, err := parseIntParameter(offset, offsetName, w)
	if err != nil {
		return
	}

	note := model.Note{}
	var notes []model.Note

	if err := note.GetNotes(a.DB, &notes, startInt, countInt).Error; err != nil {
		switch err {
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, notes)
}

func (a *App) createNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := note.CreateNote(a.DB).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, note)
}

func (a *App) updateNoteFields(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	var note model.Note
	note.ID = uint(id)

	if err := note.GetNote(a.DB).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Note not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	// set id if rewrited
	note.ID = uint(id)
	if err := note.UpdateNoteFields(a.DB).Error; err != nil {
		switch err {
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, note)
}

func (a *App) updateFullNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var note model.Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	note.ID = uint(id)

	if err := note.UpdateFullNote(a.DB).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "Note not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, note)
}

func (a *App) deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	note := model.Note{}
	note.ID = uint(id)

	if err := note.DeleteNote(a.DB).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
