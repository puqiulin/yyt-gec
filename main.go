package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	isLoginSelector := `#protal-pro > div.headBody___2_BvY.headBodyScroll___1bs59 > div > div > div > div.ant-col.headUserLogin___5ahDV.ant-col-sm-4.ant-col-lg-4.ant-col-xl-3.ant-col-xxl-3 > button > span`
	var loginText string

	err = chromedp.Run(ctx,
		network.Enable(),
		chromedp.Evaluate(`
	       // Overwrite the 'navigator.webdriver' property
	       Object.defineProperty(navigator, 'webdriver', {
	           get: () => undefined
	
	       });
	   `, nil),
		chromedp.Navigate(gecUrl),
		chromedp.Evaluate(`
	       // Overwrite the 'navigator.webdriver' property
	       Object.defineProperty(navigator, 'webdriver', {
	           get: () => undefined
	
	       });
	   `, nil),

		chromedp.WaitVisible(isLoginSelector),
		chromedp.Text(isLoginSelector, &loginText),

		chromedp.ActionFunc(func(ctx context.Context) error {
			if strings.Contains(loginText, "明珠") {
				fmt.Println(loginText)
				return AfterLogin(ctx)
			} else {
				return Login(ctx)
			}
		}),
	)
	if err != nil {
		fmt.Println(err)
	}

	for {
		time.Sleep(1)
	}
}

