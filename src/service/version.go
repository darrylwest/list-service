// simple version
//
// @author darryl.west@ebay.com
// @created 2018-01-15 09:59:37

package service

import "fmt"

const (
	major = 18
	minor = 1
	patch = 18
)

// Version - return the version number as a single string
func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
