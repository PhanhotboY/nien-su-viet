package dto

type GetEventQueryReq struct {
	ID string `json:"id" validate:"required"`
}
