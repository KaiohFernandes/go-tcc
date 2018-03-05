package database

type Database interface {
	// OpenConnection()
	// CloseConnection()
	// Get()
	// GetById(id int)
	// Insert()
	// Update()
	// Delete()
}

type DatabaseConfig struct {
	source string
	url string
	user string
	password string
	port int
}

func (db *DatabaseConfig) config(){
	db.source = "firebase"
	db.url = "path/to/serviceAccountKey.json"
}