package auth

// MicroTimerMock ...
type MicroTimerMock struct {
	result string
}

// Get returns prepared result of function running
func (m *MicroTimerMock) Get() string {
	return m.result
}
