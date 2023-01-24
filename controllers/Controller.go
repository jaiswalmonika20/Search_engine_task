package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/module_page/models"
	"github.com/module_page/services"
)

type PageController struct {
	PageService services.PageService
}

func New(pageservice services.PageService) PageController {
	return PageController{
		PageService: pageservice,
	}
}

// check api is live
func online(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "api is online"})
}

func (pagecontroller *PageController) CreateNewPage(ctx *gin.Context) {
	var page models.Page
	if err := ctx.ShouldBindJSON(&page); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	temp := strings.Split(page.Key, " ")
	//checking total number of keys
	if len(temp) <= 10 {

		err := pagecontroller.PageService.AddPage(&page)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "loaded into database"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "key length out of limit"})
	}

}

// get all data store in dbms
func (pagecontroller *PageController) GetAllPage(ctx *gin.Context) {
	pages, err := pagecontroller.PageService.GetAllPages()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pages)
}

func SortByPriority_Pages(mpp map[string]int) []string {
	keys := make([]string, 0, len(mpp))

	for key := range mpp {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return mpp[keys[i]] > mpp[keys[j]]
	})
	return keys
}

// adding new data in dbms

func (pagecontroller *PageController) GetByQuery(ctx *gin.Context) {

	var query string = ctx.Param("query")
	queries := strings.Split(query, " ")

	page, err := pagecontroller.PageService.GetAllPages()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	mpp := Calculate_rating(page, queries)

	ctx.JSON(http.StatusOK, SortByPriority_Pages(mpp))

}
func Calculate_rating(pages []*models.Page, queries []string) map[string]int {
	temp := []string{}
	for _, t := range pages {
		temp = append(temp, t.Key)
	}
	val := []int{}
	for i := range temp {
		var sum int = 0
		temp2 := strings.Split(temp[i], " ")
		for k := 0; k < len(temp2); k++ {
			for j := range queries {
				if strings.EqualFold(temp2[k], queries[j]) {
					sum += (10 - k) * (10 - j)
				}
			}
		}
		val = append(val, sum)
	}
	mpp := map[string]int{}
	for l, varr := range val {
		if varr != 0 {
			str := strconv.Itoa(l + 1)
			var pageNo string = "P" + str
			mpp[pageNo] = varr
		}
	}
	fmt.Println(mpp)
	return mpp
}

func (pagecontroller *PageController) Routes(route *gin.RouterGroup) {
	route.GET("/pages", pagecontroller.GetAllPage)
	route.GET("/:query", pagecontroller.GetByQuery)
	route.POST("/newpage", pagecontroller.CreateNewPage)
	route.GET("/", online)

}