func Login(ctx context.Context) error {
	fmt.Println("login action")
	//domain := "gec.10010.com"
	//path := "/"
	gecUrl := "https://gec.10010.com/"

	//phoneNumber := "18608272283"
	account := "minzhufandian"
	password := "y98KoDR5"
	loginSelector := `#protal-pro > div.headBody___2_BvY.headBodyScroll___1bs59 > div > div > div > div.ant-col.deletePadding.ant-col-sm-4.ant-col-lg-4.ant-col-xl-3.ant-col-xxl-3 > div > div > div.loginBtn___1B74P`
	formSelector := `#root > div.ant-spin-nested-loading > div > div > div > div.containerRight___34grG > div.content___1g3x4`
	usernameSelector := `#basic_username`
	passwordSelector := `#basic_password`
	protocolSelector := `#basic > div.agreementBody___2Kgl6 > label > span.ant-checkbox > input`
	loginButtonSelector := `#basic > div:nth-child(5) > div > div > div > div > button`
	changePhoneNumberSelector := `#root > div.ant-spin-nested-loading > div > div > div > div.containerRight___34grG > div.doubleRight___1aO1m > div > div:nth-child(2) > span`
	numberSelector := `#basic_phone`

	return chromedp.Run(ctx,
		//network.SetCookies([]*network.CookieParam{
		//	{
		//		Name:   "m45av2WWBJ5jO",
		//		Value:  `57.oGs3mWgZrtQd7895uj.xkjzLnAhDAwUEy6D9bQiI2TSWFEqMpKVSS8RSDGTX.hiBKrK68SQYyL6_FPw6OpQA`,
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "SECKEY_ABVK",
		//		Value:  `7sqmc+XpHqb+Vtwn9gFaTXS1B1pfxco/VzN01sKi5lM%3D`,
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "deviceId",
		//		Value:  "15438355f6bc7cc4f0f82d26a5b378d9",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "publicKey",
		//		Value:  "uIIRy9JtTb2wu92v",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "BMAP_SECKEY",
		//		Value:  `5CmkWtiePgmrCTEFgip3vsUx-RXTvWLrFhMtvynyQ23UADA9bs84lDMrg2dzLJOJOO3ckQZdNNihPH1tiT414h5wamfWy6qwgtc-LZfqwe4EtLyhtqIIHwTMEBD1xKE6b8aKf7l3MG0DN1Mlq8JcOomRfEQnE-OpA58rA9i5hrfIyB9QuTfhvZxKNSbBPcKF`,
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "track-params",
		//		Value:  "",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "AUTH-UNIFIED-SESSION",
		//		Value:  "S-019B3F6A13BB428FA14B077DD0FA7B09-1267069171190914048",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "ipCity",
		//		Value:  `0rs^%^2FS64oT^%^2FODtGS8g0YIGA^%^3D^%^3D`,
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "ipProvince",
		//		Value:  "qc4nUzIcdlKeSJgrVIs2ow^%^3D^%^3D",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "province",
		//		Value:  "81",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "sessionId",
		//		Value:  "52421fce05b24d4ea455defa78a",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "GEC-UNIFIED-SESSION",
		//		Value:  "S-3C4D88A55FBA46E18445459764406ED8-1267069174370508800",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "JSESSIONID",
		//		Value:  "B5120694D1588D08EC1E265A93BAF3DC",
		//		Domain: domain,
		//		Path:   path,
		//	},
		//	{
		//		Name:   "m45av2WWBJ5jP",
		//		Value:  `ZOLZn1NSW3MiLOAzaTm8sjMT.YLsTCLgwvimK0OZptvfbvz1MqU2w5yHReegGuWlx7kGbeIl6nnCcFV3AHe6UthGTA85eT9qfUumFZ6ZFiwwUe.uBDwGpHHM706WaanTnexjz3WbkPw9kO9vbGdY3UU4dJYLYCrWij8MEX1IBRDyhQQ1SdANSTlNyc8.B1kBQuAI3M03s3_zrilOqdKaEZs2n9qALhAmaP_sKPhL0l566Lm0JPZelpV3AYFWBLl3IjLqaxyF9On_W7YCm2ySIg9LtH3bAEazzLBE86mFRaOrNXFcWpnDgpryaNFoVgkAEdj97jNKS1ra.4ZTaxGiJzFHeQhgFRt76ogXPlJmoHCz0sCdx9WbhVU8spsPZhMhQujjbudaCAIhRRHve6vKkq0di1SP8JaVlY79hUkfRTq^`,
		//		Domain: domain,
		//		Path:   path,
		//	},
		//}),
		chromedp.Navigate(gecUrl),
		chromedp.Sleep(1*time.Second),

		chromedp.WaitVisible(loginSelector),
		chromedp.Click(loginSelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.WaitVisible(formSelector),
		chromedp.WaitVisible(usernameSelector),
		chromedp.WaitVisible(passwordSelector),
		chromedp.Sleep(1*time.Second),

		chromedp.Click(usernameSelector, chromedp.NodeVisible),
		chromedp.SendKeys(usernameSelector, account, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),

		chromedp.Click(passwordSelector, chromedp.NodeVisible),
		chromedp.SendKeys(passwordSelector, password, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),

		chromedp.Click(protocolSelector, chromedp.NodeVisible),
		chromedp.Click(loginButtonSelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.Sleep(5*time.Second),

		chromedp.Click(changePhoneNumberSelector, chromedp.NodeVisible),
		chromedp.Click(numberSelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		//chromedp.Click(verifiedCodeSelector, chromedp.NodeVisible),

		chromedp.Sleep(20*time.Second),

		//chromedp.Click(yesSelector, chromedp.NodeVisible),
	)
}

func AfterLogin(ctx context.Context) error {
	fmt.Println("after login action")

	//specialInvoiceUrl := "https://gec.10010.com/portal-manager/invoiceService/specialInvoiceApplyGet"

	specialInvoiceSelector := `#protal-pro > div.quickEntranceBox___1qCE_ > div > div > div:nth-child(3) > img`

	specialInvoieSelectorMenu := `#\/invoiceService\$Menu > li.ant-menu-item.ant-menu-item-selected > a`

	querySelector := `#js-skewedRefer`
	invoiceSelector := `#loading > div.site-self-service > div > div > div > div.IntentionMain > div > div > div.zealui-table-box > div.zealui-table-body.zealui-table-main > table > tbody > tr > td:nth-child(2) > div`
	billingSelector := `#protal-manager > div > section > section > aside > div > ul > li.ant-menu-submenu.ant-menu-submenu-inline.billingService_joyride > div`
	myBillingSelector := `#\/billingService\$Menu > li > a`

	commonUserSelector := `#protal-manager > div > section > section > section > main > div.payment___3u_9m > div.search___1lnYC.searchCard > div > div.ant-card-body > div > form > div:nth-child(1) > div:nth-child(2) > div > div > div.ant-col.ant-form-item-control-wrapper > div > span > div > button`
	checkSelector := `body > div:nth-child(9) > div > div.ant-modal-wrap.ant-modal-centered > div > div.ant-modal-content > div.ant-modal-body > div.table___2hrYQ > div > div > div > div > div > div > div.ant-table-body > table > tbody > tr:nth-child(1) > td.ant-table-selection-column > span > label > span > input`
	yesCheckSelector := `body > div:nth-child(9) > div > div.ant-modal-wrap.ant-modal-centered > div > div.ant-modal-content > div.ant-modal-body > div.okButton___1Iel_ > button`
	queryYesCheckSelector := `#search`

	return chromedp.Run(ctx,
		//chromedp.Navigate(specialInvoiceUrl),

		chromedp.WaitVisible(specialInvoiceSelector, chromedp.ByQuery),
		chromedp.Click(specialInvoiceSelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.WaitVisible(specialInvoieSelectorMenu, chromedp.ByQuery),
		chromedp.Click(specialInvoieSelectorMenu, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.WaitVisible(querySelector, chromedp.ByQuery),
		chromedp.Click(querySelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.WaitVisible(invoiceSelector, chromedp.ByQuery),
		chromedp.Click(invoiceSelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
		chromedp.Sleep(1*time.Second),
		chromedp.Evaluate(`
            function smoothScrollToTop() {
                const currentPosition = window.pageYOffset;
                if (currentPosition > 0) {
                    window.scrollTo(0, currentPosition - 50);
                    setTimeout(smoothScrollToTop, 20);
                }
            }
            smoothScrollToTop();
        `, nil),
		chromedp.Sleep(1*time.Second),

		chromedp.WaitVisible(billingSelector, chromedp.ByQuery),
		chromedp.Click(billingSelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.WaitVisible(myBillingSelector, chromedp.ByQuery),
		chromedp.Click(myBillingSelector, chromedp.NodeVisible),
		chromedp.Sleep(4*time.Second),

		chromedp.WaitVisible(commonUserSelector, chromedp.ByQuery),
		chromedp.Click(commonUserSelector, chromedp.NodeVisible),
		chromedp.Sleep(2*time.Second),

		chromedp.WaitVisible(checkSelector, chromedp.ByQuery),
		chromedp.Click(checkSelector, chromedp.NodeVisible),
		chromedp.Sleep(2*time.Second),

		chromedp.WaitVisible(yesCheckSelector, chromedp.ByQuery),
		chromedp.Click(yesCheckSelector, chromedp.NodeVisible),
		chromedp.Sleep(2*time.Second),

		chromedp.WaitVisible(queryYesCheckSelector, chromedp.ByQuery),
		chromedp.Click(queryYesCheckSelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
		chromedp.Evaluate(`
            function smoothScrollToTop() {
                const currentPosition = window.pageYOffset;
                if (currentPosition > 0) {
                    window.scrollTo(0, currentPosition - 50);
                    setTimeout(smoothScrollToTop, 20);
                }
            }
            smoothScrollToTop();
        `, nil),
		chromedp.Sleep(1*time.Second),
	)
}
