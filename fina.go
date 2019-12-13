package main

import (
	"fmt"
	"net/http"
	"strings"

	"net/smtp"
	"task/lib/logger"
	"time"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
	"github.com/yeeyuntech/yeego"
	"gitlab.yeeyuntech.com/yee/easyweb"
	"gitlab.yeeyuntech.com/yee/easyweb_cms"
)

type VestBag struct {
	Id         int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"` // 主键
	VestName   string `json:"vest_name"`                            // 马甲包名称(30)
	VestBagUrl string `json:"vestbag_url"`                          // 马甲包地址
	State      string `json:"state"`                                // 状态
	CreateAt   string `json:"create_at"`
}

const (
	google             string = "play.google.com"
	comName            string = "Apps on Google Play"
	apk                string = "apk"
	CommodityStateDown string = "down" // 已下架
	limit              int64  = 1      // 查询条数
	online             string = "处于上线状态"
	offline            string = "下架了"
)

func (VestBag) TableName() string {
	return "lm_vest"
}

var (
	defaultDb *gorm.DB // 默认数据库
	redirect  string
	err       error
	title     string
	// commodityName string
	state string
	c     = colly.NewCollector()
	url   string
)

func init() {
	yeego.MustInitConfig(".", "conf")
	var (
		logCfg = logger.LogConfig{
			DebugPath: "logs/debug.log",
			ErrorPath: "logs/error.log",
		}
	)
	logger.InitLogger(logCfg)
	dbConf := easyweb_cms.DbConfig{
		UserName: yeego.Config.GetString("database.UserName"),
		Password: yeego.Config.GetString("database.Password"),
		Host:     yeego.Config.GetString("database.Host"),
		Port:     yeego.Config.GetString("database.Port"),
		DbName:   yeego.Config.GetString("database.DbName"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		easyweb.Logger.Error("数据库连接失败:%s", err.Error())
	}
	db.DB().SetMaxIdleConns(2000)
	defaultDb = db
	defaultDb.AutoMigrate(VestBag{})
}

func main() {
	d := cron.New()
	d.AddFunc("@hourly", task)
	d.Start()
	t1 := time.NewTimer(time.Minute * 1)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Minute * 1)
		}
	}
}

func task() {
	// vestBag := []string{
	// 	"https://play.google.com/store/apps/details?id=com.dana.kredit.loan.cash.credit.easy.uang.cair.happycoin",
	// 	"https://play.google.com/store/apps/details?id=com.credit.kredit.loan.dana.cash.loanrupiah",
	// 	"https://play.google.com/store/apps/details?id=com.credit.easy.kilat.cepat.pinjam.uang.uangcepat",
	// 	"https://play.google.com/store/apps/details?id=com.kredit.easy.dana.cair.loan.cash.credit.pinjamuang",
	// 	"https://play.google.com/store/apps/details?id=com.easy.cash.kredit.cair.uang.loan.credit.kilatloan",
	// 	"https://play.google.com/store/apps/details?id=com.easy.dana.cair.credit.uang.creditdana",
	// }
	vestBag := []string{
		"https://play.google.com/store/apps/details?id=com.kamiksp.rupiah&referrer=utm_source%3DCaptain",
		"https://play.google.com/store/apps/details?id=kwq.yeej.syicju&referrer=utm_source%3DelPYpAxZ",
		"https://play.google.com/store/apps/details?id=com.bunga.indah.leaon&referrer=utm_source%3DCaptainLoan2",
		"https://play.google.com/store/apps/details?id=com.bunga.indah.leaon&referrer=utm_source%3DCaptainLoan2",
		"https://play.google.com/store/apps/details?id=com.tebutebu.terran&referrer=utm_source%3DCaptainloan",
		"https://play.google.com/store/apps/details?id=cn.fundingku.saku.full.rzsc&referrer=utm_source%3Dcaptainloan",
	}
	for i := 0; i < len(vestBag); i++ {
		time.Sleep(4 * time.Minute)
		url := vestBag[i]
		if strings.Index(url, "https") == -1 {
			url = strings.Replace(url, "http", "https", 1)
		}
		fmt.Println(url)
		if strings.Contains(url, google) {
			download(url, comName)
		}
	}
}

func download(url string, commodity string) {
	nowTime := time.Now().Format("2006-01-02T 15:04:05")
	start := strings.LastIndex(url, ".")
	content := url[start+1 : len(url)]
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})
	err = c.Visit(url)
	if err != nil {
		easyweb.Logger.Error("链接%s谷歌链接访问出错:%s", url, err.Error())
		vestInfos := content + url + offline + nowTime
		SendMail(vestInfos)
	} else {
		if !strings.Contains(title, commodity) {
			vestInfos := content + url + offline + nowTime
			SendMail(vestInfos)
		} else {
			vestinfo := VestBag{
				VestName:   content,
				VestBagUrl: url,
				CreateAt:   nowTime,
				State:      online,
			}
			vestInfos := content + url + offline + nowTime
			SendMail(vestInfos)
			easyweb.Logger.Error("%s%s%s:", content, nowTime, "成功")
			if err := defaultDb.Create(&vestinfo).Error; err != nil {
				easyweb.Logger.Error("%s%s数据库添加错误%s:", content, nowTime, err)
			} else {
				weekDay := int(time.Now().Weekday())
				if weekDay == 1 {
					var vest []VestBag
					defaultDb.Model(&VestBag{}).Find(&vest)
					defaultDb.Model(&VestBag{}).Delete(&VestBag{})
					var vestarr []string
					for _, val := range vest {
						vestInfos := val.VestName + val.VestBagUrl + val.State + val.CreateAt
						vestarr = append(vestarr, vestInfos)
					}
					vestMessage := strings.Join(vestarr, "\r\n")
					SendMail(vestMessage)
				}
			}
		}
	}
}

func download2(url string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		easyweb.Logger.Error("http get error:%s", err.Error())
		return "1", err
	}
	defer resp.Body.Close() //函数结束后关闭相关链接
	if len(resp.Header["Location"]) > 0 {
		return resp.Header["Location"][0], nil
	} else {
		return "1", fmt.Errorf("该链接不是谷歌链接")
	}

}

func SendMail(body string) {
	auth := smtp.PlainAuth("", "1328162839@qq.com", "waxyjhcjpneqhgaf", "smtp.qq.com")
	to := []string{"1328162839@qq.com"}
	nickname := "韩"
	user := "1328162839@qq.com"
	subject := "马甲包状态通知"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		easyweb.Logger.Error("send mail error: %s", err.Error())
	}
}
