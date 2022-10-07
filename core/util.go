/*
  Doorman Proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

import (
	src "southwinds.dev/source_client"
	"time"
)

func GetSource() (*src.Client, error) {
	uri, err := GetSourceURI()
	if err != nil {
		return nil, err
	}
	user, err := GetSourceUser()
	if err != nil {
		return nil, err
	}
	pwd, err := GetSourcePwd()
	if err != nil {
		return nil, err
	}
	s := src.New(uri, user, pwd, &src.ClientOptions{
		InsecureSkipVerify: GetSourceInsecureSkipVerify(),
		Timeout:            60 * time.Second,
	})
	return &s, nil
}
