package web

import (
	"encoding/json"
	"github.com/minakovuri/simpleVideoServer/src/app"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct{}

func (ms MockService) GetVideoList() []app.Video {
	return []app.Video{{
		ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
	}}
}

func (ms MockService) GetVideoByID(string) app.Video {
	return app.Video{
		ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		URL:       "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	}
}

func TestGetVideoList(t *testing.T) {
	w := httptest.NewRecorder()

	mockService := MockService{}
	httpHandler := HttpHandlerImpl{mockService}

	httpHandler.getVideoList(w, nil)
	response := w.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	items := make([]Video, 10)

	if err = json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
}

func TestGetVideoByID(t *testing.T) {
	// TODO
}
