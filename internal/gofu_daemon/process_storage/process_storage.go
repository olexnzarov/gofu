package process_storage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"sync"

	"github.com/olexnzarov/gofu/internal/gofu_daemon/process_manager"
	"go.uber.org/zap"
)

var (
	ErrNoProcess = errors.New("process is not persistent")
)

type ProcessStorage struct {
	log     *zap.Logger
	db      *sql.DB
	dbMutex *sync.RWMutex
}

// New returns an uninitialized process storage. You must call Initialize before using it.
func New(log *zap.Logger, db *sql.DB) *ProcessStorage {
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

func (s *ProcessStorage) List() ([]*process_manager.ProcessData, error) {
	s.dbMutex.RLock()
	defer s.dbMutex.RUnlock()

	rows, err := s.db.Query("SELECT id, configuration FROM processes")
	if err != nil {
		return nil, err
	}

	processes := []*process_manager.ProcessData{}
	for rows.Next() {
		var blob []byte
		data := process_manager.ProcessData{}

		if err := rows.Scan(&data.Id, &blob); err != nil {
			s.log.Sugar().Errorf("Failed to read from the database, process=%s: %s", data.Id, err)
			continue
		}

		if err := json.Unmarshal(blob, &data.Configuration); err != nil {
			s.log.Sugar().Errorf("Found a broken configuration, process=%s: %s", data.Id, err)
			continue
		}

		processes = append(processes, &data)
	}

	return processes, nil
}

func (s *ProcessStorage) Upsert(process *process_manager.ProcessData) error {
	s.dbMutex.Lock()
	defer s.dbMutex.Unlock()

	// Check if process record exists.
	err := s.db.QueryRow("SELECT 1 FROM processes WHERE id = ?", process.Id).Scan()

	// If there's no error, process is already saved.
	if err == nil {
		return s.update(process)
	}

	// If the error is ErrNoRows, we save the process.
	if err == sql.ErrNoRows {
		return s.save(process)
	}

	return err
}

func (s *ProcessStorage) Delete(id string) error {
	s.dbMutex.Lock()
	defer s.dbMutex.Unlock()

	r, err := s.db.Exec("DELETE FROM processes WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoProcess
	}

	return nil
}

func (s *ProcessStorage) save(process *process_manager.ProcessData) error {
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

func (s *ProcessStorage) update(process *process_manager.ProcessData) error {
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
