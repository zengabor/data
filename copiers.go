package data

import "fmt"

// Copies values, and sets flag if there was a change in the destination
type changeAwareCopier bool

func (c *changeAwareCopier) DestinationUpdated() bool {
	return bool(*c)
}

func (c *changeAwareCopier) copyString(src string, dst *string) {
	if *dst != src {
		*dst = src
		*c = true
	}
}

func (c *changeAwareCopier) copyStringIfAny(src string, dst *string) {
	if len(src) > 0 {
		c.copyString(src, dst)
	}
}

func (c *changeAwareCopier) copyStringIfMissing(src string, dst *string) {
	if len(*dst) == 0 && len(src) > 0 && *dst != src {
		*dst = src
		*c = true
	}
}

func (c *changeAwareCopier) copyStringWithWarning(src string, dst *string, fieldName string) string {
	if len(src) == 0 {
		return fmt.Sprintf("MISSING FIELD: %s", fieldName)
	}
	c.copyString(src, dst)
	return ""
}

func (c *changeAwareCopier) copyNilBool(src NilBool, dst *NilBool) {
	if *dst != src {
		*dst = src
		*c = true
	}
}

// Copies values and sets flag if destination differs from source
type diffAwareCopier bool

func (d *diffAwareCopier) copyStringIfMissing(src string, dst *string) {
	if len(*dst) == 0 && len(src) > 0 {
		*dst = src
	}
	// destination is still different, so we have something new here
	if *dst != src {
		*d = true
	}
}

func (d *diffAwareCopier) Different() bool {
	return bool(*d)
}
