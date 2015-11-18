package models
import (
	"testing"
	"encoding/json"
)


/*/ search test
func Test_Search(t *testing.T) {

	r, err := Search("毛")
	if err != nil {
		t.Error(err)
	}

	dat, _ := json.MarshalIndent(r, "", "  ")
	t.Log(string(dat))
}
*/

// view test
func Test_View(t *testing.T) {

	torrent, err := View(1)
	if err != nil {
		t.Error(err)
	}

	dat, _ := json.MarshalIndent(torrent, "", "  ")
	t.Log(string(dat))
}