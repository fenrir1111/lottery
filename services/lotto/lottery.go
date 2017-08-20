package servicesLotto

import (
	"cp33/models"
	//	"fmt"
)

func LotteryIn(id int) bool {
	u := models.Lottery{}
	err := models.Db.Model(&u).Where("id=? and enable=? and is_delete=?", id, true, false).Limit(1).Select()

	if err == nil && u.Id > 0 { //存在 成功
		return true
	}

	return false
}
