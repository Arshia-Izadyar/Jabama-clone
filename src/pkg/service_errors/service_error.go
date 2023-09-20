package service_errors

type ServiceError struct {
	EndUserMessage string `json:"end_user_message"`
	Err            error  `json:"err"`
}

func (se *ServiceError) Error() string {
	return se.EndUserMessage
}
