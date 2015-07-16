package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func TestUtils(t *testing.T) { TestingT(t) }

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

func (s *FileUtilsTestSuite) TestWriteStringFails(c *C) {
	data := []byte("This is a test")
	err := WriteStringFile(data, "/doesnt/exists")

	if err == nil {
		c.Error("Writting should fail because of unexistent path")
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

func (s *FileUtilsTestSuite) TestWriteJsonFails(c *C) {
	data := map[int]string{
		1: "Json format doens't admit map keys that are not an string",
	}
	err := WriteJsonFile(data, s.filePath)

	if err == nil {
		c.Error("Json parsing should fail, invalid json object")
	}
}

func (s *FileUtilsTestSuite) TestWriteJsonFile(c *C) {
	type JsonTest struct {
		Attr1 string   `json:"attr1,omitempty"`
		Attr2 int      `json:"attr2,omitempty"`
		Attr3 []string `json:"attr3,omitempty"`
	}

	jt := JsonTest{
		Attr1: "this is a test",
		Attr2: 923123123,
		Attr3: []string{"a", "b", "c", "d"},
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
