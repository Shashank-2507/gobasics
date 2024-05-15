package helper

import (
	"github.com/gofiber/fiber/v2/utils"
)

func Genid() string {
	return utils.UUID()
}
