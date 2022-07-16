package stringlike

type ErrableSetter struct {
	value string
	loaded bool
}

func SomeErrableSetter(value string) ErrableSetter {
	return ErrableSetter{
		value:value,
		loaded:true,
	}
}

func (receiver *ErrableSetter) SetString(value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	*receiver = ErrableSetter{
		value:value,
		loaded:true,
	}

	return nil
}
