/*
  Artisan's Doorman - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func init() {
	// load env vars from file if present
	godotenv.Load("proxy.env")
}

func IsLoggingEnabled() bool {
	value := os.Getenv("DPROXY_LOGGING")
	return len(value) > 0
}

func GetSourceURI() (string, error) {
	value := os.Getenv("DPROXY_SOURCE_URI")
	if len(value) == 0 {
		return "", fmt.Errorf("missing DPROXY_SOURCE_URI variable")
	}
	return value, nil
}

func GetSourceUser() (string, error) {
	value := os.Getenv("DPROXY_SOURCE_USER")
	if len(value) == 0 {
		return "", fmt.Errorf("missing DPROXY_SOURCE_USER variable")
	}
	return value, nil
}

func GetSourcePwd() (string, error) {
	value := os.Getenv("DPROXY_SOURCE_PASSWORD")
	if len(value) == 0 {
		return "", fmt.Errorf("missing DPROXY_SOURCE_PASSWORD variable")
	}
	return value, nil
}

func GetSourceInsecureSkipVerify() bool {
	value := os.Getenv("DPROXY_SOURCE_INSECURE_SKIP_VERIFY")
	if len(value) == 0 {
		return false
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return v
}
