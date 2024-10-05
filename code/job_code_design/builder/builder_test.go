package builder

import "testing"

func SDKTest(t *testing.T) {
	// 测试SDK
	imageItem := NewShareVideoItemBuilder().SetVideoOption("videoOption").SetTitle("title").SetUrl("url").Build()
	ManagerInstance.DoShareImage(imageItem)
	videoItem := NewShareImageItemBuilder().SetImageOption("imageOption").SetTitle("title").SetUrl("url").Build()
	ManagerInstance.DoShareVideo(videoItem)
}
