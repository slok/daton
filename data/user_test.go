package data

import (
	"encoding/json"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestDataUser(t *testing.T) { TestingT(t) }

type DataUserSuite struct {
}

var _ = Suite(&DataUserSuite{})

func (s *DataUserSuite) TestUserModel(c *C) {
	u := User{
		Login: "slok",
		Id:    1,
	}

	jsonOk := `{"login":"slok","id":1}`
	j, err := json.Marshal(u)

	if err != nil {
		c.Error(err)
	}

	c.Assert(string(j), Equals, jsonOk)
}
