package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"

	"scraping_challenge/app"
)

func main() {
	config, err := app.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// let's make it not headless
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:2],
		chromedp.DefaultExecAllocatorOptions[3:]...,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	getData(ctx, config)
}

func getData(ctx context.Context, config app.Config) {
	err := chromedp.Run(ctx, chromedp.Navigate("https://app.comet.co/freelancer"))
	if err != nil {
		log.Fatal(err)
		return
	}

	emailInputSel := `//input[@name="email"]`
	passInputSel := `//input[@name="password"]`
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.WaitVisible(emailInputSel),
		chromedp.WaitVisible(passInputSel),
		chromedp.SendKeys(emailInputSel, config.ProfileCredentials.Email),
		chromedp.SendKeys(passInputSel, config.ProfileCredentials.Pass),
	})
	if err != nil {
		log.Fatal(err)
	}

	// let's wait for `axeptio` widget to be visible.
	// Looks like it blocks form submit if it's not visible yet
	err = chromedp.Run(ctx, chromedp.WaitVisible(`//div[contains(@class, 'axeptio_widget')]`))
	if err != nil {
		log.Fatal(err)
	}

	// It's time to submit form
	submitButtonSel := `//button[@type="submit"]`
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Click(submitButtonSel),
	})
	if err != nil {
		log.Fatal(err)
	}

	// TODO: handle error if is present `Mot de passe incorrect.` after this step
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.WaitNotPresent(emailInputSel),
		chromedp.WaitNotPresent(passInputSel),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Let's go to profile page
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://app.comet.co/freelancer/profile"))
	if err != nil {
		log.Fatal(err)
		return
	}

	meViewSel := `//div[contains(@class, "MeView_main")]`
	freelancerDetailsSel := meViewSel + `//div[contains(@class, "FreelancerDetails_freelancerDetails")]`
	fullNameSel := freelancerDetailsSel + `//div[contains(@class, "FreelancerDetails_fullName")]`
	subtitleSel := freelancerDetailsSel + `//div[contains(@class, "FreelancerDetails_subtitle")]`

	var fullName, subtitle string
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Text(fullNameSel, &fullName, chromedp.NodeVisible),
		chromedp.Text(subtitleSel, &subtitle, chromedp.NodeVisible),
	})

	fmt.Println("Full name:", fullName)
	fmt.Println("Subtitle:", subtitle)
}
