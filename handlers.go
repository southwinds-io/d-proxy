/*
  Doorman Proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"southwinds.dev/d-proxy/core"
	"southwinds.dev/d-proxy/types"
	util "southwinds.dev/http"
	"southwinds.dev/types/dproxy"
	"strings"
	"time"
)

// @title Doorman Proxy
// @version 1.0.0
// @description Event Sources for Doorman
// @contact.name SouthWinds Tech ltd
// @contact.url https://www.southwinds.io/
// @contact.email info@southwinds.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @Summary A Webhook for MinIO compatible event sources
// @Description receives a s3:ObjectCreated:Put event sent by a MinIO format compatible source
// @Tags Event Sources
// @Router /events/minio [post]
// @Param event body types.MinioS3Event true "the notification information to send"
// @Accept application/json, application/yaml
// @Produce plain
// @Failure 400 {string} bad request: the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)
// @Failure 500 {string} internal server error: the server encountered an unexpected condition that prevented it from fulfilling the request.
// @Success 201 {string} event has been processed
func minioEventsHandler(w http.ResponseWriter, r *http.Request) {
	event := new(types.MinioS3Event)
	err := util.Unmarshal(r, event)
	if util.IsErr(w, err, http.StatusBadRequest, "cannot unmarshal webhook payload") {
		return
	}
	if event.Records == nil {
		util.Err(w, http.StatusBadRequest, "incorrect webhook payload, missing Records, cannot continue")
		return
	}
	object := event.Records[0].S3.Object
	if !strings.HasSuffix(object.Key, "spec.yaml") {
		util.Err(w, http.StatusBadRequest, fmt.Sprintf("invalid event, changed object was %s but required spec.yaml", object.Key))
		return
	}
	key, err := url.PathUnescape(object.Key)
	if util.IsErr(w, err, http.StatusBadRequest, fmt.Sprintf("cannot unescape object key %s", object.Key)) {
		return
	}
	// checks if the release has been done within a folder, if not it is not valid
	if !strings.Contains(key, "/") {
		util.Err(w, http.StatusBadRequest, "no release folder specified within bucket \n"+
			"(i.e. format should be s3host://bucket-name/version-folder/ it it was s3host://bucket-name, cannot accept event; \n"+
			"ensure you put objects in the bucket under a version folder\n")
		return
	}
	cut := strings.LastIndex(key, "/")
	// get the path within the bucket
	folderName := key[:cut]
	// get the unique identifier for the bucket
	deploymentId := event.Records[0].ResponseElements.XMinioDeploymentID
	// get the bucket name
	bucketName := event.Records[0].S3.Bucket.Name
	// output new release information
	fmt.Printf("︎⚡️ new release:\n")
	fmt.Printf("  ✔ from   = %s\n", event.Records[0].ResponseElements.XMinioOriginEndpoint)
	fmt.Printf("  ✔ time   = %s\n", time.Now().UTC())
	fmt.Printf("  ✔ id     = %s\n", deploymentId)
	fmt.Printf("  ✔ bucket = %s\n", bucketName)
	fmt.Printf("  ✔ folder = %s\n", folderName)
	fmt.Printf("  ✔ type   = minio compatible\n")
	fmt.Printf("--------------------------------------------------------------------\n\n")

	s, err := core.GetSource()
	if util.IsErr(w, err, http.StatusInternalServerError, fmt.Sprintf("cannot connect to source")) {
		return
	}
	err = s.Save(dproxy.ReleaseKey, dproxy.ReleaseType, dproxy.Release{
		Origin:       event.Records[0].ResponseElements.XMinioOriginEndpoint,
		DeploymentId: deploymentId,
		BucketName:   bucketName,
		FolderName:   folderName,
		Time:         time.Now().UTC(),
	})
	if util.IsErr(w, err, http.StatusInternalServerError, fmt.Sprintf("cannot persist release information")) {
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// @Summary Retrieves next release to process
// @Description Doorman uses this endpoint to get the next release to process or none (404)
// @Tags Event Sources
// @Router /release [get]
// @Accept application/json, application/yaml
// @Produce json
// @Failure 404 {string} not found: no releases found
// @Failure 400 {string} bad request: the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)
// @Failure 500 {string} internal server error: the server encountered an unexpected condition that prevented it from fulfilling the request.
// @Success 200 {string} request has been successful
func getReleaseHandler(w http.ResponseWriter, r *http.Request) {
	s, err := core.GetSource()
	if util.IsErr(w, err, http.StatusInternalServerError, fmt.Sprintf("cannot connect to source")) {
		return
	}
	release, err := s.PopOldest(dproxy.ReleaseType, new(dproxy.Release))
	if util.IsErr(w, err, http.StatusInternalServerError, fmt.Sprintf("cannot persist release information")) {
		return
	}
	if release == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	util.Write(w, r, release)
}
