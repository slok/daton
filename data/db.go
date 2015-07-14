package data

import (
	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
)

const (
	// Permissions to use on the db file.
	dbFileMode = 0600
)

var (
	// Global connection
	db *BoltDb = nil
)

// Bolt database will be the storage for the data
type BoltDb struct {
	Conn *bolt.DB
	Path string
}

// Returns a boltdb connection (if doesn't exists it creates)
func GetBoltDb() (*BoltDb, error) {
	if db == nil {
		path := viper.GetString("BoltdbName")
		bdb, err := newBoltDb(path)
		if err != nil {
			return nil, err
		}
		db = bdb
	}

	log.WithFields(log.Fields{"path": db.Path}).Info("New bolt database connection")
	return db, nil
}

func newBoltDb(path string) (*BoltDb, error) {

	// Connect
	connection, err := bolt.Open(path, dbFileMode, nil)
	if err != nil {
		log.WithFields(log.Fields{"path": db.Path}).Errorf("Error connecting to bolt: %v", err)
		return nil, err
	}

	// Create the new store
	bdb := &BoltDb{
		Conn: connection,
		Path: path,
	}

	return bdb, nil
}

func (b *BoltDb) Disconnect() error {
	if b.Conn != nil {
		log.WithFields(log.Fields{"path": db.Path}).Info("Disconnect bolt database")
		err := b.Conn.Close()
		if err != nil {
			log.WithFields(log.Fields{"path": db.Path}).Errorf("Error disconnecting to bolt: %v", err)
			return err
		}
		db = nil
	}
	return nil
}
