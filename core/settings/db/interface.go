package db

type IDBSettings interface {
	GetDsn() string
	InitFromEnv() error
}
