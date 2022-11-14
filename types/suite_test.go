package types

import (
	"encoding/json"
	"os"
	"southwinds.dev/types/dproxy"
	"testing"
)

func TestSaveRelease(t *testing.T) {
	r := dproxy.WebhookInfo{
		WebhookToken: "JFkxnsn++02UilVkYFFC9w==",
		ReferrerURL:  "",
		IpSafeList:   []string{"127.0.0.1"},
		Filter:       "",
	}
	b, _ := json.MarshalIndent(r, "", "  ")
	os.WriteFile("webhookInfo.json", b, os.ModePerm)
}
