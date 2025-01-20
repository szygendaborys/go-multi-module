package main

import (
	"fmt"

	"example.com/learning"
)

func BillPlayground() {
	bill := learning.NewBill("test bill")

	fmt.Println(bill)

	bill.AddPrice(2.04)
	bill.AddPrice(4.00)

	fmt.Println(bill.Format())
}
