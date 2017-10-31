package example

type Stack struct {
	Callers []uintptr
}

func (t Stack) Marshal() ([]byte, error) {
	return nil, nil
}

func (t *Stack) MarshalTo(data []byte) (n int, err error) {
	return 0, nil
}

func (t *Stack) Unmarshal(data []byte) error {
	return nil
}

func (t Stack) MarshalJSON() ([]byte, error) {
	return []byte(`null`), nil
}

func (t *Stack) UnmarshalJSON(data []byte) error {
	return nil
}

func (t *Stack) Size() int {
	return 0
}
