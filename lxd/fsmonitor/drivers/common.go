package drivers

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/canonical/lxd/lxd/fsmonitor"
	"github.com/canonical/lxd/shared/logger"
)

type common struct {
	logger     logger.Logger
	mu         sync.Mutex
	watches    map[string]map[string]func(string, fsmonitor.Event) bool
	prefixPath string
	events     []fsmonitor.Event
}

func (d *common) init(logger logger.Logger, path string, events []fsmonitor.Event) {
	d.logger = logger
	d.watches = make(map[string]map[string]func(string, fsmonitor.Event) bool)
	d.prefixPath = path
	d.events = events
}

// PrefixPath returns the prefix path.
func (d *common) PrefixPath() string {
	return d.prefixPath
}

// Watch creates a watch for a path which may or may not yet exist. If the provided path gets an
// inotify event, f() is called. If there already is a watch on the provided path, the callback
// function will simply be replaced without returning an error.
// Note: If f() returns false, the watch is removed.
func (d *common) Watch(path string, identifier string, f func(path string, event fsmonitor.Event) bool) error {
	d.logger.Info("Watching path", logger.Ctx{"path": path})

	if f == nil {
		return ErrInvalidFunction
	}

	path = filepath.Clean(path)

	if !strings.HasPrefix(path, d.prefixPath) {
		return &ErrInvalidPath{PrefixPath: d.prefixPath}
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	_, ok := d.watches[path]
	if !ok {
		d.watches[path] = make(map[string]func(string, fsmonitor.Event) bool)
	}

	_, ok = d.watches[path][identifier]
	if ok {
		return ErrWatchExists
	}

	d.watches[path][identifier] = f

	return nil
}

// Unwatch removes a watch.
func (d *common) Unwatch(path string, identifier string) error {
	d.logger.Info("Unwatching path", logger.Ctx{"path": path})

	d.mu.Lock()
	defer d.mu.Unlock()

	path = filepath.Clean(path)

	_, ok := d.watches[path]
	if !ok {
		return nil
	}

	delete(d.watches[path], identifier)

	if len(d.watches[path]) == 0 {
		delete(d.watches, path)
	}

	return nil
}
