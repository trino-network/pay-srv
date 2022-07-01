package api

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	contracts "github.com/trino-network/contracts/goabi"
	"github.com/trino-network/pay-srv/internal/models"
	"github.com/trino-network/pay-srv/internal/types"
	"github.com/trino-network/pay-srv/internal/util/db"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"math/big"
	"net/http"
	"time"
)

func (s *Server) WatchPay(ctx context.Context) {
	// goerli network
	go s.WatchEthPay(ctx, "goerli", "wss://goerli.infura.io/ws/v3/4cd91fcd955643eab22f2c0dcebf87e2",
		"0x889fB20cA7D3B2Bbb7bc888F18d1fE68742087Af")

	go s.Notify(ctx)
}

func (s *Server) WatchEthPay(ctx context.Context, network string, url string, contractAddr string) error {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Error().Msg(err.Error())

		return err
	}

	contractAddress := common.HexToAddress(contractAddr)

	filter, err := contracts.NewTroinPaymentFilterer(contractAddress, client)
	sink := make(chan *contracts.TroinPaymentPay)
	sub, err := filter.WatchPay(nil, sink)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			log.Error().Msg(err.Error())
			return err
		case p := <-sink:
			log.Info().Msgf("Watched payment InvoiceId: %s,  Recipient: %s, Network: %s, Amount: %s, TxHash: %s", p.InvoiceId, p.Recipient.String(), network, p.Amount, p.Raw.TxHash.Hex())
			_ = s.processInvoice(ctx, p.InvoiceId, p.Recipient.String(), p.Amount, network, p.Raw.TxHash.Hex())
		}
	}
}

func (s *Server) processInvoice(ctx context.Context, invoiceId string, recipient string, amount *big.Int, network string, txHash string) error {
	invoice, err := models.Invoices(models.InvoiceWhere.ID.EQ(invoiceId),
		models.InvoiceWhere.Recipient.EQ(recipient),
		models.InvoiceWhere.Amount.EQ(null.Int64From(amount.Int64())),
		models.InvoiceWhere.Network.EQ(network),
		models.InvoiceWhere.Status.EQ(types.GetInvoiceResponseStatusUnpaid),
	).One(ctx, s.DB)

	if err != nil || invoice == nil {
		log.Error().Msg(err.Error())
		return err
	}

	invoice.TXNHash = null.StringFrom(txHash)
	invoice.Status = types.GetInvoiceResponseStatusPaid

	invoice.UpdatedAt = time.Now()

	if err := s.updateInvoice(ctx, invoice); err != nil {
		return err
	}
	return s.notify(ctx, invoice)
}

func (s *Server) updateInvoice(ctx context.Context, invoice *models.Invoice) error {

	if err := db.WithTransaction(ctx, s.DB, func(tx boil.ContextExecutor) error {

		if _, err := invoice.Update(ctx, tx, boil.Blacklist(models.InvoiceColumns.CreatedAt)); err != nil {
			log.Debug().Err(err).Msgf("Failed to update invoice: %s ", err.Error())
			return err
		}
		return nil
	}); err != nil {
		log.Debug().Err(err).Msgf("Failed to update invoice: %s ", err.Error())
		return err
	}

	return nil
}

func (s *Server) Notify(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		invoices, _ := models.Invoices(models.InvoiceWhere.Status.EQ(types.GetInvoiceResponseStatusPaid)).All(ctx, s.DB)
		for _, v := range invoices {
			s.notify(ctx, v)
		}
	}

}

func (s *Server) notify(ctx context.Context, invoice *models.Invoice) error {
	// header 对 token进行 hmac sha256编码
	token, err := models.AccessTokens(models.AccessTokenWhere.UserID.EQ(invoice.UserID), qm.OrderBy("created_at desc")).One(ctx, s.DB)
	if err != nil {
		return err
	}
	data := make(map[string]string)
	data["invoice_id"] = invoice.ID
	data["status"] = invoice.Status
	jsonData, _ := json.Marshal(data)
	sig := s.hmacSha256(jsonData, token.Token)

	// 写数据invoice_id、status
	req, _ := http.NewRequest("POST", invoice.NotifyURL, bytes.NewReader(jsonData))
	req.Header.Set("trino-pay-Sig", fmt.Sprintf("sha256=%s", sig))
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Error().Msgf("payment notify failed: %s ", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Info().Msgf("invoice id: %s, payment notify success", invoice.ID)
		invoice.Status = types.GetInvoiceResponseStatusNotified
		if err := s.updateInvoice(ctx, invoice); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) hmacSha256(data []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
