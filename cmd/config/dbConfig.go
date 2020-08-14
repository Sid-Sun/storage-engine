package config

import "fmt"

type DBConfig struct {
	env        string
	user       string
	pass       string
	host       string
	port       string
	database   string
	collection string
	timeout    int
}

func (db DBConfig) Address() string {
	if db.env == dev {
		return fmt.Sprintf("%s://%s:%s@%s:%s/", "mongodb", db.user, db.pass, db.host, db.port)
	}
	return fmt.Sprintf("%s://%s:%s@%s/?retryWrites=true&w=majority", "mongodb+srv", db.user, db.pass, db.host)
}

func (db DBConfig) Collection() string {
	return db.collection
}

func (db DBConfig) Database() string {
	return db.database
}

func (db DBConfig) TimeoutInSec() int {
	return db.timeout
}
