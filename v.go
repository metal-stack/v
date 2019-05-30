package v

type version struct{}

var (
	// Version of the application
	Version   = ""
	// Revision of the application
	Revision  = ""
	// GitSHA1 of the application
	GitSHA1   = ""
	// BuildDate of the application
	BuildDate = ""
)

// V the version
var V = &version{}

func (v *version) String() string {
	var versionString = Version
	if GitSHA1 != "" {
		versionString += " (" + GitSHA1 + ")"
	}
	if Revision != "" {
		versionString += ", " + Revision
	}
	if BuildDate != "" {
		versionString += ", " + BuildDate
	}
	return versionString
}
