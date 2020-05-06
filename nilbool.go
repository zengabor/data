package data

type NilBool int32

const (
	NilBoolUnset NilBool = 0
	NilBoolTrue  NilBool = 1
	NilBoolFalse NilBool = -1
)

func NilBoolFromBool(value bool) NilBool {
	if value {
		return NilBoolTrue
	} else {
		return NilBoolFalse
	}
}

func NilBoolFromString(value string) NilBool {
	switch value {
	case "true":
		return NilBoolTrue
	case "false":
		return NilBoolFalse
	default:
		return NilBoolUnset
	}
}

func (nb NilBool) Unset() {
	nb = NilBoolUnset
}

func (nb NilBool) Set(value bool) {
	nb = NilBoolFromBool(value)
}

func (nb NilBool) IsSet() bool {
	return nb == NilBoolTrue || nb == NilBoolFalse
}

func (nb NilBool) IsTrue() bool {
	return nb == NilBoolTrue
}

func (nb NilBool) IsFalse() bool {
	return nb == NilBoolFalse
}
