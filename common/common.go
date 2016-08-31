package common

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	currentPath string
	locker      *sync.Mutex = new(sync.Mutex)
)

func GetCurrentPath() string {
	if currentPath == "" {
		locker.Lock()
		defer locker.Unlock()
		if currentPath == "" {
			file, _ := exec.LookPath(os.Args[0])
			path, _ := filepath.Abs(file)
			index := strings.LastIndex(path, string(os.PathSeparator))
			currentPath = path[:index] + "/"
		}
	}
	return currentPath
}

func truncateID(id string) string {
	shortLen := 12
	if len(id) < shortLen {
		shortLen = len(id)
	}
	return id[:shortLen]
}

// GenerateRandomID returns an unique id
func GenerateRandomID() string {
	for {
		id := make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, id); err != nil {
			panic(err) // This shouldn't happen
		}
		value := hex.EncodeToString(id)
		// if we try to parse the truncated for as an int and we don't have
		// an error then the value is all numberic and causes issues when
		// used as a hostname. ref #3869
		if _, err := strconv.ParseInt(truncateID(value), 10, 32); err == nil {
			continue
		}
		return value
	}
}

func FindConf(confs []string) string {
	for _, conf := range confs {
		_, err := os.Stat(conf)
		if err == nil {
			return conf
		}
	}
	return ""
}
