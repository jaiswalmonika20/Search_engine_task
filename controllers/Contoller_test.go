package controllers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	mocks "github.com/module_page/mocks/services"
	"github.com/module_page/services"
)

func TestPageController_online(t *testing.T) {

	t.Run("validate activeness", func(t *testing.T) {
		response := httptest.NewRecorder()
		mockRepo := mocks.PageService{}
		context, _ := gin.CreateTestContext(response)
		mockRepo.On("online", context)
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
		// TODO: Add test cases.
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

// func TestPageController_GetAllPage(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	resp := httptest.NewRecorder()
// 	pageRepositroy := mocks.PageService{}
// 	p := []*models.Page{
// 		{
// 			ID:  1,
// 			Key: "ford",
// 		},
// 	}
// 	pageRepositroy.On("GetAllPages").Return(p, nil)
// 	userServic := New(&pageRepositroy)
// 	userServic.GetAllPage(&gin.Context{})
// 	if resp != nil {
// 		t.Error("service not working")
// 	}
// }

// func TestPageController_CreateNewPage(t *testing.T) {
// 	page := []*models.Page{
// 		{
// 			ID:  1,
// 			Key: "ford",
// 		},
// 	}
// 	var ctx *gin.Context
// 	mockRepo := mocks.PageService{}
// 	mockRepo.On("AddPage", page).Return(nil)
// 	pageService := New(&mockRepo)
// 	pageService.CreateNewPage(ctx)

// }
