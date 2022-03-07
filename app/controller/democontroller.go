package controller

import (
	"carefree/app/agreement"
	"fmt"
	"github.com/carefreex-io/xhttp"
	"github.com/gin-gonic/gin"
	"sync"
)

type DemoController struct {
}

var (
	demoController     *DemoController
	demoControllerOnce sync.Once
)

func NewDemoController() *DemoController {
	if demoController == nil {
		demoControllerOnce.Do(func() {
			demoController = &DemoController{}
		})
	}

	return demoController
}

func (c *DemoController) Home(ctx *gin.Context) {
	params := agreement.HomeRequest{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		xhttp.BadRequestResponse(ctx, err)
		return
	}

	fmt.Println("params: ", params)

	xhttp.SuccessResponse(ctx, agreement.HomeResponse{Uid: params.Uid})
}

func (c *DemoController) Test(ctx *gin.Context) {
	fmt.Println("this is DemoController.Test method")

	xhttp.SuccessResponse(ctx, gin.H{
		"info": "this is DemoController.Test method",
	})
}
