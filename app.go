package main

import (
	"context"
	"v-notes/notes"
)

// App struct
type App struct {
	ctx         context.Context
	notesManager *notes.NotesManager
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		notesManager: notes.NewNotesManager(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetNotes(name string) ([]string, error) {
	return a.notesManager.GetNotes(name)
}

func (a *App) AddNote(name string, note string) {
	a.notesManager.AddNote(name, note)
}

func (a *App) UpdateNote(name, note string, index int) error {
	return a.notesManager.UpdateNote(name, note, index)
}
