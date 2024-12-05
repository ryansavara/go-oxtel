package oxtel

import "fmt"

type BaseError struct {
	Message string
}
type InvalidRateError struct {
	BaseError
}

type InvalidAngleError struct {
	BaseError
}

type InvalidDurationError struct {
	BaseError
}

type InvalidMixError struct {
	BaseError
}

type InvalidFieldError struct {
	BaseError
}

type InvalidGainError struct {
	BaseError
}

type InvalidSdiError struct {
	BaseError
}

type InvalidAudioProfileError struct {
	BaseError
}

type TimeoutError struct {
	BaseError
}

type InvalidDolbyProfileError struct {
	BaseError
}

type InvalidAudioChannelError struct {
	BaseError
}

type InvalidParametersError struct {
	BaseError
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}
