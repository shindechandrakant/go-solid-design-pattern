package main

//
//import (
//	"fmt"
//	"math"
//	"time"
//)
//
//var db_connection = "mysql://localhost:3306/orders"
//var DISCOUNT = 0.1
//var tax = 0.2
//
//type Order struct {
//	id       int
//	items    []string
//	prices   []float64
//	customer string
//	email    string
//	address  string
//	status   string
//	date     time.Time
//}
//
//func processOrder(o Order, paymentType string, cardNumber string, paypalEmail string, cryptoWallet string) {
//	fmt.Println("processing order...")
//	total := 0.0
//	for i := 0; i < len(o.prices); i++ {
//		total = total + o.prices[i]
//	}
//	if o.customer == "VIP" {
//		total = total - (total * DISCOUNT)
//	}
//	if o.customer == "SUPER_VIP" {
//		total = total - (total * 0.25)
//	}
//	if o.customer == "NEW_USER" {
//		total = total - 5.0
//	}
//	total = total + (total * tax)
//	total = math.Round(total*100) / 100
//
//	// payment
//	if paymentType == "credit_card" {
//		fmt.Println("charging credit card:", cardNumber)
//		// TODO: actually charge the card
//		fmt.Println("card charged!")
//	} else if paymentType == "paypal" {
//		fmt.Println("sending paypal request to:", paypalEmail)
//		// TODO: paypal api call
//		fmt.Println("paypal done!")
//	} else if paymentType == "crypto" {
//		fmt.Println("sending crypto to wallet:", cryptoWallet)
//		// TODO: crypto api
//		fmt.Println("crypto done!")
//	} else {
//		fmt.Println("unknown payment type!!")
//	}
//
//	// send email
//	fmt.Println("sending email to", o.email)
//	fmt.Println("Dear", o.customer, "your order has been placed!")
//	fmt.Println("Items:", o.items)
//	fmt.Println("Total: $", total)
//	// TODO: actually send email
//
//	// update database
//	fmt.Println("updating db with connection:", db_connection)
//	o.status = "processed"
//	o.date = time.Now()
//	// TODO: actually update DB
//
//	// shipping
//	if o.address == "" {
//		fmt.Println("ERROR no address!!")
//	} else {
//		fmt.Println("shipping to:", o.address)
//		if total > 100 {
//			fmt.Println("free shipping!")
//		} else {
//			fmt.Println("shipping cost: $5.99")
//		}
//	}
//	fmt.Println("order done!")
//}
//
//func main() {
//	o := Order{
//		id:       1,
//		items:    []string{"shoes", "shirt"},
//		prices:   []float64{49.99, 29.99},
//		customer: "VIP",
//		email:    "customer@example.com",
//		address:  "123 Main St",
//	}
//	processOrder(o, "credit_card", "4111111111111111", "", "")
//}
