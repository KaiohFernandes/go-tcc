package database

type Database struct {
	source string
	url string
	user string
	password string
	port int
}

func (db *Database) Initialize() string{

	db.Config()

	if db.source == "firebase" {
		fb := Firebase{}
		fb.Connect(db)
		return ""//fb.Firebase{}.Connect(db)
	}

	return ""

}

func (db *Database) Config(){
	db.source = "firebase"
	db.url = "path/to/serviceAccountKey.json"
}