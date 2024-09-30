package builder

type ShareServices interface {
	DoShare()
}

// TencentClient 腾讯SDK
type TencentClient struct {
	appid string

	Share ShareServices
}

func NewClient(appid string) TencentClient{

}
func initShareServices(client *TencentClient) {
	client.Share=
}
