package process_manager

type PersistentStorage interface {
	List() ([]*ProcessData, error)
	Upsert(process *ProcessData) error
	Delete(id string) error
}
