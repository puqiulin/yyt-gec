package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var (
	err error
)

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	userDataDir := filepath.Join(userHomeDir, "yyt-gec")

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("no-first-run", true),
		chromedp.Flag("no-default-browser-check", true),
		chromedp.Flag("ignore-certificate-errors", true),
		//chromedp.Flag("start-maximized", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
		chromedp.UserDataDir(userDataDir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*60*24*time.Second)
	defer cancel()

	gecUrl := "https://gec.10010.com/"

	userHeader := `#protal-pro > div.headBody___2_BvY.headBodyScroll___1bs59 > div > div > div > div.ant-col.ant-col-sm-2.ant-col-lg-2.ant-col-xl-2.ant-col-xxl-2 > div.ant-dropdown-trigger.internationalLogin___2OJgA > div`
	isLoginSelector := `#protal-pro > div.headBody___2_BvY.headBodyScroll___1bs59 > div > div > div > div.ant-col.headUserLogin___5ahDV.ant-col-sm-4.ant-col-lg-4.ant-col-xl-3.ant-col-xxl-3 > button`

	err = chromedp.Run(ctx,
		network.Enable(),
		chromedp.Evaluate(`
	       // Overwrite the 'navigator.webdriver' property
	       Object.defineProperty(navigator, 'webdriver', {
	           get: () => undefined
	
	       });
	   `, nil),
		chromedp.Navigate(gecUrl),
		chromedp.WaitReady(userHeader, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),

		chromedp.ActionFunc(func(ctx context.Context) error {
			var exists bool
			err = chromedp.Evaluate(fmt.Sprintf(`!!document.querySelector('%s')`, isLoginSelector), &exists).Do(ctx)
			if err != nil {
				return err
			}

			//if exists {
			//	return AfterLogin(ctx)
			//} else {
			//	return Login(ctx)
			//}
			if exists {
				return nil
			} else {
				return Login(ctx)
			}
		}),
	)
	if err != nil {
		fmt.Println(err)
	}
	for i := range math.MaxInt {
		err := AfterLogin(ctx, i)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
