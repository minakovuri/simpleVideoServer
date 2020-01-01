package app

type Service interface {
	GetVideoList() []Video
	GetVideoByID(string) Video
}

type ServiceImpl struct{}

func (s ServiceImpl) GetVideoList() []Video {
	return []Video{{
		ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
	}}
}

func (s ServiceImpl) GetVideoByID(ID string) Video {
	return Video{
		ID:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		URL:       "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	}
}
