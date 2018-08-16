package container

import "github.com/aghape/helpers"

var PREFIX string

func init() {
	PREFIX = helpers.GetCalledDir()
}
