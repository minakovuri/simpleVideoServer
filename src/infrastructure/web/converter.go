package web

import "github.com/minakovuri/simpleVideoServer/src/app"

func convertDomainVideoListToWebVideoList(domainVideoList []app.Video) []Video {
	webVideoList := make([]Video, len(domainVideoList))

	for i, domainVideo := range domainVideoList {
		webVideoListItem := convertDomainVideoToWebVideo(domainVideo)
		webVideoList[i] = webVideoListItem
	}

	return webVideoList
}

func convertDomainVideoToWebVideo(domainVideo app.Video) Video {
	return Video{
		ID:        domainVideo.ID,
		Name:      domainVideo.Name,
		Duration:  domainVideo.Duration,
		Thumbnail: domainVideo.Thumbnail,
		URL:       domainVideo.URL,
	}
}
