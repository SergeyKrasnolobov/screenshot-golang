package api

const (
	ErrorEmptySource = "Поле source не может быть пустым"
	ErrorValidation  = "Ошибка обработки запроса, попробуйте чуть позже"
	ErrorInternal    = "internal error"
	ErrorUnknown     = "Неизвестная ошибка"
)

const (
	BadRequestStatus          = "400"
	InternalServerErrorStatus = "500"
)
