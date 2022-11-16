/*
  Doorman Proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package core

import (
	"southwinds.dev/types/dproxy"
)

func LoadWebHookInfo() ([]*dproxy.WebhookInfo, error) {
	source, err := GetSource()
	if err != nil {
		return nil, err
	}
	items, err := source.LoadItemsByType(func() any {
		return new(dproxy.WebhookInfo)
	}, dproxy.WebHookInfoType)
	if err != nil {
		return nil, err
	}
	result := make([]*dproxy.WebhookInfo, 0)
	for _, item := range items {
		result = append(result, item.(*dproxy.WebhookInfo))
	}
	return result, nil
}

func SetMeta() error {
	source, err := GetSource()
	if err != nil {
		return err
	}
	err = source.SetType(dproxy.WebHookInfoType, dproxy.WebHookInfoProto)
	if err != nil {
		return err
	}
	return source.SetType(dproxy.ReleaseType, dproxy.ReleaseProto)
}
