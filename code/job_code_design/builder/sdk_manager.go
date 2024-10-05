package builder

var ManagerInstance *Manager

// Manager 旧的Manager模块，负责管理SDK的初始化，旧的代码无法大批量修改，所以需要保留
type Manager struct {
	// 旧的SDK
	SDK *TencentClient
}

func init() {
	// 初始化Manager
	ManagerInstance = &Manager{
		SDK: NewClient("123"),
	}
}

func (m *Manager) DoShareVideo(req *ShareItem) error {
	// 转换参数
	transformReq := ShareReq{
		Title:       req.Title,
		Url:         req.Url,
		VideoOption: req.VideoOption,
	}
	return m.SDK.Share.DoShare(transformReq)
}

func (m *Manager) DoShareImage(req *ShareItem) error {
	// 转换参数
	transformReq := ShareReq{
		Title:       req.Title,
		Url:         req.Url,
		ImageOption: req.ImageOption,
	}
	return m.SDK.Share.DoShare(transformReq)
}
