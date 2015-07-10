package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func TestFileUtils(t *testing.T) { TestingT(t) }

type FileUtilsTestSuite struct {
	filePath string
}

var _ = Suite(&FileUtilsTestSuite{})

func (s *FileUtilsTestSuite) SetUpSuite(c *C) {
	s.filePath = "/tmp/daton-test/FileUtilsTestSuite.txt"
}

func (s *FileUtilsTestSuite) TearDownTest(c *C) {
	err := os.RemoveAll(s.filePath)
	if err != nil {
		panic(err)
	}
}

func (s *FileUtilsTestSuite) TestWriteStringFile(c *C) {
	data := []byte("This is a test")
	err := WriteStringFile(data, s.filePath)

	if err != nil {
		c.Error(err)
	}

	dataAfter, err := ioutil.ReadFile(s.filePath)

	if err != nil {
		c.Error(err)
	}

	c.Assert(string(dataAfter), Equals, string(data))
}

func (s *FileUtilsTestSuite) TestWriteJsonFile(c *C) {
	type JsonTest struct {
		attr1 string   `json:"attr1,omitempty"`
		attr2 int      `json:"attr2,omitempty"`
		attr3 []string `json:"attr3,omitempty"`
	}

	jt := JsonTest{
		attr1: "this is a test",
		attr2: 923123123,
		attr3: []string{"a", "b", "c", "d"},
	}

	err := WriteJsonFile(jt, s.filePath)

	if err != nil {
		c.Error(err)
	}

	dataAfter, err := ioutil.ReadFile(s.filePath)

	if err != nil {
		c.Error(err)
	}

	data, _ := json.Marshal(jt)
	c.Assert(string(data), Equals, string(dataAfter))
}
