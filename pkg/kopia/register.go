// Copyright 2021 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package kopia provides integration with Kopia, a fast and secure backup tool.
// It registers supported blob storage providers for use with Kopia repositories.
package kopia

// Register supported blob storage providers
import (
	_ "github.com/kopia/kopia/repo/blob/azure"      // Azure blob storage support
	_ "github.com/kopia/kopia/repo/blob/filesystem" // Filesystem storage support
	_ "github.com/kopia/kopia/repo/blob/gcs"        // Google Cloud Storage support
	_ "github.com/kopia/kopia/repo/blob/s3"         // S3-compatible storage support
)
