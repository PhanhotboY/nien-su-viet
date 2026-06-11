package adto

type GetPlanByIdReqDto struct {
	Id string `json:"id" validate:"required,uuid"`
}
