package craigslist

import (
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	//"github.com/walterjwhite/go-application/libraries/logging"
	"github.com/walterjwhite/go-application/libraries/chromedpexecutor"
	"github.com/walterjwhite/go-application/libraries/sleep"

	"github.com/rs/zerolog/log"
	//"os"
)

const (
	craigslistBasePostUrl = "https://post.craigslist.org/c/"
)

var (
	minimumDelayBetweenActionsFlag = flag.Int("CraigslistMinimumDelayBetweenActions", 250, "Minimum Delay between actions (ms)")
	deviationBetweenActionsFlag    = flag.Int("CraigslisDeviationBetweenActions", 5000, "Deviation between actions (ms)")
	devToolsWsUrlFlag              = flag.String("DevToolsWsUrl", "", "Dev Tools WS URL")

	//delayBetweenActions     time.Duration
	delay *sleep.RandomDelay
)

func init() {
	//var err error

	delay = &sleep.RandomDelay{MinimumDelay: *minimumDelayBetweenActionsFlag, Deviation: *deviationBetweenActionsFlag}

	//delayBetweenActions, err = time.ParseDuration(*delayBetweenActionsFlag)
	//logging.Panic(err)
}

func (p *CraigslistPost) Create(ctx context.Context) {
	log.Info().Msgf("post: %v", p)

	p.session = chromedpexecutor.New(ctx)

	p.session.Execute(chromedp.Navigate(craigslistBasePostUrl + p.Region))

	p.session.Execute(p.doForSaleBy()...)
	p.session.Execute(p.doCategory()...)

	p.session.Execute(p.doPostDetails()...)
	p.session.Execute(p.doPhone()...)
	p.session.Execute(p.doMedia()...)
	p.session.Execute(p.publish()...)
}

func (p *CraigslistPost) publish() []chromedp.Action {
	return []chromedp.Action{chromedp.Click("//*[@id=\"publish_top\"]/button")}
}

// TODO: this is generic code, unrelated to craigslist
// move this out into chromedp helper ...

func (p *CraigslistPost) HasDefault() bool {
	return false
}

func (p *CraigslistPost) Refreshable() bool {
	return false
}

func (p *CraigslistPost) EncryptedFields() []string {
	return nil
}

func Wait() {
	delay.Wait()
}
