package expression

type EBytes []byte

func (v EBytes) Keys() []string {
	panic("implement me")
}

func (v EBytes) Get(name string) (EValue, error) {
	switch name {
	case "bcontains":
		return &eBytesBcontains{b: v, ci: false}, nil
	case "icontains":
		return &eBytesBcontains{b: v, ci: true}, nil
	}
	panic("implement me")
}

func (v EBytes) ToString() EString {
	return EString(v)
}
