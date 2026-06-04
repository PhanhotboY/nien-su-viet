package adto

type GetPurchaseReqDto struct {
	PurchaseId string `json:"id" validate:"required"`
}
