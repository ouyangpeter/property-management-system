package responses

import "github.com/beevik/guid"

type BaseResponse struct {
    RequestId string `json:"RequestId"`
}

func NewBaseResponse() *BaseResponse {
    b := &BaseResponse{}

    b.RequestId = guid.NewString()
    return b
}
