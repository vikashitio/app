package handlers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"template/database"
	"template/function"
	"template/models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type APIResponsePaylinkSuccess struct {
	Status      string
	ReferanceID string
	PayURL      string
}
type APIResponseFailed struct {
	Status string
	Error  string
}

type APIResponseBalanceSuccess struct {
	Receivedcurrency string
	Balance          string
}

type CreateLink struct {
	ProductName         string  `json:"ProductName"`
	Description         string  `json:"Description"`
	Currency            string  `json:"Currency"`
	Amount              float64 `json:"Amount"`
	CustomerName        string  `json:"CustomerName"`
	CustomerEmail       string  `json:"CustomerEmail"`
	OrderID             string  `json:"OrderID"`
	Is_fee_paid_by_user bool    `json:"is_fee_paid_by_user"`
}

// Function for Generate Pay Link By API
func ApiPaymentLink(c *fiber.Ctx) error {
	apiError := ""
	// Retrieve a specific header
	apikey := c.Get("Apikey")

	link := new(CreateLink)

	// Parse the request body into the User struct
	if err := c.BodyParser(link); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON !!",
		})
	}

	MID, errorx := function.GetMIDByApikey(apikey)
	// Retrieve data from header
	price_currency := strings.TrimSpace(link.Currency)

	requestedamount := link.Amount

	//fmt.Println("Amount - > ".requestedamount)

	productName := strings.TrimSpace(link.ProductName)
	description := strings.TrimSpace(link.Description)
	customerName := strings.TrimSpace(link.CustomerName)
	customerEmail := strings.TrimSpace(link.CustomerEmail)
	orderID := strings.TrimSpace(link.OrderID)
	is_fee_paid_by_user := link.Is_fee_paid_by_user

	fmt.Println(customerName, customerEmail, orderID, is_fee_paid_by_user)

	if errorx != "" {
		fmt.Println(errorx)
		apiError = errorx
	} else if price_currency == "" {
		apiError = "Currency Not Found"
	} else if link.Amount == 0.0 {
		apiError = "Amount Not Found"
	} else if productName == "" {
		apiError = "Product Name Name Not Found"
	} else if description == "" {
		apiError = "Description Not Found"
	} else if orderID == "" {
		apiError = "orderID Not Found"
	} else {

		// Generate randomly Transaction ID
		trackID, err := function.GenerateRandomID(16) // 16 bytes will give us a 32 character hex string
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to generate random ID",
			})
		}

		Ip := c.Context().RemoteIP().String()
		qry := models.Invoice_Master{Client_id: MID, Requestedamount: requestedamount, Requestedcurrency: price_currency, Product_name: productName, Description: description, Ip: Ip, Trackid: trackID, Name: customerName, Email: customerEmail, Order_id: orderID, Is_fee_paid_by_user: is_fee_paid_by_user}
		result := database.DB.Db.Table("invoice").Select("client_id", "requestedamount", "requestedcurrency", "product_name", "description", "ip", "trackid", "name", "email", "order_id", "is_fee_paid_by_user").Create(&qry)
		invoice_id := strconv.FormatUint(uint64(qry.Invoice_id), 10)
		fmt.Println(invoice_id)
		if result.Error != nil {
			fmt.Println("Data Not Inserted")

		}
		var commonURL = os.Getenv("CommonURL")
		PayURL := commonURL + "/pay?iid=" + trackID
		status := "Ok"

		response := APIResponsePaylinkSuccess{
			Status:      status,
			ReferanceID: trackID,
			PayURL:      PayURL,
		}

		return c.JSON(response)
	}
	status := "Fail"
	response := APIResponseFailed{
		Status: status,
		Error:  apiError,
	}

	return c.JSON(response)
}

// Function for Get Balance By API
func ApiBalanceByCrypto(c *fiber.Ctx) error {

	apiError := ""
	// Retrieve a specific header
	apikey := c.Get("Apikey")

	MID, errorx := function.GetMIDByApikey(apikey)
	fmt.Println(MID)
	if errorx != "" {
		fmt.Println(errorx)
		apiError = errorx

	} else {

		assetList := []APIResponseBalanceSuccess{}
		var totalWallet int64
		// fetch query for wallet with balance
		database.DB.Db.Table("transaction").Select("assetid, receivedcurrency, SUM(receivedamount)  as balance").Where("client_id = ? AND status = ?", MID, "Success").Group("assetid,receivedcurrency").Having("COUNT(assetid) > ?", 0).Order("assetid ASC").Find(&assetList).Count(&totalWallet)

		return c.JSON(assetList)
	}

	status := "Fail"
	response := APIResponseFailed{
		Status: status,
		Error:  apiError,
	}

	return c.JSON(response)
}

// API Function for Get Transaction details by ID
func ApiTransactionByTransID(c *fiber.Ctx) error {

	apiError := ""
	// Retrieve a specific header
	apikey := c.Get("Apikey")
	TransID := c.Params("TransID")
	fmt.Println("TransID==>", TransID)

	MID, errorx := function.GetMIDByApikey(apikey)
	fmt.Println(MID)
	if errorx != "" {
		fmt.Println(errorx)
		apiError = errorx
	} else if TransID == "" {
		apiError = "TransID Not Found"

	} else {

		transData := models.Transaction_Pay{}
		database.DB.Db.Table("transaction").Where("transaction_id = ? AND client_id = ? ", TransID, MID).Omit("client_id", "assetid", "response_json").Find(&transData)

		return c.JSON(transData)
	}

	status := "Fail"
	response := APIResponseFailed{
		Status: status,
		Error:  apiError,
	}

	return c.JSON(response)
}

// API Function for Get Transaction details by ID
func ApiTransactionByReferenceID(c *fiber.Ctx) error {

	apiError := ""
	// Retrieve a specific header
	apikey := c.Get("Apikey")
	ReferenceID := c.Params("ReferenceID")
	fmt.Println("TransID==>", ReferenceID)

	MID, errorx := function.GetMIDByApikey(apikey)
	fmt.Println(MID)
	if errorx != "" {
		fmt.Println(errorx)
		apiError = errorx
	} else if ReferenceID == "" {
		apiError = "ReferenceID Not Found"

	} else {

		transData := models.Transaction_Pay{}
		database.DB.Db.Table("transaction").Where("customerrefid = ? AND client_id = ? ", ReferenceID, MID).Omit("client_id", "assetid", "response_json").Find(&transData)

		return c.JSON(transData)
	}

	status := "Fail"
	response := APIResponseFailed{
		Status: status,
		Error:  apiError,
	}

	return c.JSON(response)
}

// API Function for Get Transaction List last 10
func ApiTransactionList(c *fiber.Ctx) error {

	apiError := ""
	// Retrieve a specific header
	apikey := c.Get("Apikey")

	MID, errorx := function.GetMIDByApikey(apikey)
	fmt.Println(MID)
	if errorx != "" {
		fmt.Println(errorx)
		apiError = errorx

	} else {

		transData := []models.Transaction_Pay{}
		database.DB.Db.Table("transaction").Where("client_id = ? ", MID).Omit("client_id", "assetid", "response_json").Order("id DESC").Limit(10).Find(&transData)

		return c.JSON(transData)
	}

	status := "Fail"
	response := APIResponseFailed{
		Status: status,
		Error:  apiError,
	}

	return c.JSON(response)
}
