package view_models

type SuccessResponseWithOutMetaVm struct {
	Data interface{} `json:"data"`
}

type SuccessResponseWithMeta struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

type ErrorResponseVm struct {
	Message interface{}
}
