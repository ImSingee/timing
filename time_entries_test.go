package timing

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var apiKey string

func InitTest() {
	if apiKey == "" {
		f, err := os.OpenFile("token.test", os.O_RDONLY, 0)
		if err != nil {
			panic("Open token.test file failed")
		}
		defer f.Close()

		k, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Read token.test file failed")
		}

		apiKey = string(k)
	}
	log.Println("apiKey:", apiKey)

	Init(apiKey)
}

func TestInit(t *testing.T) {
	InitTest()

	if !IsInitialed() {
		t.Fatal("Init fail")
	}
}

func TestTimeEntriesIsRunning(t *testing.T) {
	InitTest()

	response, err := TimeEntries(&TimeEntriesRequest{
		IsRunning: IsRunningTrue,
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", response)
}

func TestTimeEntriesInProject(t *testing.T) {
	InitTest()

	response, err := TimeEntries(&TimeEntriesRequest{
		Projects: []string{"/projects/1"},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", response)
}
