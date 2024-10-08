package handlers

import (
	"fmt"
	"os"
	"template/database"
	"template/function"
	"template/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Function for Display Failed payment page
func FailedView(c *fiber.Ctx) error {

	transID := c.Params("TransID")
	//fmt.Println(transID)

	//=============Update Transaction status from Transaction id===============
	currentTime := time.Now()
	database.DB.Db.Table("transaction").Where("transaction_id = ?", transID).UpdateColumns(models.FailedStatus{Status: "Declined", Substatus: 8, Response_timestamp: currentTime})

	//=============Fetch Transaction details from Transaction id===============
	transData := models.Transaction_Pay{}
	database.DB.Db.Table("transaction").Where("transaction_id = ?", transID).Find(&transData)
	mID := transData.Client_id

	//  Email///
	template_code := "PAYMENT-STATUS"
	getName := function.GetNameByMID(mID)
	getEmail := function.GetEmailByMID(mID)
	hashCode := transData.Response_hash
	num := transData.Receivedamount
	amount := fmt.Sprintf("%.2f", num) // Convert to string with 2 decimal places
	crypto := transData.Receivedcurrency
	status := transData.Status

	emailData := models.EmailData{FullName: getName, Email: getEmail, Status: status, Amount: amount, Crypto: crypto, HashCode: hashCode, TransID: transID}

	err := function.SendEmail(template_code, emailData)
	if err != nil {
		fmt.Println("issue sending verification email")
	}

	// Fetch Webhook Url from client store
	storeData := models.ClientStore{}
	database.DB.Db.Table("client_store").Where("client_id = ?", mID).Find(&storeData)
	webhookURL := storeData.Webhookurl
	redirectURL := ""
	if webhookURL != "" {
		redirectURL = webhookURL + "?referanceID=" + transData.Customerrefid + "&transactionID=" + transData.Transaction_id + "&orderID=" + transData.Order_id + "&hash=" + transData.Response_hash + "&amount=" + fmt.Sprintf("%f", transData.Receivedamount) + "&currency=" + transData.Receivedcurrency + "&status=" + transData.Status
	}
	return c.Render("failed", fiber.Map{
		"Title":       "Declined Payment",
		"Subtitle":    "Declined Payment",
		"TransID":     transID,
		"TransData":   transData,
		"RedirectURL": redirectURL,
	})
}

// Function for Display Dispute payment page
func DisputeView(c *fiber.Ctx) error {

	transID := c.Params("TransID")
	fmt.Println(transID)

	//=============Update Transaction status from Transaction id===============
	currentTime := time.Now()
	database.DB.Db.Table("transaction").Where("transaction_id = ?", transID).UpdateColumns(models.FailedStatus{Status: "Declined", Substatus: 9, Response_timestamp: currentTime})

	//=============Fetch Transaction details from Transaction id===============
	transData := models.Transaction_Pay{}
	database.DB.Db.Table("transaction").Where("transaction_id = ?", transID).Find(&transData)
	mID := transData.Client_id

	//  Email///
	template_code := "PAYMENT-STATUS"
	getName := function.GetNameByMID(mID)
	getEmail := function.GetEmailByMID(mID)
	hashCode := transData.Response_hash
	num := transData.Receivedamount
	amount := fmt.Sprintf("%.2f", num) // Convert to string with 2 decimal places
	crypto := transData.Receivedcurrency
	status := transData.Status

	emailData := models.EmailData{FullName: getName, Email: getEmail, Status: status, Amount: amount, Crypto: crypto, HashCode: hashCode, TransID: transID}

	err := function.SendEmail(template_code, emailData)
	if err != nil {
		fmt.Println("issue sending verification email")
	}

	// Fetch Webhook Url from client store
	storeData := models.ClientStore{}
	database.DB.Db.Table("client_store").Where("client_id = ?", mID).Find(&storeData)
	webhookURL := storeData.Webhookurl

	redirectURL := ""
	if webhookURL != "" {
		redirectURL = webhookURL + "?referanceID=" + transData.Customerrefid + "&transactionID=" + transData.Transaction_id + "&orderID=" + transData.Order_id + "&hash=" + transData.Response_hash + "&amount=" + fmt.Sprintf("%f", transData.Receivedamount) + "&currency=" + transData.Receivedcurrency + "&status=" + transData.Status
	}

	return c.Render("dispute", fiber.Map{
		"Title":       "Dispute Payment",
		"Subtitle":    "Dispute Payment",
		"TransID":     transID,
		"TransData":   transData,
		"RedirectURL": redirectURL,
	})
}

// Function for Display Success payment page
func SuccessView(c *fiber.Ctx) error {

	transID := c.Params("TransID")
	fmt.Println(transID)
	//=============Fetch Invoice Details by trackid===============

	transData := models.Transaction_Pay{}
	database.DB.Db.Table("transaction").Where("transaction_id = ?", transID).Find(&transData)
	mID := transData.Client_id

	//  Email///
	template_code := "PAYMENT-STATUS"
	getName := function.GetNameByMID(mID)
	getEmail := function.GetEmailByMID(mID)
	hashCode := transData.Response_hash
	num := transData.Receivedamount
	amount := fmt.Sprintf("%.2f", num) // Convert to string with 2 decimal places
	crypto := transData.Receivedcurrency
	status := transData.Status

	emailData := models.EmailData{FullName: getName, Email: getEmail, Status: status, Amount: amount, Crypto: crypto, HashCode: hashCode, TransID: transID}

	err := function.SendEmail(template_code, emailData)
	if err != nil {
		fmt.Println("issue sending verification email")
	}

	// Fetch Webhook Url from client store
	storeData := models.ClientStore{}
	database.DB.Db.Table("client_store").Where("client_id = ?", mID).Find(&storeData)
	webhookURL := storeData.Webhookurl

	redirectURL := ""
	if webhookURL != "" {
		redirectURL = webhookURL + "?referanceID=" + transData.Customerrefid + "&transactionID=" + transData.Transaction_id + "&orderID=" + transData.Order_id + "&hash=" + transData.Response_hash + "&amount=" + fmt.Sprintf("%f", transData.Receivedamount) + "&currency=" + transData.Receivedcurrency + "&status=" + transData.Status
	}
	return c.Render("success", fiber.Map{
		"Title":       "Success Payment",
		"Subtitle":    "Success Payment",
		"TransID":     transID,
		"TransData":   transData,
		"RedirectURL": redirectURL,
	})
}

// Function for Display  payment detail on checkout with qr code
func PayQRView(c *fiber.Ctx) error {

	PaymentID := c.Query("iid")
	//fmt.Println(PaymentID)
	//=============Fetch Invoice Details by trackid===============
	invoiceData := models.Invoice_Master{}
	database.DB.Db.Table("invoice").Where("trackid = ?", PaymentID).Find(&invoiceData)
	//=============Fetch coin list ===============
	coinList := []models.CoinList{}
	database.DB.Db.Table("coin_list").Order("coin ASC").Where("status = ?", 1).Find(&coinList)
	fmt.Println(invoiceData)
	var commonURL = os.Getenv("CommonURL")
	return c.Render("checkout-pay-views", fiber.Map{
		"Title":       "Checkout",
		"Subtitle":    "Checkout",
		"CoinList":    coinList,
		"InvoiceData": invoiceData,
		"CommonURL":   commonURL,
	})
}
