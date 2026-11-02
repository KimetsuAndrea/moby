//go:build !windows && !darwin && !freebsd

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package platforms

import (
	"runtime"

	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

func DefaultSpec() specs.Platform {
	os := runtime.GOOS

	// Android uses the Linux kernel and runs Linux OCI container images.
	// The Go binary is built for Android, but the container platform is Linux.
	if os == "android" {
		os = "linux"
	}

	return specs.Platform{
		OS:           os,
		Architecture: runtime.GOARCH,
		Variant:      cpuVariant(),
	}
}

// Default returns the default matcher for the platform.
func Default() MatchComparer {
	return Only(DefaultSpec())
}
