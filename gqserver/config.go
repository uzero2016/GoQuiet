package gqserver

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

// Provides either the real current time or a fake time
type TimeFn func() time.Time

// Global states
type State struct {
	WebServerAddr  string
	Key            string
	TicketTimeHint int
	AESKey         []byte
	Now            TimeFn
	SS_LOCAL_HOST  string
	SS_LOCAL_PORT  string
	SS_REMOTE_HOST string
	SS_REMOTE_PORT string
	UsedRandom     map[[32]byte]int
}

// ParseConfig parses the config file into a State variable
func ParseConfig(configPath string, sta *State) error {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.New("Failed to read config file. File may not exist")
	}
	err = json.Unmarshal(content, &sta)
	if err != nil {
		return errors.New("Bad config json format")
	}
	MakeAESKey(sta)
	return nil
}

// MakeAESKey calculates the SHA256 of the string key and writes it to the AESKey field
func MakeAESKey(sta *State) {
	h := sha256.New()
	h.Write([]byte(sta.Key))
	sta.AESKey = h.Sum(nil)
}
