package processmanagerstorage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/olexnzarov/gofu/internal/daemon/logger"
	"github.com/olexnzarov/gofu/internal/daemon/processmanager"
	"github.com/olexnzarov/gofu/internal/daemon/processmanager/service/managedprocess"
)

type SQLDB struct {
	*sql.DB
}

type ProcessStorage struct {
	log     logger.Logger
	db      *SQLDB
	dbMutex *sync.RWMutex
}

// New returns an uninitialized process storage. You must call Initialize before using it.
func NewSQL(log logger.Logger, db *SQLDB) processmanager.Storage {
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

func (s *ProcessStorage) List() ([]processmanager.ProcessData, error) {
	s.dbMutex.RLock()
	defer s.dbMutex.RUnlock()

	rows, err := s.db.Query("SELECT id, configuration FROM processes")
	if err != nil {
		return nil, fmtError("failed to list", err)
	}

	processes := []processmanager.ProcessData{}
	for rows.Next() {
		var blob []byte
		data := managedprocess.ProcessData{}

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

func (s *ProcessStorage) Upsert(process processmanager.ProcessData) error {
	s.dbMutex.Lock()
	defer s.dbMutex.Unlock()

	// Check if process record exists.
	t := ""
	err := s.db.QueryRow("SELECT id FROM processes WHERE id = ?", process.GetID()).Scan(&t)

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
		return processmanager.ErrStorageNoProcess
	}

	return nil
}

func (s *ProcessStorage) save(process processmanager.ProcessData) error {
	bytes, err := json.Marshal(process.GetConfiguration())
	if err != nil {
		return nil
	}

	_, err = s.db.Exec(
		"INSERT INTO processes(id, configuration) VALUES(?, ?)",
		process.GetID(),
		bytes,
	)

	return err
}

func (s *ProcessStorage) update(process processmanager.ProcessData) error {
	bytes, err := json.Marshal(process.GetConfiguration())
	if err != nil {
		return nil
	}

	_, err = s.db.Exec(
		"UPDATE processes SET configuration = ? WHERE id = ?",
		bytes,
		process.GetID(),
	)

	return err
}
