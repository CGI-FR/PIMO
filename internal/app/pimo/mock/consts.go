package mock

const (
	keyMethod      = "method"   // request only
	keyStatus      = "status"   // response only
	keyURL         = "url"      // request only
	keyURLPath     = "path"     // request only
	keyURLQuery    = "query"    // request only
	keyURLFragment = "fragment" // request only
	keyProtocol    = "protocol"
	keyHeaders     = "headers"
	keyBody        = "body"
	keyTrailers    = "trailers" // response only

	// following properties are always empty and are not keyed
	// keyURLScheme       = "scheme"
	// keyURLUser         = "user"
	// keyURLUserName     = "name"
	// keyURLUserPassword = "pass"
	// keyURLHost         = "host"
	// keyURLHostName     = "name"
	// keyURLHostPort     = "port"
)
