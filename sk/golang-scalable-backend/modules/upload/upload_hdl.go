package upload

import (
	"clean-architecture/sk/golang-scalable-backend/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Upload(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			panic(common.ErrValidation(err))
		}

		dst := fmt.Sprintf("./uploads/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)

		if err := ctx.SaveUploadedFile(fileHeader, dst); err != nil {
		}

		img := common.Image{
			Id:        0,
			Url:       dst,
			Width:     100,
			Height:    100,
			CloudName: "local",
			Extension: "",
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
