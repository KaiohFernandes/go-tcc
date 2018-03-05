package database

type Database struct {
	source string
	url string
	user string
	password string
	port int
}

func (db *Database) config(){
	db.source = "firebase"
	db.url = "path/to/serviceAccountKey.json"
}