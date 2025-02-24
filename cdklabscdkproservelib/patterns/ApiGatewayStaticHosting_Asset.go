package patterns


// Static Asset Definition.
// Experimental.
type ApiGatewayStaticHosting_Asset struct {
	// Unique identifier to delineate an asset from other assets.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Path(s) on the local file system to the static asset(s).
	//
	// Each path must be either a directory or zip containing the assets.
	// Experimental.
	Path interface{} `field:"required" json:"path" yaml:"path"`
	// Name of the index page for a Single Page Application (SPA).
	//
	// This is used as a default key to load when the path provided does not map to a concrete static asset.
	//
	// Example:
	//   index.html
	//
	// Experimental.
	SpaIndexPageName *string `field:"optional" json:"spaIndexPageName" yaml:"spaIndexPageName"`
}

