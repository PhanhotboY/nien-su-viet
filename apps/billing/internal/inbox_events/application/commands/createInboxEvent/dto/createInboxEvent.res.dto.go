package adto // application command dto

type CreateInboxEventResDto interface {
	GetData() *CreateInboxEventResData
}

type CreateInboxEventResData struct {
	Id string `json:"id"`
}

type createInboxEventResDto struct {
	Data *CreateInboxEventResData `json:"data"`
}

func NewCreateInboxEventResDto(id string) CreateInboxEventResDto {
	return &createInboxEventResDto{
		Data: &CreateInboxEventResData{
			Id: id,
		},
	}
}

func (d *createInboxEventResDto) GetData() *CreateInboxEventResData {
	return d.Data
}
