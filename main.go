/*
  Doorman Proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"southwinds.dev/d-proxy/core"
	"southwinds.dev/d-proxy/types"
	h "southwinds.dev/http"
	"strings"
	"time"
)

var (
	defaultAuth func(r http.Request) *h.UserPrincipal
)

func main() {
	// creates a generic http server
	s := h.New("doorman-proxy", types.Version)
	// add handlers
	s.Http = func(router *mux.Router) {
		// add http request login for debugging purposes (using DPROXY_LOGGING env variable)
		if core.IsLoggingEnabled() {
			router.Use(s.LoggingMiddleware)
		}
		// apply authentication
		router.Use(s.AuthenticationMiddleware)

		router.HandleFunc("/events/minio", minioEventsHandler).Methods("POST")
	}
	// grab a reference to default auth to use it in the proxy override below
	defaultAuth = s.DefaultAuth
	// set up specific authentication for doorman proxy
	s.Auth = map[string]func(http.Request) *h.UserPrincipal{
		"^/events/.*": whAuth,
	}
	fmt.Print(`
+++++++++++++++++++++++++++++++++++++++++++++
|     ___   ___   ___   ___   _     _       |
|    | | \ | |_) | |_) / / \ \ \_/ \ \_/    |
|    |_|_/ |_|   |_| \ \_\_/ /_/ \  |_|     |
+++++++++++|  doorman's proxy  |+++++++++++++
`)
	// ensure source is set up for handling configuration
	if err := core.SetMeta(); err != nil {
		log.Fatalf(err.Error())
	}
	s.Serve()
}

// whAuth authenticates web hook requests using opaque string (bearer token)
func whAuth(r http.Request) *h.UserPrincipal {
	// load the configuration info on every request to prevent caching issues
	whInfo, err := core.LoadWebHookInfo()
	if err != nil {
		log.Printf("ERROR: cannot load webhook configuration, cannot authenticate request: %s", err)
		return nil
	}
	ip := h.FindRealIP(&r)
	token := r.Header.Get("Authorization")
	for _, info := range whInfo {
		// if the bearer token is ok
		if strings.HasSuffix(token, info.WebhookToken) {
			var safeListed bool
			// if a whitelist has been set up for the webhook
			if info.IpSafeList != nil && len(info.IpSafeList) > 0 {
				// check that the requester real IP is in the whitelist
				for _, listedIp := range info.IpSafeList {
					if ip == listedIp {
						safeListed = true
					}
				}
				// if it is not then block the request
				if !safeListed {
					log.Printf("WARNING: authentication failed, requester IP '%s' is not safe listed\n", ip)
					return nil
				}
			}
			// if a referrer URL has been specified
			if len(info.ReferrerURL) > 0 {
				// if the referrer in the request does not match the required one
				if strings.EqualFold(r.Referer(), info.ReferrerURL) {
					log.Printf("WARNING: authentication failed, referrer URL '%s' does not match required value '%s'\n", r.Referer(), info.ReferrerURL)
					//  blocks the request
					return nil
				}
			}
			return &h.UserPrincipal{
				Username: "webhook-user",
				Created:  time.Now(),
				Context:  token,
			}
		}
	}
	// try with admin credentials
	if defaultAuth != nil {
		if principal := defaultAuth(r); principal != nil {
			return principal
		}
	}
	// otherwise, fail authentication
	log.Printf("WARNING: authentication failed, invalid token '%s'\n", token)
	return nil
}
