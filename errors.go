package cim

type NotImplementedCimObjectError struct {
	Message string
}

func (error NotImplementedCimObjectError) Error() string {
	return "Not implemented unmarshalling of: " + error.Message
}
