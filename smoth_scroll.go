package main

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func SmoothScroll(ctx context.Context) error {
	return chromedp.Run(ctx,
		chromedp.Evaluate(`
		   window.scrollTo({
				top: document.body.scrollHeight,
				behavior: 'smooth'
			});
        `, nil),
		chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(`
			 window.scrollTo({
				top: 0,
				behavior: 'smooth'
			});
        `, nil),
	)
}
