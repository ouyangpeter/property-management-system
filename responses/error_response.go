package responses

import (
    "encoding/json"
    "github.com/astaxie/beego/context"
)

type ErrorBaseResponse struct {
    BaseResponse
    Status  int    `json:"Status"`
    Message string `json:"Message"`
}

func (response *ErrorBaseResponse) Handler(output *context.BeegoOutput) {
    s, _ := json.Marshal(response)

    output.ContentType("json")
    output.SetStatus(response.Status)
    output.Body(s)
}

func NewErrorBaseResponse(status int, message string) *ErrorBaseResponse {
    errResponse := &ErrorBaseResponse{BaseResponse: *NewBaseResponse(), Status: status, Message: message}
    return errResponse
}

func NewInvalidParameterResponse() *ErrorBaseResponse {
    return NewErrorBaseResponse(400, "InvalidParameter")
}

func NewForbiddenResponse() *ErrorBaseResponse {
    return NewErrorBaseResponse(403, "Forbidden")
}

func NewUnauthorizedResponse() *ErrorBaseResponse {
    return NewErrorBaseResponse(401, "Unauthorized")
}
