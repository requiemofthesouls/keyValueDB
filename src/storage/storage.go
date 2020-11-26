package storage

type KVStorage interface {
	NewStorage() KVStorage

	// String crud
	CreateString(key, value string) error
	ReadString(key string) (string, error)
	UpdateString(key, value string) error
	DeleteString(key string) error

	PrintDump() error
	String() string
}
