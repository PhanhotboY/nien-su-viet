package dbcontracts

import "go.uber.org/fx"

type DbModelParam struct {
	Order int
	Model any
}

func NewDbModelParam(order int, model any) any {
	return fx.Annotate(
		func() DbModelParam {
			return DbModelParam{
				Order: order,
				Model: model,
			}
		},
		fx.ResultTags(`group:"db_models"`),
	)
}
