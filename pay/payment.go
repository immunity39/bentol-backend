package pay

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/mythrnr/paypayopa-sdk-go"
)

func Pay(ReservID uint, Name string, TotalAmount uint) (string, error) {
	creds := paypayopa.NewCredentials(
		paypayopa.EnvSandbox,
	)

	wp := paypayopa.NewWebPayment(creds)
	ctx := context.Background()
	int_id := int(ReservID)
	str_id := strconv.Itoa(int_id)

	res, info, err := wp.CreateQRCode(ctx, &paypayopa.CreateQRCodePayload{
		MerchantPaymentID: str_id,
		OrderDescription:  Name,
		Amount: &paypayopa.MoneyAmount{
			Amount:   int(TotalAmount),
			Currency: paypayopa.CurrencyJPY,
		},
		CodeType:     paypayopa.CodeTypeOrderQR,
		RedirectURL:  "https://paypay.ne.jp",
		RedirectType: paypayopa.RedirectTypeWebLink,
	})

	if err != nil {
		log.Fatalf("%+v", err)
		return "", err
	}

	b, _ := json.MarshalIndent(info, "", "  ")
	log.Println(string(b))

	if !info.Success() {
		log.Fatalf("%+v", info)
		return "", err
	}

	b, _ = json.MarshalIndent(res, "", "  ")
	log.Println(string(b))

	return res.URL, nil
}

func Refund(ReservID uint, Amount uint) error {
	creds := paypayopa.NewCredentials(
		paypayopa.EnvSandbox,
	)

	wp := paypayopa.NewWebPayment(creds)
	ctx := context.Background()
	int_id := int(ReservID)
	str_id := strconv.Itoa(int_id)

	_, info, err := wp.RefundPayment(ctx, &paypayopa.RefundPaymentPayload{
		MerchantRefundID: str_id,
		Amount: &paypayopa.MoneyAmount{
			Amount:   int(Amount),
			Currency: paypayopa.CurrencyJPY,
		},
	})

	if err != nil {
		log.Fatalf("%+v", err)
		return err
	}

	b, _ := json.MarshalIndent(info, "", "  ")
	log.Println(string(b))

	if !info.Success() {
		log.Fatalf("%+v", info)
		return err
	}

	return nil
}
