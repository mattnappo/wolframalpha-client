package common

var (
	// UnsafeChars are the chars that are not safe within a url.
	UnsafeChars = []string{
		"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",",
		"/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^",
		"`", "{", "|", "}",
	}

	// SafeChars are the safe chars, respectively.
	SafeChars = []string{
		"%21", "%22", "%23", "%24", "%25", "%26", "%27", "%28", "%29",
		"%2A", "%2B", "%2C", "%2F", "%3A", "%3B", "%3C", "%3D", "%3E",
		"%3F", "%40", "%5B", "%5C", "%5D", "%5E", "%5F", "%60",
	}
)
