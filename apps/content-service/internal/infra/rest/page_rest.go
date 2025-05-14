package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
)

type PageRest struct {
	page.PageService
}

func NewPageRest(rg *gin.RouterGroup, pageService page.PageService) {
	pageRest := &PageRest{
		PageService: pageService,
	}

	rg.POST("/", pageRest.post)
	rg.GET("/", pageRest.get)
}

func (pr *PageRest) post(ctx *gin.Context) {
	page, err := pr.PageService.CreateDraftPage()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"page": page,
	})
}

func (pr *PageRest) get(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"page": page.NewDraft(),
	})
}
