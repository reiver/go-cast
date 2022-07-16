package stringlike

type Flag struct {
	value string
	loaded bool
}

func SomeFlag(value string) Flag {
	return Flag{
		value:value,
		loaded:true,
	}
}

func (receiver *Flag) Set(value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	*receiver = Flag{
		value:value,
		loaded:true,
	}

	return nil
}
