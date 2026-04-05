package main

import (
	"fmt"
	"math"
)

// God struct that does everything
type App struct {
	users    [][]string      // [name, email, role]
	orders   [][]interface{} // [userIndex, product, qty, price]
	discount float64
}

func (a *App) DoStuff(action string, data []interface{}) interface{} {
	if action == "addUser" {
		name := data[0].(string)
		email := data[1].(string)
		role := data[2].(string)
		// validate
		if name == "" || email == "" {
			fmt.Println("bad user data")
			return nil
		}
		if role != "admin" && role != "customer" {
			fmt.Println("bad role")
			return nil
		}
		a.users = append(a.users, []string{name, email, role})
		fmt.Println("user added:", name)
		return len(a.users) - 1

	} else if action == "placeOrder" {
		userIdx := data[0].(int)
		product := data[1].(string)
		qty := data[2].(int)
		price := data[3].(float64)

		if userIdx >= len(a.users) {
			fmt.Println("no user")
			return nil
		}
		if qty <= 0 || price <= 0 {
			fmt.Println("bad order")
			return nil
		}

		total := float64(qty) * price
		// apply discount
		if a.users[userIdx][2] == "admin" {
			total = total * 0.8 // admin gets 20% off
		} else if a.discount > 0 {
			total = total * (1 - a.discount)
		}
		// round
		total = math.Round(total*100) / 100

		a.orders = append(a.orders, []interface{}{userIdx, product, qty, price, total})
		fmt.Println("order placed for", a.users[userIdx][0], "total:", total)
		return total

	} else if action == "report" {
		fmt.Println("=== REPORT ===")
		for i, o := range a.orders {
			uIdx := o[0].(int)
			fmt.Printf("Order %d: user=%s product=%s qty=%d total=%.2f\n",
				i, a.users[uIdx][0], o[1], o[2], o[4])
		}
		return nil

	} else if action == "sendEmail" {
		userIdx := data[0].(int)
		msg := data[1].(string)
		if userIdx >= len(a.users) {
			fmt.Println("no user")
			return nil
		}
		// pretend to send
		fmt.Printf("Sending email to %s: %s\n", a.users[userIdx][1], msg)
		return true
	}

	fmt.Println("unknown action")
	return nil
}

func main() {
	app := &App{discount: 0.1}

	i := app.DoStuff("addUser", []interface{}{"Alice", "alice@example.com", "admin"})
	j := app.DoStuff("addUser", []interface{}{"Bob", "bob@example.com", "customer"})

	app.DoStuff("placeOrder", []interface{}{i.(int), "Widget", 3, 25.0})
	app.DoStuff("placeOrder", []interface{}{j.(int), "Gadget", 2, 50.0})

	app.DoStuff("sendEmail", []interface{}{i.(int), "Your order is confirmed!"})
	app.DoStuff("report", nil)
}
