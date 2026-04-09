package dto

type BaseRequest struct {
	IdInstance       string `json:"idInstance"       schema:"idInstance"       validate:"required"`
	APITokenInstance string `json:"apiTokenInstance" schema:"apiTokenInstance" validate:"required"`
}

type SendMessageRequest struct {
	BaseRequest

	ChatID  string `json:"chatId"  validate:"required"`
	Message string `json:"message" validate:"required"`
}

type SendFileRequest struct {
	BaseRequest

	ChatID   string `json:"chatId"   validate:"required"`
	FileName string `json:"fileName" validate:"required"`
	UrlFile  string `json:"urlFile"  validate:"required"`
	Caption  string `json:"caption"  validate:"required"`
}
