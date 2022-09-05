package main

import (
	"context"
	"log"
	"sync"

	"github.com/chromedp/cdproto/network"
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

	done := make(chan struct{})
	requestWillBeSentCh := make(chan *network.EventRequestWillBeSent)
	loadingFinishedCh := make(chan *network.EventLoadingFinished)
	chromedp.ListenTarget(ctx, func(v interface{}) {
		switch ev := v.(type) {
		case *network.EventRequestWillBeSent:
			if ev.Request.URL == "https://app.comet.co/api/graphql" {
				go func() {
					requestWillBeSentCh <- ev
				}()
			}
		case *network.EventLoadingFinished:
			go func() {

				loadingFinishedCh <- ev
			}()
		}
	})

	// Let's go to profile page
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://app.comet.co/freelancer/profile"))
	if err != nil {
		log.Fatal(err)
		return
	}

	loadedReqIDch := finishedRequestsListener(ctx, requestWillBeSentCh, loadingFinishedCh)

	go func() {
		for {
			select {
			case reqID, ok := <-loadedReqIDch:
				if !ok {
					break
				}

				var respBody []byte
				err = chromedp.Run(ctx, chromedp.ActionFunc(func(cxt context.Context) error {
					respBody, err = network.GetResponseBody(reqID).Do(cxt)
					return err
				}))
				if err == nil {
					respBodyStr := string(respBody)
					if respBodyStr != "" {
					}
				}

			}

		}
	}()

	<-done
	//
	// freelancerProfile, err := factory.NewFreelancerProfile(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println("profile:", freelancerProfile)
}

type requestListener struct {
	requestWillBeSentCh <-chan network.RequestID
	loadingFinishedCh   <-chan network.RequestID
	reqMap              map[network.RequestID][]byte
}

func finishedRequestsListener(ctx context.Context, requestWillBeSentCh <-chan *network.EventRequestWillBeSent, loadingFinishedCh <-chan *network.EventLoadingFinished) <-chan network.RequestID {
	ch := make(chan network.RequestID)
	var mu sync.Mutex
	requests := make(map[network.RequestID]struct{})
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case t := <-requestWillBeSentCh:
				mu.Lock()
				requests[t.RequestID] = struct{}{}
				mu.Unlock()

			case t := <-loadingFinishedCh:
				mu.Lock()
				if _, ok := requests[t.RequestID]; ok {
					ch <- t.RequestID
				}
				mu.Unlock()
			}
		}
	}()
	return ch
}
