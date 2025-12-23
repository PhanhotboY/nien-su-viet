package initialize

import (
	"encoding/json"
	"os"

	"github.com/danielgtaylor/huma/v2"
)

func ExportOpenAPI(api huma.API) error {
	data, err := json.MarshalIndent(api.OpenAPI(), "", "  ")
	if err != nil {
		return err
	}
	outDir := "../../openapi"
	err = os.Mkdir(outDir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return os.WriteFile(outDir+"/cms-service.json", data, 0644)
}
