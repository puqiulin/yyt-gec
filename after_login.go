package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func AfterLogin(ctx context.Context, i int) error {
	fmt.Printf("after login action(%d)\n", i)

	//specialInvoiceUrl := "https://gec.10010.com/portal-manager/invoiceService/specialInvoiceApplyGet"
	specialInvoiceSelector := `#protal-pro > div.quickEntranceBox___1qCE_ > div > div > div:nth-child(3) > img`
	specialInvoiceSelectorMenu := `#\/invoiceService\$Menu > li:nth-child(1) > a`
	iframe := `#micro-iframe`
	querySelector := `#js-skewedRefer`
	invoiceSelector := `#loading > div.site-self-service > div > div > div > div.IntentionMain > div > div > div.zealui-table-box > div.zealui-table-body.zealui-table-main > table > tbody > tr > td:nth-child(2) > div`
	closeInvoiceSelector := `#loading > div.electronicInvoice > div.electronicInvoicePrint > i > img`
	billingSelector := `#protal-manager > div > section > section > aside > div > ul > li.ant-menu-submenu.ant-menu-submenu-inline.billingService_joyride > div`
	myBillingSelector := `#\/billingService\$Menu > li > a`
	commonUserSelector := `#protal-manager > div > section > section > section > main > div.payment___3u_9m > div.search___1lnYC.searchCard > div > div.ant-card-body > div > form > div:nth-child(1) > div:nth-child(2) > div > div > div.ant-col.ant-form-item-control-wrapper > div > span > div > button`
	//checkSelector0 := `body > div:nth-child(10) > div > div.ant-modal-wrap.ant-modal-centered > div > div.ant-modal-content > div.ant-modal-body > div.table___2hrYQ > div > div > div > div > div > div > div.ant-table-header > table > thead > tr > th.ant-table-selection-column > span > div > span.ant-table-column-title > div > label > span > input`
	//checkSelectorXpath1 := `/html/body/div[3]/div/div[2]/div/div[2]/div[2]/div[1]/div/div/div/div/div/div/div[1]/table/thead/tr/th[1]/span/div/span[1]/div/label/span/input`
	checkSelectorXpath2 := `/html/body/div[2]/div/div[2]/div/div[2]/div[2]/div[1]/div/div/div/div/div/div/div[1]/table/thead/tr/th[1]/span/div/span[1]/div/label/span/input`
	//yesCheckSelectorXpath1 := `/html/body/div[3]/div/div[2]/div/div[2]/div[2]/div[2]/button`
	yesCheckSelectorXpath2 := `/html/body/div[2]/div/div[2]/div/div[2]/div[2]/div[2]/button`
	//yesCheckSelector2 := `body > div:nth-child(10) > div > div.ant-modal-wrap.ant-modal-centered > div > div.ant-modal-content > div.ant-modal-body > div.okButton___1Iel_ > button`
	//yesCheckSelector2 := `body > div:nth-child(10) > div > div.ant-modal-wrap.ant-modal-centered > div > div.ant-modal-content > div.ant-modal-body > div.okButton___1Iel_ > button`

	queryYesCheckSelector := `#search`

	return chromedp.Run(ctx,
		//loading page
		chromedp.ActionFunc(func(ctx context.Context) error {
			fmt.Println("loading special invoice page...")
			return nil
		}),

		chromedp.ActionFunc(func(ctx context.Context) error {
			do, err := target.GetTargets().Do(ctx)
			if err != nil {
				fmt.Println(err)
				return err
			}
			if len(do) == 1 {
				return chromedp.Run(ctx,
					chromedp.Click(specialInvoiceSelector, chromedp.NodeVisible),
					chromedp.Sleep(1*time.Second),
				)
			}

			for _, info := range do {
				fmt.Printf("%#v\n", info)
				if info.URL != "https://gec.10010.com/" {
					fmt.Printf("%#v\n", info)
					targetCtx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(info.TargetID))

					return chromedp.Run(targetCtx,
						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking invoice menu...")
							return nil
						}),
						chromedp.Click(specialInvoiceSelectorMenu, chromedp.NodeVisible),
						chromedp.Sleep(4*time.Second),

						//loading iframe
						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("loading iframe...")
							return nil
						}),
						chromedp.WaitReady(iframe, chromedp.ByQuery),

						//chromedp.Evaluate(`
						//	const clickButton = ()=>{
						//	  let button = document.querySelector("#micro-iframe")?.contentDocument.querySelector("#js-skewedRefer")
						//	  if (button){
						//		  console.log(button)
						//		  setInterval(()=>button.click(),1000)
						//	  }else{
						//		  setTimeout(clickButton, 1000)
						//	  }
						//	}
						//	clickButton()
						//`, nil),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking invoice...")
							return nil
						}),
						chromedp.Evaluate(fmt.Sprintf(`
							var clickQueryButton = ()=>{
							  let button = document.querySelector("%s")?.contentDocument.querySelector("%s")
							  if (button){
								  console.log("clickQueryButton",button)
								  button.click()
							  }else{
								  setTimeout(clickQueryButton, 1000)
							  }
							}
							clickQueryButton()
						`, iframe, querySelector), nil),

						chromedp.Sleep(2*time.Second),

						chromedp.Evaluate(fmt.Sprintf(`
							var clickInvoiceButton = ()=>{
							  let button = document.querySelector("%s")?.contentDocument.querySelector("%s")
							  if (button){
								  console.log("clickInvoiceButton",button)
								  button.click()
							  }else{
								  setTimeout(clickInvoiceButton, 1000)
							  }
							}
							clickInvoiceButton()
						`, iframe, invoiceSelector), nil),

						chromedp.Sleep(3*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("closing invoice...")
							return nil
						}),
						chromedp.Evaluate(fmt.Sprintf(`
							var clickInvoiceCloseButton = ()=>{
							  let button = document.querySelector("%s")?.contentDocument.querySelector("%s")
							  if (button){
								  console.log("clickInvoiceCloseButton",button)
								  button.click()
							  }else{
								  setTimeout(clickInvoiceCloseButton, 1000)
							  }
							}
							clickInvoiceCloseButton()
						`, iframe, closeInvoiceSelector), nil),
						chromedp.Sleep(1*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							if i == 1 {
								fmt.Println("clicking billing...")
								return chromedp.Run(ctx,
									chromedp.Click(billingSelector, chromedp.NodeVisible),
									chromedp.Sleep(1*time.Second),
								)
							}
							return nil
						}),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking myBilling...")
							return nil
						}),
						chromedp.Click(myBillingSelector, chromedp.NodeVisible),
						chromedp.Sleep(5*time.Second),

						//chromedp.ActionFunc(func(ctx context.Context) error {
						//	fmt.Println("clicking common user...")
						//	var clickable bool
						//	for start := time.Now(); time.Since(start) < time.Minute; time.Sleep(500 * time.Millisecond) {
						//		err := chromedp.Evaluate(`
						//			var btn = document.querySelector('`+commonUserSelector+`');
						//			btn && !btn.disabled && btn.offsetWidth > 0 && btn.offsetHeight > 0;
						//		`, &clickable).Do(ctx)
						//		if err != nil {
						//			fmt.Println(err)
						//			return err
						//		}
						//		if clickable {
						//			return chromedp.Run(ctx,
						//				chromedp.Click(commonUserSelector, chromedp.NodeVisible),
						//				chromedp.Sleep(4*time.Second),
						//			)
						//		}
						//	}
						//	return nil
						//}),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking common user...")
							return nil
						}),
						chromedp.Click(commonUserSelector, chromedp.NodeVisible),
						chromedp.Sleep(2*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking check user...")
							return nil
						}),
						//chromedp.Click(checkSelectorXpath1, chromedp.BySearch),
						chromedp.Click(checkSelectorXpath2, chromedp.BySearch),
						chromedp.Sleep(2*time.Second),

						//chromedp.ActionFunc(func(ctx context.Context) error {
						//	var exists bool
						//	//err = chromedp.Evaluate(fmt.Sprintf(`!!document.querySelector('%s').contentDocument.querySelector("%s")`, iframe, checkSelector1), &exists).Do(ctx)
						//	err = chromedp.Evaluate(fmt.Sprintf(`!!document.querySelector('%s')`, checkSelector1), &exists).Do(ctx)
						//	if err != nil {
						//		fmt.Println(err)
						//		return err
						//	}
						//	if exists {
						//		fmt.Println("click checkSelector1...")
						//		return chromedp.Run(ctx,
						//			chromedp.WaitVisible(checkSelector1, chromedp.ByQuery),
						//			chromedp.Click(checkSelector1, chromedp.NodeVisible),
						//			chromedp.Sleep(3*time.Second),
						//		)
						//	} else {
						//		fmt.Println("click checkSelector2...")
						//		return chromedp.Run(ctx,
						//			chromedp.WaitVisible(checkSelector2, chromedp.ByQuery),
						//			chromedp.Click(checkSelector2, chromedp.NodeVisible),
						//			chromedp.Sleep(3*time.Second),
						//		)
						//	}
						//}),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking yes check user...")
							return nil
						}),
						//chromedp.Click(yesCheckSelectorXpath1, chromedp.BySearch),
						chromedp.Click(yesCheckSelectorXpath2, chromedp.BySearch),
						chromedp.Sleep(2*time.Second),

						//chromedp.ActionFunc(func(ctx context.Context) error {
						//	var exists bool
						//	//err = chromedp.Evaluate(fmt.Sprintf(`!!document.querySelector('%s').contentDocument.querySelector("%s")`, iframe, yesCheckSelector1), &exists).Do(ctx)
						//	err = chromedp.Evaluate(fmt.Sprintf(`!!document.querySelector('%s')`, yesCheckSelector1), &exists).Do(ctx)
						//	if err != nil {
						//		fmt.Println(err)
						//		return err
						//	}
						//	if exists {
						//		fmt.Println("click yesCheckSelector1...")
						//		return chromedp.Run(ctx,
						//			chromedp.WaitVisible(yesCheckSelector1, chromedp.ByQuery),
						//			chromedp.Click(yesCheckSelector1, chromedp.NodeVisible),
						//			chromedp.Sleep(2*time.Second),
						//		)
						//	} else {
						//		fmt.Println("click yesCheckSelector2...")
						//		return chromedp.Run(ctx,
						//			chromedp.WaitVisible(yesCheckSelector2, chromedp.ByQuery),
						//			chromedp.Click(yesCheckSelector2, chromedp.NodeVisible),
						//			chromedp.Sleep(2*time.Second),
						//		)
						//	}
						//}),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking yes check user search...")
							return nil
						}),
						chromedp.Click(queryYesCheckSelector, chromedp.ByQuery),
						chromedp.Sleep(3*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							return SmoothScroll(ctx)
						}),
					)
				}
			}
			return nil
		}),
	)
}
