package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/module_page/mocks/services"
	"github.com/module_page/pkg/models"
	"github.com/module_page/pkg/services"
)

var pagecontroller *PageController

func TestPageController_online(t *testing.T) {

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	online(context)
	if response.Code != http.StatusOK {
		t.Errorf("api not active")
	}

}

func TestSortByPriority_Pages(t *testing.T) {
	type args struct {
		mpp map[string]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "queryTest",
			args: args{
				map[string]int{
					"P": 1,
					"Q": 2,
					"R": 7,
					"S": 4,
				},
			},
			want: []string{"R", "S", "Q", "P"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortByPriority_Pages(tt.args.mpp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortByPriority_Pages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		pageservice services.PageService
	}
	tests := []struct {
		name string
		args args
		want PageController
	}{
		{
			name: "testing New ",
			args: args{
				nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.pageservice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCalculate_rating(t *testing.T) {
	type args struct {
		pages   []*models.Page
		queries []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "calculating rating",
			args: args{
				pages: []*models.Page{
					{ID: 1, Key: "ford car"},
					{ID: 2, Key: "ford review"},
					{ID: 3, Key: "car ford review"},
				},
				queries: []string{
					"ford",
				},
			},
			want: map[string]int{
				"P1": 100,
				"P2": 100,
				"P3": 90,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate_rating(tt.args.pages, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate_rating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageController_CreateNewPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	response := httptest.NewRecorder()

	page := models.Page{
		ID:  1,
		Key: "lamborgini",
	}
	ServiceMock := mocks.NewPageService(t)
	ServiceMock.On("AddPage", page).Return(nil)

	router.POST("/newpage", pagecontroller.CreateNewPage)
	input := `{}`
	req, err := http.NewRequest("POST", "/newpage", bytes.NewBuffer([]byte(input)))
	if err != nil {
		log.Println(err)
	}
	router.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Errorf("data loaded test successful")
	}

}
func TestPageController_GetAllPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	response := httptest.NewRecorder()
	mockRepo := mocks.NewPageService(t)
	mockRepo.On("GetAllPages").Return([]models.Page{}, nil)
	if response.Code != http.StatusOK {
		t.Errorf("database not connected")
	}
	router.GET("/pages", pagecontroller.GetAllPage)
	request, err := http.NewRequest("GET", "/pages", nil)
	if err != nil {
		log.Println(err)
	} else {

	}
	router.ServeHTTP(response, request)

}

func TestPageController_GetByQuery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockRepo := mocks.NewPageService(t)
	mockRepo.On("GetAllPages").Return([]models.Page{}, nil)
	expectedMap := map[string]int{
		"P1": 100, "P2": 100, "P3": 90,
	}
	pages := []*models.Page{
		{ID: 1, Key: "ford car"},
		{ID: 2, Key: "ford review"},
		{ID: 3, Key: "car ford review"},
	}
	queries := []string{
		"ford",
	}
	response := httptest.NewRecorder()
	fmt.Println(response)
	actualMap := Calculate_rating(pages, queries)
	if got := Calculate_rating(pages, queries); !reflect.DeepEqual(got, expectedMap) {
		t.Errorf("Calculate_rating() = %v, want %v", got, expectedMap)
	}
	expectedSortedAns := []string{
		"P1", "P2", "P3",
	}
	router.GET("/:query", pagecontroller.GetByQuery)
	request, err := http.NewRequest("GET", "/:query", nil)
	mockRepo.On("SortByPriority_Pages", actualMap).Return(expectedSortedAns)
	if err != nil {
		log.Println(err)
	}
	router.ServeHTTP(response, request)
	fmt.Println(response.Code)
	if http.StatusBadGateway == response.Code {
		t.Error("bad gateway error")
	} else {
		t.Error(response.Code)
	}
}
