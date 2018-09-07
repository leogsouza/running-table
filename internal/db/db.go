package db

// Record simulates a single database record
type Record struct {
	Name string  `json:"name"`
	Time float32 `json:"time"`
}

// NewRecord returns a new Record object
func NewRecord(name string, time float32) Record {
	return Record{name, time}
}

// Database holds all records
type Database struct {
	contents []Record
}

// New return new Database
func New() Database {
	contents := make([]Record, 0)
	return Database{contents}
}

// AddRecord add a record into database contents
func (database *Database) AddRecord(r Record) {
	database.contents = append(database.contents, r)
}

// GetRecords returns all records from database
func (database *Database) GetRecords() []Record {
	return database.contents
}
