package constant

import (
	"time"
)

const (
	// DefaultUser is used if there is no user given
	DefaultUser = "Administrator"

	// DefaultPort is used if there is no port given
	DefaultPort = 5985

	// DefaultScriptPath is used as the path to copy the file to
	// for remote execution if not provided otherwise.
	DefaultScriptPath = "C:/Temp/terraform_%RAND%.cmd"

	// DefaultTimeout is used if there is no timeout given
	DefaultTimeout = 5 * time.Minute

	// DefaultPassword is used if there is no password given
	DefaultPassword = "hello123"

	// Region is used if there is no region given
	DefaultRegion = "cn-hangzhou"

	// Defaultfile is used if there is no file given
	Defaultfile = "application.yml"
)
