package builder

type ShareReq struct {
	Title string
	Url   string
	// 图片分享必传参数
	ImageOption string
	// 视频分享必传参数
	VideoOption string
}

// ShareServices 分享服务接口
type ShareServices interface {
	DoShare(req ShareReq)
}

// TencentClient 腾讯SDK
type TencentClient struct {
	appid string

	Share ShareServices
}

func NewClient(appid string) TencentClient {
	client := TencentClient{
		appid: appid,
	}
	initShareServices(&client)
	return client
}
func initShareServices(client *TencentClient) {
	client.Share = newShareServices()
	// 其他服务的初始化
}

func newShareServices() ShareServices {
	// 这里不管了，直接返回一个空的实现
	return nil
}
