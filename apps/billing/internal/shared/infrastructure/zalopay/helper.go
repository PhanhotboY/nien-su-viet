package zalopay

import (
	"fmt"
	"time"
)

func GenerateAppTransID() string {
	return fmt.Sprintf("%s_%s", time.Now().Format("060102"), time.Now().Format("20060102150405"))
}
