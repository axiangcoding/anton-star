package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"net/url"
)

const (
	StatusQueryFailed = 1
	StatusNotFound    = 2
	StatusFound       = 3
)

func GetProfileFromWTOfficial(nick string, callback func(status int, user *table.GameUser)) error {
	urlTemplate := "https://warthunder.com/zh/community/userinfo/?nick=%s"
	queryUrl := fmt.Sprintf(urlTemplate, url.QueryEscape(nick))

	c := colly.NewCollector(
		colly.AllowedDomains("warthunder.com"),
		colly.MaxDepth(1),
		colly.IgnoreRobotsTxt(),
	)
	extensions.RandomUserAgent(c)

	c.OnHTML("div[class=user__unavailable-title]", func(e *colly.HTMLElement) {
		logging.Warnf("%s userinfo not found", nick)
		callback(StatusNotFound, nil)
	})

	c.OnHTML("div[class=user-info]", func(e *colly.HTMLElement) {
		data := ExtractGaijinData(e)
		callback(StatusFound, &data)
	})

	c.OnRequest(func(r *colly.Request) {
		logging.Infof("start visiting %s", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		logging.Warnf("visiting %s failed. %s", r.Request.URL.String(), err)
		callback(StatusQueryFailed, nil)

	})

	err := c.Post(queryUrl, nil)
	if err != nil {
		logging.Warn("colly post failed. ", err)
		callback(StatusQueryFailed, nil)
		return err
	}
	return nil
}

func GetProfileFromThunderskill(nick string, callback func(status int, skill *ThunderSkillResp)) error {
	urlTemplate := "https://thunderskill.com/en/stat/%s/export/json"
	url := fmt.Sprintf(urlTemplate, nick)

	c := colly.NewCollector(
		colly.AllowedDomains("thunderskill.com"),
		colly.MaxDepth(1),
		colly.IgnoreRobotsTxt(),
	)

	c.OnResponse(func(r *colly.Response) {
		var resp ThunderSkillResp
		_ = json.Unmarshal(r.Body, &resp)
		callback(StatusFound, &resp)
	})

	c.OnRequest(func(r *colly.Request) {
		logging.Infof("start visiting %s", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		logging.Warnf("visiting %s failed. %s", r.Request.URL.String(), err)
	})

	err := c.Visit(url)
	if err != nil {
		logging.Warn("colly get failed. ", err)
		return err
	}
	return nil
}