package servicesLotto

import (
	"cp33/models"
	"cp33/services/user"
	"fmt"
	"strconv"
)

func DoBets(us []models.Bets, uid int) (result models.Result) {
	tx, err := models.Db.Begin()
	if err != nil {
		result = models.Result{Code: 500, Message: err.Error(), Data: nil}
		return
	}

	m := models.Members{}

	_, err = tx.QueryOne(&m, fmt.Sprintf("select coin from members where uid=%v limit 1 for update", uid))
	//err = tx.Model(&m).Where("uid=?", uid).Returning("coin").Select()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		result = models.Result{Code: 601, Message: "数据库错误!", Data: nil}
		return
	}

	coin := m.Coin - us[0].Amount*float64(us[0].BetMore)
	if coin < 0 {
		tx.Rollback()
		result = models.Result{Code: 601, Message: "余额不足本次消费!", Data: nil}
		return
	}
	_, err = tx.Model(&m).Set("coin=?", strconv.FormatFloat(coin, 'f', 3, 64)).Where("uid=?", uid).Update()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		result = models.Result{Code: 601, Message: "数据库错误!", Data: nil}
		return
	}

	err = tx.Insert(&us)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		result = models.Result{Code: 600, Message: "数据库错误!", Data: nil}
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		result = models.Result{Code: 600, Message: "数据库错误!", Data: nil}
		return
	}

	result = models.Result{Code: 200, Message: "ok", Data: nil}
	return
}

func BetList(bl *models.AjaxBetList, platform, username string) (result models.Result) {
	//strSql := fmt.Sprintf("select id,amount,bet_count,bet_prize,bet_reward,ctime,is_win,sub_name,bet_code,play_id,game_id,win_amount,game_period,open_num,status,bet_pos,etime from bets where uid=%v ", services.GetUid(platform, username))
	strSql := fmt.Sprintf("uid=%v", services.GetUid(platform, username))
	switch bl.OrderType {
	case 0: //0 全部
		break
	case 1: //1追号
		strSql = fmt.Sprintf(" %s%s", strSql, " and bet_next<>0")
	case 2: //中奖
		strSql = fmt.Sprintf(" %s%s", strSql, " and is_win=true")
	case 3: //3待开奖
		strSql = fmt.Sprintf(" %s%s", strSql, " and open_num<>null")
	case 4: //4撤单
		strSql = fmt.Sprintf(" %s%s", strSql, " and status=2")
	}

	var us []models.Bets
	total, err := models.Db.Model(&us).Where(strSql).Count()
	if err != nil {
		result = models.Result{Code: 500, Message: err.Error(), Data: nil}
		return
	} else if total == 0 {
		result = models.Result{Code: 200, Message: "没有数据！", Data: nil}
		return
	}

	err = models.Db.Model(&us).Where(strSql).Limit(20).Offset((bl.PageIndex - 1) * 20).Order("id DESC").Select()
	if err != nil {
		result = models.Result{Code: 500, Message: err.Error(), Data: nil}
		return
	}
	out := map[string]interface{}{"PageCount": total / 20, "Records": &us}
	result = models.Result{Code: 200, Message: "ok", Data: &out}
	return
}

func EndLottery(gameId, issue int) {
	//	bets := models.Bets{}
	//	_, err = tx.QueryOne(&m, fmt.Sprintf("select coin from members where game_id=%v and game_period=%v for update", gameId,issue))
	//	//err = tx.Model(&m).Where("uid=?", uid).Returning("coin").Select()
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		tx.Rollback()
	//		result = models.Result{Code: 601, Message: "数据库错误!", Data: nil}
	//		return
	//	}
}
