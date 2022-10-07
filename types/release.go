/*
  Doorman Proxy - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package types

import "time"

type Release struct {
	Origin       string    `json:"origin"`
	DeploymentId string    `json:"deployment_id"`
	BucketName   string    `json:"bucket_name"`
	FolderName   string    `json:"folder_name"`
	Time         time.Time `json:"time"`
}
