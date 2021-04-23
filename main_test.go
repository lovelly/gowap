package gowap

import (
	"encoding/json"
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGowap(t *testing.T) {
	url := "https://www.boce.com"
	wapp, err := Init()
	if err != nil {
		log.Errorln(err)
		t.FailNow()
	}
	res, err := wapp.Analyze(url)
	if err != nil {
		log.Errorln(err)
		t.FailNow()
	}
	prettyJSON, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Errorln(err)
		t.FailNow()
	}
	fmt.Println( string(prettyJSON))
}
