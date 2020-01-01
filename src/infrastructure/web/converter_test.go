package web

import (
	"github.com/minakovuri/simpleVideoServer/src/app"
	"reflect"
	"testing"
)

func TestConvertDomainVideoListToWebVideoList(t *testing.T) {
	domainVideoList := []app.Video{
		{
			ID:        "id1",
			Name:      "name1",
			Duration:  0,
			Thumbnail: "thumbnail1",
		},
		{
			ID:        "id2",
			Name:      "name2",
			Duration:  0,
			Thumbnail: "thumbnail2",
		},
	}

	expectedWebVideoList := []Video{
		{
			ID:        "id1",
			Name:      "name1",
			Duration:  0,
			Thumbnail: "thumbnail1",
		},
		{
			ID:        "id2",
			Name:      "name2",
			Duration:  0,
			Thumbnail: "thumbnail2",
		},
	}

	receivedWebVideoList := convertDomainVideoListToWebVideoList(domainVideoList)

	if !reflect.DeepEqual(receivedWebVideoList, expectedWebVideoList) {
		t.Errorf("call convertDomainVideoListToWebVideoList with %v data: want %v, but got %v", domainVideoList, expectedWebVideoList, receivedWebVideoList)
	}

}

func TestConvertDomainVideoToWebVideo(t *testing.T) {
	domainVideo := app.Video{
		ID:        "id",
		Name:      "name",
		Duration:  0,
		Thumbnail: "thumbnail",
	}

	expectedWebVideo := Video{
		ID:        "id",
		Name:      "name",
		Duration:  0,
		Thumbnail: "thumbnail",
	}

	receivedWebVideo := convertDomainVideoToWebVideo(domainVideo)

	if !reflect.DeepEqual(receivedWebVideo, expectedWebVideo) {
		t.Errorf("call convertDomainVideoToWebVideo with %v data: want %v, but got %v", domainVideo, expectedWebVideo, receivedWebVideo)
	}
}
