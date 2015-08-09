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

//Deloyment database keys
const (
	//    ------------- Key formats ----------------
	DeployBucketDbKey        = "deployments"
	DeployCounterKeyFmt      = "%s:counter"
	DeployObjectDbKeyFmt     = "%s:data:%d"
	DeployObjectListDbKeyFmt = "%s:data:"
	// This key will contain the deployment json body
	// examples:
	//	- {NAMESPACE}:data:{INCREMENTAL DEPLOY ID}
	//  - slok/daton:data:1
	//  - slok/daton:data:98
	//	- docker/docker:data:4
	DeployQueryDbKeyFmt = "%s:query:%s"
	// This key will contain a list with deploy keys
	// examples:
	//	- {NAMESPACE}:query:{REF/SHA/ENV}
	//
	//  - byEnv:	slok/daton:query:production
	//			  	slok/daton:query:staging
	//	- byTask:	slok/daton:query:deploy
	//				slok/daton:query:migrate
	//	- byRef:	slok/daton:query:master
	//				slok/daton:query:tagv1
	//				slok/daton:query:aa271b21ae983e8dc188a111699c368888a2fed7
)

var (
	// Global connection
	globalDb *BoltDb = nil
)

// Bolt database will be the storage for the data
type BoltDb struct {
	Conn *bolt.DB
	Path string
}

// Returns a boltdb connection (if doesn't exists it creates)
func GetBoltDb() (*BoltDb, error) {
	if globalDb == nil {
		path := viper.GetString("BoltdbName")
		bdb, err := newBoltDb(path)
		if err != nil {
			return nil, err
		}
		globalDb = bdb
		log.WithFields(log.Fields{"path": globalDb.Path}).Info("New bolt database connection")
	}

	return globalDb, nil
}

func newBoltDb(path string) (*BoltDb, error) {

	// Connect
	connection, err := bolt.Open(path, dbFileMode, nil)
	if err != nil {
		log.WithFields(log.Fields{"path": globalDb.Path}).Errorf("Error connecting to bolt: %v", err)
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
	if b == nil {
		// If no database dont panic
		log.Warning("Database is nil, no need to disconnect")
		return nil
	}
	if b.Conn != nil {
		log.WithFields(log.Fields{"path": globalDb.Path}).Info("Disconnect bolt database")
		err := b.Conn.Close()
		if err != nil {
			log.WithFields(log.Fields{"path": globalDb.Path}).Errorf("Error disconnecting to bolt: %v", err)
			return err
		}
		globalDb = nil
	}
	return nil
}
