package internal

import (
	"runtime/debug"
	"sync"
)

const (
	defaultChalkVersion = "dev"
	userAgentPrefix     = "chalk-go/"
)

var (
	userAgentOnce   sync.Once
	cachedUserAgent string
)

func UserAgent() string {
	userAgentOnce.Do(func() {
		cachedUserAgent = userAgentPrefix + chalkVersion()
	})
	return cachedUserAgent
}

func chalkVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return defaultChalkVersion
	}
	version := info.Main.Version
	if version == "" || version == "(devel)" {
		return defaultChalkVersion
	}
	return version
}
