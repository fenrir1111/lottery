package servicesLotto

import (
	"cp33/models"
)

func Played(platformId, intPlayId, intSubId int) *models.Played {
	u := models.Played{}
	_ = models.Db.Model(&u).Where("group_id=? and platform_id=? and sub_id=?", intPlayId, platformId, intSubId).Limit(1).Select()
	return &u
}

//func PostCheckPlayedPrize(platformId int, f64BetPrize float64) bool {
//	u := models.Played{}
//	err := models.Db.Model(&u).Where("id=? and platform_id=? and sub_id=? and enable=? and name=?", intPlayId, platformId, intSubId, true, subName).Limit(1).Select()

//	if err == nil && u.Id > 0 { //存在 成功
//		return true
//	}
//	return false

//}
