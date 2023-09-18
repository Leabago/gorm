package main

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/notes", a.getNotes).Methods("GET")
	a.Router.HandleFunc("/notes", a.createNote).Methods("POST")
	a.Router.HandleFunc("/note/{id:[0-9]+}", a.getNote).Methods("GET")
	a.Router.HandleFunc("/note/{id:[0-9]+}", a.updateNoteFields).Methods("PATCH")
	a.Router.HandleFunc("/note/{id:[0-9]+}", a.updateFullNote).Methods("PUT")
	a.Router.HandleFunc("/note/{id:[0-9]+}", a.deleteNote).Methods("DELETE")
}
