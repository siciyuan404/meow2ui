package provider

type ProviderError struct {
	Code       string
	ProviderID string
	ModelID    string
	Retryable  bool
	Message    string
}

func (e *ProviderError) Error() string {
	return e.Message
}
