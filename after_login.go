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
	checkSelector := `body > div:nth-child(10) > div > div.ant-modal-wrap.ant-modal-centered > div > div.ant-modal-content > div.ant-modal-body > div.table___2hrYQ > div > div > div > div > div > div > div.ant-table-body > table > tbody > tr:nth-child(1) > td.ant-table-selection-column > span > label > span > input`
	yesCheckSelector := `body > div:nth-child(9) > div > div.ant-modal-wrap.ant-modal-centered > div > div.ant-modal-content > div.ant-modal-body > div.okButton___1Iel_ > button`
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
							fmt.Println("clicking billing...")
							if i == 1 {
								return chromedp.Run(ctx,
									chromedp.Click(billingSelector, chromedp.NodeVisible),
									chromedp.Sleep(1*time.Second),
								)
							}
							return nil
						}),

						chromedp.Click(myBillingSelector, chromedp.NodeVisible),
						chromedp.Sleep(2*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking myBilling...")
							return nil
						}),
						chromedp.Click(commonUserSelector, chromedp.NodeVisible),
						chromedp.Sleep(3*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking common user...")
							return nil
						}),
						chromedp.Click(commonUserSelector, chromedp.NodeVisible),
						chromedp.Sleep(4*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking check user...")
							return nil
						}),
						chromedp.WaitVisible(checkSelector, chromedp.ByQuery),
						chromedp.Click(checkSelector, chromedp.NodeVisible),
						chromedp.Sleep(3*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking yes check user...")
							return nil
						}),
						chromedp.WaitVisible(yesCheckSelector, chromedp.ByQuery),
						chromedp.Click(yesCheckSelector, chromedp.NodeVisible),
						chromedp.Sleep(2*time.Second),

						chromedp.ActionFunc(func(ctx context.Context) error {
							fmt.Println("clicking yes check user search...")
							return nil
						}),
						chromedp.Click(queryYesCheckSelector, chromedp.ByQuery),
						chromedp.Sleep(3*time.Second),

						//chromedp.ActionFunc(func(ctx context.Context) error {
						//	return SmoothScroll(ctx)
						//}),
					)
				}
			}
			return nil
		}),
	)
}
