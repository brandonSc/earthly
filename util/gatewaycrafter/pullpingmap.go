package gatewaycrafter

import (
	"fmt"
	"sync"

	"github.com/earthly/earthly/util/stringutil"
)

// PullPingMap is a thread-save map used for coordinating pullpings
type PullPingMap struct {
	m           sync.Mutex
	localImages map[string]string
}

// NewPullPingMap returns a new PullPingMap
func NewPullPingMap() *PullPingMap {
	return &PullPingMap{
		localImages: map[string]string{},
	}
}

// Get fetches an existing entry from the map or returns false if none exists
func (ppm *PullPingMap) Get(k string) (string, bool) {
	ppm.m.Lock()
	defer ppm.m.Unlock()
	v, ok := ppm.localImages[k]
	return v, ok
}

// Insert creates a new entry for the value under sessionID/<v'>-<uuid>
// Where v' is v without special chars
func (ppm *PullPingMap) Insert(sessionID, v string) string {
	k := fmt.Sprintf("sess-%s/pullping:%s-%s", sessionID, stringutil.AlphanumericOnly(v), stringutil.RandomAlphanumeric(32))
	ppm.m.Lock()
	defer ppm.m.Unlock()
	ppm.localImages[k] = v
	return k
}
