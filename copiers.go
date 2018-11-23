package data

import "fmt"

// Copies values, and sets flag if there was a change in the destination
type ChangeAwareCopier bool

func (c *ChangeAwareCopier) DestinationUpdated() bool {
	return bool(*c)
}

func (c *ChangeAwareCopier) CopyString(src string, dst *string) {
	if *dst != src {
		*dst = src
		*c = true
	}
}

func (c *ChangeAwareCopier) CopyStringIfAny(src string, dst *string) {
	if len(src) > 0 {
		c.copyString(src, dst)
	}
}

func (c *ChangeAwareCopier) CopyStringIfMissing(src string, dst *string) {
	if len(*dst) == 0 && len(src) > 0 && *dst != src {
		*dst = src
		*c = true
	}
}

func (c *ChangeAwareCopier) CopyStringWithWarning(src string, dst *string, fieldName string) string {
	if len(src) == 0 {
		return fmt.Sprintf("MISSING FIELD: %s", fieldName)
	}
	c.copyString(src, dst)
	return ""
}

func (c *ChangeAwareCopier) CopyNilBool(src NilBool, dst *NilBool) {
	if *dst != src {
		*dst = src
		*c = true
	}
}

// Copies values and sets flag if destination differs from source
type DiffAwareCopier bool

func (d *DiffAwareCopier) CopyStringIfMissing(src string, dst *string) {
	if len(*dst) == 0 && len(src) > 0 {
		*dst = src
	}
	// destination is still different, so we have something new here
	if *dst != src {
		*d = true
	}
}

func (d *DiffAwareCopier) Different() bool {
	return bool(*d)
}
