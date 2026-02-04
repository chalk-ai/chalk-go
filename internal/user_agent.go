package internal

import (
	"runtime/debug"
	"sync"
)

const (
	chalkModulePath     = "github.com/chalk-ai/chalk-go"
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
	if v := versionFromBuildInfo(); v != "" {
		return v
	}
	return defaultChalkVersion
}

func versionFromBuildInfo() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	if v := moduleVersion(info.Main); v != "" {
		return v
	}
	for _, dep := range info.Deps {
		if dep == nil {
			continue
		}
		if v := moduleVersion(*dep); v != "" {
			return v
		}
	}
	return ""
}

func moduleVersion(mod debug.Module) string {
	if mod.Path != chalkModulePath {
		return ""
	}
	version := mod.Version
	if version == "" && mod.Replace != nil {
		version = mod.Replace.Version
	}
	if version == "" || version == "(devel)" {
		return ""
	}
	return version
}
