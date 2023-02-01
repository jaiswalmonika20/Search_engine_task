package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/module_page/mocks/pkg/services"
	"github.com/module_page/pkg/models"
	"github.com/module_page/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pagecontroller *PageController

func TestPageController_online(t *testing.T) {

	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	Online(context)
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
	mockRepo := mocks.NewPageService(t)
	mockRepo.On("AddPage", mock.Anything, mock.Anything).Return(nil)
	controller := New(mockRepo)

	router := gin.Default()
	router.POST("/newpage", controller.CreateNewPage)
	//status 200
	page := models.Page{
		ID:  1,
		Key: "lamborgini",
	}
	jsonInput, _ := json.Marshal(page)
	req := httptest.NewRequest("POST", "/newpage", bytes.NewBuffer(jsonInput))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.NotEmpty(t, resp.Body)
	//status key size limit
	page = models.Page{
		ID:  1,
		Key: "ford dee dd gf d e   d d f d e   d dh s",
	}
	jsonInput, _ = json.Marshal(page)
	req = httptest.NewRequest("POST", "/newpage", bytes.NewBuffer(jsonInput))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.NotEmpty(t, resp.Body)

	//status bad_request
	// input = &models.Page{
	// 	ID:  1,
	// 	Key: "ford dee",
	// }
	// jsonInput, _ = json.Marshal(input)
	// req = httptest.NewRequest("POST", "/newpage", bytes.NewBuffer(jsonInput))
	// resp = httptest.NewRecorder()
	// router.ServeHTTP(resp, req)
	// assert.Equal(t, http.StatusBadRequest, resp.Code)
	// assert.NotEmpty(t, resp.Body)

}
func TestPageController_GetAllPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockRepo := mocks.NewPageService(t)
	pageController := New(mockRepo)
	mockRepo.On("GetAllPages").Return([]*models.Page{}, nil)
	router.GET("/pages", pageController.GetAllPage)

	req := httptest.NewRequest("GET", "/pages", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	// assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestPageController_GetByQuery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockRepo := mocks.NewPageService(t)
	pageController := New(mockRepo)
	mockRepo.On("GetAllPages").Return([]*models.Page{}, nil)
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
	actualMap := Calculate_rating(pages, queries)
	if got := Calculate_rating(pages, queries); !reflect.DeepEqual(got, expectedMap) {
		t.Errorf("Calculate_rating() = %v, want %v", got, expectedMap)
	}
	expectedSortedAns := []string{
		"P1", "P2", "P3",
	}
	if got := SortByPriority_Pages(actualMap); !reflect.DeepEqual(got, expectedSortedAns) {
		t.Errorf("Calculate_rating() = %v, want %v", got, expectedSortedAns)
	}
	router.GET("/:query", pageController.GetByQuery)

	req := httptest.NewRequest("GET", "/:query", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

}
