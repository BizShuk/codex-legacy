package ui

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	. "propertychain/utils"
	"strings"
)

// Handler handle source
type Handler func() (readCloser io.ReadCloser)

// GetPropertySourceHandler factory
func GetProtocolHandler(uri string) (handler Handler) {
	protocol := GetProtocol(uri)
	switch protocol {
	case OS:
		handler = loadFromOS(uri)
	case HOME:
		handler = loadFromHomeEnv(uri)
	case HTTP:
		handler = loadFromHttp(uri)
	case FTP:
		handler = loadFromFtp(uri)
	default:
		handler = loadFromFile(uri)
	}
	return
}

var loadFromOS = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		lines := os.Environ()
		reader := strings.NewReader(strings.Join(lines, "\n"))
		readCloser = ioutil.NopCloser(reader)
		return
	}
}

// loadFromHomeEnv reads .env (or path relative to $HOME) from the user's home directory.
// Priority in chain should be lower than current-dir .env and OS env vars.
var loadFromHomeEnv = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		// uri looks like "home://.env" → strip scheme and join with $HOME
		rel := strings.TrimPrefix(uri, "home://")
		if rel == "" {
			rel = ".env"
		}
		home, err := os.UserHomeDir()
		if err != nil {
			log.Println("[propertySources] cannot resolve $HOME:", err)
			readCloser = ioutil.NopCloser(strings.NewReader(""))
			return
		}
		path := filepath.Join(home, rel)
		f, err := os.Open(path)
		if err != nil {
			// home .env is optional — silently skip if not present
			readCloser = ioutil.NopCloser(strings.NewReader(""))
			return
		}
		readCloser = f
		return
	}
}

var loadFromFile = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		readCloser, err := os.Open(uri)
		if err != nil {
			log.Fatal("File not found!!")
			return nil
		}

		return readCloser
	}
}

var loadFromHttp = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		resp, err := http.Get(uri)
		if err != nil {
			log.Fatal("Erro when fetch ", uri)
			return nil
		}
		readCloser = resp.Body
		return
	}
}

// loadFromFtp is a placeholder — real FTP needs a dedicated client library
// (e.g. github.com/jlaffaye/ftp). For now, log a warning and return an empty stream
// instead of incorrectly using os.Open (which would silently treat FTP URIs as local paths).
var loadFromFtp = func(uri string) Handler {
	return func() (readCloser io.ReadCloser) {
		log.Println("[propertySources] FTP source not implemented yet, skipping:", uri)
		readCloser = ioutil.NopCloser(strings.NewReader(""))
		return
	}
}
