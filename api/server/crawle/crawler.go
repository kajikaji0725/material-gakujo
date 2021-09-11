package crawle

import (
	"log"

	"github.com/earlgray283/material-gakujo/api/db"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
)

type Crawler struct {
	controller *db.Controller
	client     *gakujo.Client
	userID     int
}

func NewCrawler(controller *db.Controller, username, password string, userID int) (*Crawler, error) {
	c := gakujo.NewClient()
	if err := c.Login(username, password); err != nil {
		return nil, err
	}
	return &Crawler{
		controller: controller,
		client:     c,
		userID:     userID,
	}, nil
}

func (c *Crawler) CrawleAll() error {
	crawleFuncs := []func() error{c.CrawleSeiseki}
	for _, crawleFunc := range crawleFuncs {
		if err := crawleFunc(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Crawler) CrawleSeiseki() error {
	kc, err := c.client.NewKyoumuClient()
	if err != nil {
		return err
	}
	seisekis, err := kc.SeisekiRows()
	if err != nil {
		return err
	}
	for _, seiseki := range seisekis {
		err := c.controller.CreateFirstSeiseki(seiseki, c.userID)
		if err != nil {
			return err
		}
	}
	log.Println("成績 crawle task completed")
	return nil
}

/*
func (c *Crawler) CrawleClassNotice() error {
	years := []int{2018, 2019, 2020, 2021}
	for _, year := range years {
		opt := gakujomodel.AllClassNoticeSearchOpt(year)
		rows, err := c.client.ClassNoticeRows(opt)
		if err != nil {
			return err
		}
		for _, row := range rows {
			time.Sleep(time.Second)
			detail, err := c.client.ClassNoticeDetail(&row, opt)
			if err != nil {
				return err
			}

		}
	}

}
*/
