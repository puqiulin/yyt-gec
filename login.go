package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

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

		chromedp.WaitVisible(loginSelector),
		chromedp.Click(loginSelector, chromedp.NodeVisible),

		chromedp.WaitVisible(formSelector),
		chromedp.WaitVisible(usernameSelector),
		chromedp.WaitVisible(passwordSelector),

		chromedp.Click(usernameSelector, chromedp.NodeVisible),
		chromedp.SetValue(usernameSelector, account, chromedp.ByQuery),

		chromedp.Click(passwordSelector, chromedp.NodeVisible),
		chromedp.SetValue(passwordSelector, password, chromedp.ByQuery),

		chromedp.Click(protocolSelector, chromedp.NodeVisible),
		chromedp.Click(loginButtonSelector, chromedp.NodeVisible),

		chromedp.Sleep(5*time.Second),

		chromedp.Click(changePhoneNumberSelector, chromedp.NodeVisible),
		chromedp.Click(numberSelector, chromedp.NodeVisible),
		chromedp.Sleep(1*time.Second),

		chromedp.Sleep(20*time.Second),
	)
}
