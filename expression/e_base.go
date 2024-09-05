package expression

type EFunction interface {
	Call(args []EValue) (EValue, error)
	ToString() EString
}

type EValue interface {
	ToString() EString
}

type EObject interface {
	EValue
	Get(name string) (EValue, error)
	Keys() []string
}
