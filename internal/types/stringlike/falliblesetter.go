package stringlike

type FallibleSetter struct {
	value string
	loaded bool
}

func SomeFallibleSetter(value string) FallibleSetter {
	return FallibleSetter{
		value:value,
		loaded:true,
	}
}

func (receiver *FallibleSetter) SetString(value string) bool {
	if nil == receiver {
		return false
	}

	*receiver = FallibleSetter{
		value:value,
		loaded:true,
	}

	return true
}
