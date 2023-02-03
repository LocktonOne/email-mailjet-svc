//go:build appengine || go1.5
// +build appengine go1.5

package envconfig

import "os"

func lookupEnv(key string) (string, bool) {
	value, found := os.LookupEnv(key)
	if found {
		os.Unsetenv(key)
	}

	return value, found
}
