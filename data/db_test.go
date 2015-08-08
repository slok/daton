package data

import (
	. "gopkg.in/check.v1"

	"github.com/spf13/viper"

	"github.com/slok/daton/configuration"
)

type BoltDbSuite struct {
	db *BoltDb
}

var _ = Suite(&BoltDbSuite{})

func (s *BoltDbSuite) SetUpSuite(c *C) {
	configuration.LoadSettingsFromFile()
	viper.Set("BoltdbName", "/tmp/datontest.db")
}

func (s *BoltDbSuite) TearDownTest(c *C) {
	s.db.Disconnect()
}

func (s *BoltDbSuite) TestDbConnect(c *C) {
	db, err := GetBoltDb()

	if err != nil {
		c.Errorf("Error creating the database: %v", err)
	}
	if err := db.Conn.Sync(); err != nil {
		c.Errorf("Error creating the database: %v", err)
	}
	db.Conn.Info()
}

func (s *BoltDbSuite) TestDbSingleton(c *C) {
	db, _ := GetBoltDb()
	db2, _ := GetBoltDb()

	c.Assert(db, Equals, db2)
}

func (s *BoltDbSuite) TestDbDisconnect(c *C) {
	defer func() {
		// This function should panic so we need an error
		if r := recover(); r == nil {
			c.Error("Error on boltdb disconnection")
		}
	}()
	db, _ := GetBoltDb()
	db.Disconnect()
	db.Conn.Info()
}
