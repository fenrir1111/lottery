package models

import (
	"sync"
	"time"

	_ "github.com/lib/pq"
	//"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

var App *iris.Application

const (
	PwdKey  string = "(i0dj2A#;ll01"
	AseSalt string = "#^UVBN_+~vTbz,.q"
)

var (
	MapCaptcha = make(map[string]string) //保存生成的验证码
	WsMutex    sync.Mutex
	WsConn     = make(map[websocket.Connection]bool)
)
var (
	//前三、后三 直选和值
	Sum_90_56 = map[string]int{"0": 1, "1": 3, "2": 6, "3": 10, "4": 15, "5": 21, "6": 28, "7": 36, "8": 45, "9": 55, "10": 63, "11": 69, "12": 73, "13": 75, "14": 75, "15": 73, "16": 69, "17": 63, "18": 55, "19": 45, "20": 36, "21": 28, "22": 21, "23": 15, "24": 10, "25": 6, "26": 3, "27": 1}
	//直选跨度
	Skip91_57 = map[string]int{"0": 10, "1": 54, "2": 96, "3": 126, "4": 144, "5": 150, "6": 144, "7": 126, "8": 96, "9": 54}
	//后三、前三组选和值
	Sum_97_63 = map[string]int{"1": 1, "2": 2, "3": 2, "4": 4, "5": 5, "6": 6, "7": 8, "8": 10, "9": 11, "10": 13, "11": 14, "12": 14, "13": 15, "14": 15, "15": 14, "16": 14, "17": 13, "18": 11, "19": 10, "20": 8, "21": 6, "22": 5, "23": 4, "24": 2, "25": 2, "26": 1}
	//前二直选和值
	Sum_40 = map[string]int{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5, "5": 6, "6": 7, "7": 8, "8": 9, "9": 10, "10": 9, "11": 8, "12": 7, "13": 6, "14": 5, "15": 4, "16": 3, "17": 2, "18": 1}
	//前二直选跨度
	Skip41 = map[string]int{"0": 10, "1": 18, "2": 16, "3": 14, "4": 12, "5": 10, "6": 8, "7": 6, "8": 4, "9": 2}
	//前二组选和值
	Sum48 = map[string]int{"0": 0, "1": 1, "2": 1, "3": 2, "4": 2, "5": 3, "6": 3, "7": 4, "8": 4, "9": 5, "10": 4, "11": 4, "12": 3, "13": 3, "14": 2, "15": 2, "16": 1, "17": 1, "18": 0}
	//任二直选和值
	Sum124 = map[string]int{"0": 1, "1": 2, "2": 3, "3": 4, "4": 5, "5": 6, "6": 7, "7": 8, "8": 9, "9": 10, "10": 9, "11": 8, "12": 7, "13": 6, "14": 5, "15": 4, "16": 3, "17": 2, "18": 1}
	//任二组选和值
	Sum127 = map[string]int{"0": 0, "1": 1, "2": 1, "3": 2, "4": 2, "5": 3, "6": 3, "7": 4, "8": 4, "9": 5, "10": 4, "11": 4, "12": 3, "13": 3, "14": 2, "15": 2, "16": 1, "17": 1, "18": 0}
	//任三直选复式
	CombArr128 = map[int]interface{}{0: []int{0, 1, 2}, 1: []int{0, 1, 3}, 2: []int{0, 1, 4}, 3: []int{0, 2, 3}, 4: []int{0, 2, 4}, 5: []int{0, 3, 4}, 6: []int{1, 2, 3}, 7: []int{1, 2, 4}, 8: []int{1, 3, 4}, 9: []int{2, 3, 4}}
	//任三直选和值
	Sum130 = map[string]int{"0": 1, "1": 3, "2": 6, "3": 10, "4": 15, "5": 21, "6": 28, "7": 36, "8": 45, "9": 55, "10": 63, "11": 69, "12": 73, "13": 75, "14": 75, "15": 73, "16": 69, "17": 63, "18": 55, "19": 45, "20": 36, "21": 28, "22": 21, "23": 15, "24": 10, "25": 6, "26": 3, "27": 1}
	//任三组选和值
	Sum137 = map[string]int{"1": 1, "2": 2, "3": 2, "4": 4, "5": 5, "6": 6, "7": 8, "8": 10, "9": 11, "10": 13, "11": 14, "12": 14, "13": 15, "14": 15, "15": 14, "16": 14, "17": 13, "18": 11, "19": 10, "20": 8, "21": 6, "22": 5, "23": 4, "24": 2, "25": 2, "26": 1}
	//[0, 1, 2, 3]
	CombArr139 = map[int]interface{}{0: []int{0, 1, 2, 3}, 1: []int{0, 1, 2, 4}, 2: []int{0, 1, 3, 4}, 3: []int{0, 2, 3, 4}, 4: []int{1, 2, 3, 4}}
)

type Result struct { //全站通用json返回结果
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WebsiteInfo struct { //网站基础
	Title string
	Host  string
}

type LoginPost struct { //post登陆
	Platform string `form:"platform"` //平台uuid
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginCookie struct {
	Platform       string `form:"platform"` //平台uuid
	Username       string `form:"username"`
	Enclientpasswd string `form:"enclientpasswd"`
}

type Pingtai struct { //注册平台相关
	Id       int    `form:"id"`
	Platform string `form:"platform"` //平台uuid
	Ctime    string `form:"ctime"`
	Stime    string `form:"Stime"`
	Etime    string `form:"etime"`
	Qq       string `form:"qq"`
}

type AjaxBetList struct {
	OrderType int `form:"orderType"` //0 全部，1追号 2中奖  3待开奖 4撤单
	PageIndex int `form:"pageIndex"`
}

type SignupPost struct { //post提交的注册信息
	Platform string `form:"platform"` //平台uuid
	Username string `form:"username"` //用户名
	Password string `form:"password"` //密码
	Captcha  string `form:"captcha"`  //验证码
	Uuid     string `form:"uuid"`     //uuid
}

type Trend struct { //走势图post请求数据
	Gid   int `form:"gid"`   //对应数据库data->type
	Count int `form:"count"` //要取的数据条目
	Pos   int `form:"pos"`   //位数，1万 2千 3百 4十 5个
}

type DataTime struct {
	Id         int
	Type       int
	ActionNo   int
	ActionTime string
	StopTime   string
}

type OpenInfo struct { //输出当前期号、开奖信息等
	Type                  int    `json:"type"`
	Last_period           int    `json:"last_period"` //上期期号
	Last_open             string `json:"last_open"`   //上期开奖号码
	Current_period        int    `json:"current_period"`
	Current_period_status string `json:"current_period_status"`
	Timeleft              int64  `json:"timeleft"`
}

type Data struct { //数据库
	Id    int
	Type  int
	Time  time.Time
	Data  string
	Issue int
}

type Played struct { //数据库
	Id            int
	SubName       string
	Enable        bool
	SubId         int
	BonusProp     string
	BonusPropBase string
	GroupId       int
	SimpleInfo    string
	Info          string
	Example       string
	Sort          int
	MinCharge     float64
	MaxCount      int
	PlatformId    int
}

type Lottery struct { //数据库
	Id int
	//PlatformId   int
	Type         int //1为时时彩
	Enable       bool
	IsDelete     bool
	sort         int
	Name         string
	ShortName    string
	Delay_second int
	Count        int //每次投注最大注数
}

type PostBet struct { //投注单
	GameId     int                       `form:"game_id"`
	GamePeriod int                       `form:"game_period"`
	BetNext    int                       `form:"bet_next"`
	Amount     float64                   `form:"amount"`
	BetMore    int                       `form:"bet_more"`
	BetWinStop int                       `form:"bet_win_stop"`
	Bet_list   map[int]map[string]string `form:"bet_list"`
	//	Bet_list   map[int]interface{} `formam:"bet_list"`
}

type Bets struct { //对应数据库bets
	Id         int64
	PlatformId int
	Uid        int
	GameId     int
	GamePeriod int
	BetNext    int
	Amount     float64
	BetMore    int
	BetWinStop int
	Ctime      string
	Etime      string
	IsWin      bool
	WinAmount  float64
	OpenNum    string
	Status     int
	//Bet_list   *Bet_list

	PlayId       int     `form:"playId";pg:"playId"` //对应数据库played表 group_id ,played_group表id
	SubId        int     `form:"subId"`
	SubName      string  `form:"subName"`
	BetCode      string  `form:"betCode"`
	BetCount     int     `form:"betCount"`
	BetMoney     float64 `form:"betMoney"`
	BetEachMoney float64 `form:"betEachMoney"`
	BetPrize     string  `form:"betPrize"`
	BetPrizeShow string  `form:"betPrizeShow"`
	BetReward    float64 `form:"betReward"`
	BetPos       string  `form:"betPos"`
}

type Bet_list struct {
	PlayId       int     `form:"playId";pg:"playId"`
	SubId        int     `form:"subId"`
	SubName      string  `form:"subName"`
	BetCode      string  `form:"betCode"`
	BetCount     int     `form:"betCount"`
	BetMoney     float64 `form:"betMoney"`
	BetEachMoney float64 `form:"betEachMoney"`
	BetPrize     float64 `form:"betPrize"`
	BetPrizeShow float64 `form:"betPrizeShow"`
	BetReward    int     `form:"betReward"`
	BetPos       string  `form:"betPos"` //0|1|2|3|4 万、千、百、十、个
}

type Members struct {
	PlatformId     int     `json:"platform_id"`
	Uid            int     `json:"uid"`      //1
	Username       string  `json:"username"` //tfyghb
	Password       string  `json:"password"`
	Uuid           string  `json:"uuid"`
	Source         int     `json:"source"`
	IsDelete       bool    `json:"isDelete"`
	Enable         bool    `json:"enable"`
	ParentId       int     `json:"parentId"`
	Parents        []int   `json:"parents"`
	CoinPassword   string  `json:"coinPassword"`
	Type           int     `json:"type"`
	RegIp          string  `json:"regIP"`
	RegTime        string  `json:"regTime"`
	UpdateTime     string  `json:"updateTime"`
	Grade          int     `json:"grade"`      //等级
	Score          int     `json:"score"`      //积分
	ScoreTotal     int     `json:"scoreTotal"` //累计积分
	Coin           float64 `json:"coin"`
	Fcoin          float64 `json:"fcoin"`
	FanDian        float64 `json:"fanDian"`
	FanDianBdw     float64 `json:"fanDianBdw"` //不定位返点
	Qq             string  `json:"qq"`
	ConCommStatus  bool    `json:"conCommStatus"`
	LossCommStatus bool    `json:"lossCommStatus"`
	Info           string  `json:"info"`
}

type SessionRedis interface { //session
	Load(string) interface{}
	Update(string, map[string]interface{})
}

type BaseInfo struct { //基础信息
	Platform string `form:"platform";json:"platform"`
	Uuid     string `form:"uuid";json:"uuid"`
}

type Dept struct {
	Name string
	Data interface{}
}

var Ss map[string]interface{}

func (self Dept) Update() {
	Ss = make(map[string]interface{})
	Ss[self.Name] = self.Data
}

func (self Dept) Load() interface{} {
	//	Ss = make(map[string]interface{})
	//	Ss[self.Name] = self.Data
	return Ss[self.Name]
}
