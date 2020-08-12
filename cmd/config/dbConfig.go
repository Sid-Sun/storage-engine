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
}

func (db DBConfig) Address() string {
	if db.env == dev {
		return fmt.Sprintf("%s://%s:%s@%s:%s/", "mongodb", db.user, db.pass, db.host, db.port)
	}
	return fmt.Sprintf("%s://%s:%s@%s:%s/?ssl=true", "mongodb", db.user, db.pass, db.host, db.port)
}

func (db DBConfig) Collection() string {
	return db.collection
}

func (db DBConfig) Database() string {
	return db.database
}
