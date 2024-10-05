package builder

// ShareItem 分享内容，给外界构造用的
type ShareItem struct {
	Title string
	Url   string
	// 图片分享必传参数
	ImageOption string
	// 视频分享必传参数
	VideoOption string
}
