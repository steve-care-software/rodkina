package programs

type attachment struct {
	value Value
	local uint
}

func createAttachment(
	value Value,
	local uint,
) Attachment {
	out := attachment{
		value: value,
		local: local,
	}

	return &out
}

// Value returns the value
func (obj *attachment) Value() Value {
	return obj.value
}

// Local returns the local
func (obj *attachment) Local() uint {
	return obj.local
}
