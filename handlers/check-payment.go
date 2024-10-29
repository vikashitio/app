package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	return_url := storeData.Return_url

	redirectURL := ""
	if return_url != "" {
		redirectURL = return_url + "?referanceID=" + transData.Customerrefid + "&transactionID=" + transData.Transaction_id + "&orderID=" + transData.Order_id + "&hash=" + transData.Response_hash + "&amount=" + fmt.Sprintf("%f", transData.Receivedamount) + "&currency=" + transData.Receivedcurrency + "&status=" + transData.Status
	}

	// For Post response on webhook URL
	webhookurl := storeData.Webhookurl //Get Webhook Url
	if webhookurl != "" {

		// Convert the struct to JSON
		jsonData, err := json.Marshal(transData)
		if err != nil {
			fmt.Println("Error in jsonData")
		}
		// Create the POST request
		resp, err := http.Post(webhookurl, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error in resp")
		}
		defer resp.Body.Close()

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
	return_url := storeData.Return_url

	redirectURL := ""
	if return_url != "" {
		redirectURL = return_url + "?referanceID=" + transData.Customerrefid + "&transactionID=" + transData.Transaction_id + "&orderID=" + transData.Order_id + "&hash=" + transData.Response_hash + "&amount=" + fmt.Sprintf("%f", transData.Receivedamount) + "&currency=" + transData.Receivedcurrency + "&status=" + transData.Status
	}

	// For Post response on webhook URL
	webhookurl := storeData.Webhookurl //Get Webhook Url
	if webhookurl != "" {

		// Convert the struct to JSON
		jsonData, err := json.Marshal(transData)
		if err != nil {
			fmt.Println("Error in jsonData")
		}
		// Create the POST request
		resp, err := http.Post(webhookurl, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error in resp")
		}
		defer resp.Body.Close()

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

	//fmt.Println(" Payment Status Success = ", transData.Status)

	if transData.Status != "Success" {
		return c.Redirect("/failed/" + transID)
	}

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
	return_url := storeData.Return_url //Get Return URL

	redirectURL := ""
	if return_url != "" {
		redirectURL = return_url + "?referanceID=" + transData.Customerrefid + "&transactionID=" + transData.Transaction_id + "&orderID=" + transData.Order_id + "&hash=" + transData.Response_hash + "&amount=" + fmt.Sprintf("%f", transData.Receivedamount) + "&currency=" + transData.Receivedcurrency + "&status=" + transData.Status
	}
	// For Post response on webhook URL
	webhookurl := storeData.Webhookurl //Get Webhook Url
	if webhookurl != "" {

		// Convert the struct to JSON
		jsonData, err := json.Marshal(transData)
		if err != nil {
			fmt.Println("Error in jsonData")
		}
		// Create the POST request
		resp, err := http.Post(webhookurl, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error in resp")
		}
		defer resp.Body.Close()

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
	MID := invoiceData.Client_id
	//=============Fetch coin list ===============
	coinList := []models.CoinList{}
	//database.DB.Db.Table("coin_list").Order("coin ASC").Where("status = ?", 1).Find(&coinList)
	database.DB.Db.Table("coin_list A ").Select("a.*").Joins("LEFT JOIN client_wallet B ON A.coin_id = B.assetid ").Where(" a.status = ? AND B.client_id = ? AND B.status = ?", 1, MID, 1).Order("a.coin_title asc").Find(&coinList)

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
