package mErr

type ErrResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
}

type Error struct {
	Message  string
	Error    error
	Response *ErrResponse
}

func ErrorHandler(message string, error error, response *ErrResponse) *Error {
	return &Error{Message: message, Response: response, Error: error}
}
