package builder

// ShareItem 分享内容，给外界构造用的,不同的分享类型，需要传不同的参数
type ShareItem struct {
	Title string
	Url   string
	// 图片分享必传参数
	ImageOption string
	// 视频分享必传参数
	VideoOption string
}

func NewShareVideoItemBuilder() *ShareVideoItemBuilder {
	return &ShareVideoItemBuilder{}
}

type ShareVideoItemBuilder struct {
	Title string
	Url   string
	// 视频分享必传参数
	VideoOption string
}

func (b *ShareVideoItemBuilder) SetTitle(title string) *ShareVideoItemBuilder {
	b.Title = title
	return b
}

func (b *ShareVideoItemBuilder) SetUrl(url string) *ShareVideoItemBuilder {
	b.Url = url
	return b
}

func (b *ShareVideoItemBuilder) SetVideoOption(videoOption string) *ShareVideoItemBuilder {
	b.VideoOption = videoOption
	return b
}

func (b *ShareVideoItemBuilder) Build() *ShareItem {
	return &ShareItem{
		Title:       b.Title,
		Url:         b.Url,
		VideoOption: b.VideoOption,
	}
}

type ShareImageItemBuilder struct {
	Title string
	Url   string
	// 图片分享必传参数
	ImageOption string
}

func (b *ShareImageItemBuilder) SetTitle(title string) *ShareImageItemBuilder {
	b.Title = title
	return b
}

func (b *ShareImageItemBuilder) SetUrl(url string) *ShareImageItemBuilder {
	b.Url = url
	return b
}

func (b *ShareImageItemBuilder) SetImageOption(imageOption string) *ShareImageItemBuilder {
	b.ImageOption = imageOption
	return b
}

func (b *ShareImageItemBuilder) Build() *ShareItem {
	return &ShareItem{
		Title:       b.Title,
		Url:         b.Url,
		ImageOption: b.ImageOption,
	}
}

func NewShareImageItemBuilder() *ShareImageItemBuilder {
	return &ShareImageItemBuilder{}
}
