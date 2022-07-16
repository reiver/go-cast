package stringlike

type Setter struct {
	value string
	loaded bool
}

func SomeSetter(value string) Setter {
	return Setter{
		value:value,
		loaded:true,
	}
}

func (receiver *Setter) SetString(value string) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	*receiver = SomeSetter(value)
}
