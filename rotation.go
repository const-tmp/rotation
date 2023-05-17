package rotation

// Trigger triggers rotation
type Trigger <-chan struct{}

type Rotator interface {
	// Rotate rotates resource
	Rotate()
}

type ValueGetter interface {
	// GetValue returns most recent resource value
	GetValue() string
}

// Rotate starts rotation process
func Rotate(trigger Trigger, rotators ...Rotator) {
	for range trigger {
		for _, r := range rotators {
			r.Rotate()
		}
	}
}
