package servicesLotto

import (
	"cp33/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

//t为type彩种 o为offset issue期号
func OpenData(t, o, l, issue int) (err error, result models.Result) {
	u := models.Data{}
	if issue <= 0 {
		err = models.Db.Model(&u).Where("type=?", t).Order("issue DESC").Limit(l).Select()
	} else {
		err = models.Db.Model(&u).Where("type=? and issue=?", t, issue).Order("issue DESC").Limit(l).Offset(o).Select()
	}

	//fmt.Println(u.Id)
	if err == nil && u.Id > 0 { //存在 成功
		result = models.Result{Code: 200, Message: "ok", Data: &u}
		return
	} else if err != nil {
		result = models.Result{Code: 400, Message: err.Error(), Data: nil}
		return
	} else {
		err = errors.New("记录不存在！")
		result = models.Result{Code: 404, Message: "error!", Data: nil}
		return
	}

	return
}

func OpenInfo(t int) (err error, result models.Result) { // 上一期开奖信息及当前可购买期号
	last_period := 0
	current_period := 0
	var delaySecond int = 10 //截止投注前n秒
	err, result = OpenData(t, -1, 1, -1)
	if result.Code != 200 {
		fmt.Println("OpenInfo 42")
		return
	}

	last_period = result.Data.(*models.Data).Issue //数据库内有数据的期号为上一期数据 即有开奖数据的那期绝对不能投了
	//fmt.Println("last_period:", last_period)
	dt := models.DataTime{}
	sTime := time.Now().Add(time.Second * time.Duration(delaySecond)).Format("15:04:05") //数据库检索时间

	err = models.Db.Model(&dt).Where("type=? and action_time>?", t, sTime).Order("action_time").Limit(1).Select()
	if err == nil && dt.Type >= 0 {
		tmpCurrent_period, _ := strconv.Atoi(time.Now().Format("060102"))
		tmpCurrent_period = tmpCurrent_period * 1000
		current_period = tmpCurrent_period + dt.ActionNo
		//fmt.Println(current_period, "ww		", dt.ActionNo, "	", dt.ActionTime, "	day:", time.Now(), "	")
	} else {
		result = models.Result{Code: 590, Message: "系统错误", Data: nil}
		return
	}
	var timeleft int64
	var ttActionTime time.Time
	ttActionTime, err = time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" "+dt.ActionTime, time.Local)
	if err != nil {
		timeleft = 100000000
	}
	timeleft = ttActionTime.Unix() - time.Now().Local().Unix()
	//fmt.Println(":::time::", ttActionTime.Unix()-time.Now().Local().Unix(), "	", time.Now().Local().Unix(), "	", time.Now().Sub(ttActionTime).String())
	timeleft = timeleft - 10
	if last_period != 0 && last_period+1 >= current_period { //
		out := (models.OpenInfo{Last_period: last_period, Last_open: result.Data.(*models.Data).Data, Current_period: last_period + 1, Current_period_status: "截止", Timeleft: timeleft, Type: t})
		result = models.Result{Code: 200, Message: "ok", Data: &out}
		//fmt.Println(out)
		return
	}
	var Last_open string
	err, result = OpenData(t, 0, 1, current_period-1)
	if result.Code != 200 { //
		//fmt.Println(result.Message, current_period-1)
		Last_open = ""
	} else {
		Last_open = result.Data.(*models.Data).Data
	}
	out := (models.OpenInfo{Last_period: current_period - 1, Last_open: Last_open, Current_period: current_period, Current_period_status: "截止", Timeleft: timeleft, Type: t})
	result = models.Result{Code: 200, Message: "ok", Data: &out}
	//fmt.Println("22:", out)

	return
}
