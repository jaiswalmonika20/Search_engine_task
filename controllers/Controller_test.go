package controllers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/module_page/mocks/services"
	"github.com/module_page/models"
	"github.com/module_page/services"
)

var pagecontroller *PageController

func TestPageController_online(t *testing.T) {

	t.Run("validate activeness", func(t *testing.T) {
		response := httptest.NewRecorder()
		//mockRepo := mocks.PageService{}
		context, _ := gin.CreateTestContext(response)
		online(context)
		if response.Code != http.StatusOK {
			t.Errorf("api not active")
		}
	})

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
			name: "testing New Constructor",
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

func TestPageController_CreateNewPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
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
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

}
func TestPageController_GetAllPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockRepo := mocks.NewPageService(t)
	mockRepo.On("GetAllPages").Return([]models.Page{}, nil)
	router.GET("/pages", pagecontroller.GetAllPage)
	req, err := http.NewRequest("GET", "/pages", nil)
	if err != nil {
		log.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
}

func TestPageController_GetByQuery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockRepo := mocks.NewPageService(t)
	mockRepo.On("GetAllPages").Return([]models.Page{}, nil)
	router.GET("/:query", pagecontroller.GetByQuery)
	req, err := http.NewRequest("GET", "/:query", nil)
	if err != nil {
		log.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
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
		// TODO: Add test cases.
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
