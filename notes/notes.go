package notes

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type NotesManager struct {
	notes map[string][]string
	mu    sync.Mutex
}

func NewNotesManager() *NotesManager {
	manager := &NotesManager{
		notes: make(map[string][]string),
	}
	manager.loadNotes()
	return manager
}

func (n *NotesManager) GetNotes(name string) ([]string, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	notes, exists := n.notes[name]
	if !exists {
		return nil, fmt.Errorf("no notes found for user: %s", name)
	}

	return notes, nil
}

func (n *NotesManager) AddNote(name, note string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.notes[name] = append(n.notes[name], note)
	n.saveNotes()
}

func (n *NotesManager) UpdateNote(name, note string, index int) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if index < 0 || index >= len(n.notes[name]) {
		return fmt.Errorf("index out of range")
	}
	n.notes[name][index] = note
	n.saveNotes()
	return nil
}

func (n *NotesManager) saveNotes() {
	file, err := os.Create("notes.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(n.notes)
	if err != nil {
		fmt.Println("Error encoding notes:", err)
	}
}

func (n *NotesManager) loadNotes() {
	file, err := os.Open("notes.json")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error opening file:", err)
		}
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&n.notes)
	if err != nil {
		fmt.Println("Error decoding notes:", err)
	}
}
