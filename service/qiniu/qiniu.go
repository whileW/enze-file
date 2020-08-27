package qiniu

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/config"
)

func GetSetting() *config.Settings {
	return global.GVA_CONFIG.Setting.GetChild("qiniu")
}

func GetMac() *qbox.Mac {
	return qbox.NewMac(GetSetting().GetString("ak"), GetSetting().GetString("sk"))
}

func GetZone() *storage.Zone {
	switch GetSetting().GetString("zone") {
	case "huadong":
		return &storage.ZoneHuadong
	case "huabei":
		return &storage.ZoneHuabei
	case "huanan":
		return &storage.ZoneHuanan
	case "beimei":
		return &storage.ZoneBeimei
	case "xinjiapo":
		return &storage.ZoneXinjiapo
	default:
		return &storage.ZoneHuanan
	}
}