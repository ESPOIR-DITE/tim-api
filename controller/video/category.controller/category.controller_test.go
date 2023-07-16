package categoryController

//
//import (
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"log"
//	"net/http"
//	"testing"
//	"tim-api/api"
//	"tim-api/domain/video.controller/category"
//)
//
//const categoryURL = api.TestBaseURL + "video.controller/category/"
//
//func TestGetRoles(t *testing.T) {
//	req, err := http.NewRequest("GET", categoryURL+"getAll", nil)
//	if err != nil {
//		t.Errorf("Error creating a new request: %v", err)
//	}
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println(err)
//	}
//	defer resp.Body.Close()
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//	fmt.Println(resp.Body)
//}
//
//func TestDeleteCategory(t *testing.T) {
//	id := ""
//	var response bool
//	resp, _ := api.Rest().Post(categoryURL + "delete/" + id)
//	if resp.IsError() {
//		assert.Failf(t, resp.Status(), "Fail")
//	}
//	err := api.JSON.Unmarshal(resp.Body(), &response)
//	if err != nil {
//		assert.Failf(t, resp.Status(), "Messy Response Fail")
//	}
//	assert.True(t, response)
//}
//
//func TestCreateCategory(t *testing.T) {
//	var response category.Category
//	resp, respErr := api.Rest().SetBody(category.Category{"", "Cartoon", "sketch"}).Post(categoryURL + "create")
//	assert.Nil(t, respErr)
//	if resp.IsError() {
//		assert.Failf(t, resp.Status(), "Fail")
//	}
//	err := api.JSON.Unmarshal(resp.Body(), &response)
//	if err != nil {
//		assert.Failf(t, resp.Status(), "Messy Response Fail")
//	}
//	assert.NotNil(t, response)
//}
