package image

import (
	"errors"
	"runtime"
	"strings"

	"github.com/moby/moby/v2/errdefs"
)

// HostOS returns the OCI operating system name for this host.
// Android uses the Linux kernel and runs Linux OCI container images.
func HostOS() string {
	if runtime.GOOS == "android" {
		return "linux"
	}
	return runtime.GOOS
}

// CheckOS checks if the given OS matches the host's platform.
func CheckOS(os string) error {
	if !strings.EqualFold(HostOS(), os) {
		return errdefs.InvalidParameter(errors.New("operating system is not supported"))
	}
	return nil
}
