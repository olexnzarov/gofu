package sqlitestorage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/olexnzarov/gofu/internal/gofudaemon/procmanager"
	"github.com/olexnzarov/gofu/internal/logger"
)

var (
	ErrNoProcess = errors.New("storage: process not found")
)

type ProcessStorage struct {
	log     logger.Logger
	db      *sql.DB
	dbMutex *sync.RWMutex
}

// New returns an uninitialized process storage. You must call Initialize before using it.
func New(log logger.Logger, db *sql.DB) *ProcessStorage {
	storage := &ProcessStorage{
		log:     log,
		db:      db,
		dbMutex: &sync.RWMutex{},
	}

	return storage
}

// Initialize makes sure the storage is ready to use.
func (s *ProcessStorage) Initialize() error {
	s.dbMutex.Lock()
	defer s.dbMutex.Unlock()

	s.log.Info("Initializing the process storage...")

	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS processes(id TEXT, configuration BLOB)")

	return err
}

func fmtError(label string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("storage: %s: %w", label, err)
}

func (s *ProcessStorage) List() ([]*procmanager.ProcessData, error) {
	s.dbMutex.RLock()
	defer s.dbMutex.RUnlock()

	rows, err := s.db.Query("SELECT id, configuration FROM processes")
	if err != nil {
		return nil, fmtError("failed to list", err)
	}

	processes := []*procmanager.ProcessData{}
	for rows.Next() {
		var blob []byte
		data := procmanager.ProcessData{}

		if err := rows.Scan(&data.Id, &blob); err != nil {
			s.log.Errorf("Failed to read from the database, process=%s: %s", data.Id, err)
			continue
		}

		if err := json.Unmarshal(blob, &data.Configuration); err != nil {
			s.log.Errorf("Found a broken configuration, process=%s: %s", data.Id, err)
			continue
		}

		processes = append(processes, &data)
	}

	return processes, nil
}

func (s *ProcessStorage) Upsert(process *procmanager.ProcessData) error {
	s.dbMutex.Lock()
	defer s.dbMutex.Unlock()

	// Check if process record exists.
	err := s.db.QueryRow("SELECT id FROM processes WHERE id = ?", process.Id).Scan(&process.Id)

	// If there's no error, process is already saved.
	if err == nil {
		return fmtError("failed to update", s.update(process))
	}

	// If the error is ErrNoRows, we save the process.
	if err == sql.ErrNoRows {
		return fmtError("failed to save", s.save(process))
	}

	return fmtError("failed to upsert", err)
}

func (s *ProcessStorage) Delete(id string) error {
	s.dbMutex.Lock()
	defer s.dbMutex.Unlock()

	r, err := s.db.Exec("DELETE FROM processes WHERE id = ?", id)
	if err != nil {
		return fmtError("failed to delete", err)
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return fmtError("RowsAffected is not supported", err)
	}
	if rowsAffected == 0 {
		return ErrNoProcess
	}

	return nil
}

func (s *ProcessStorage) save(process *procmanager.ProcessData) error {
	bytes, err := json.Marshal(process.Configuration)
	if err != nil {
		return nil
	}

	_, err = s.db.Exec(
		"INSERT INTO processes(id, configuration) VALUES(?, ?)",
		process.Id,
		bytes,
	)

	return err
}

func (s *ProcessStorage) update(process *procmanager.ProcessData) error {
	bytes, err := json.Marshal(process.Configuration)
	if err != nil {
		return nil
	}

	_, err = s.db.Exec(
		"UPDATE processes SET configuration = ? WHERE id = ?",
		bytes,
		process.Id,
	)

	return err
}
