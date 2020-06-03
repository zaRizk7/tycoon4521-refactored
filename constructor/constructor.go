package constructor

import (
	"../constant"
	"../helper"
	"../recordtype"
)

// Item constructor
func Item(name string, price int, quantity int) recordtype.Item {
	var item recordtype.Item

	item.Name = name
	item.Price = price
	item.Quantity = quantity

	return item
}

// DailyReport constructor
func DailyReport(balance int, day int) recordtype.DailyReport {
	var dailyReport recordtype.DailyReport

	dailyReport.Balance = balance
	dailyReport.Day = day

	return dailyReport
}

// Player constructor
func Player(name string) recordtype.Player {
	var player recordtype.Player

	var itemName string
	var itemPrice int

	player.Name = name
	player.Balance = 2000
	player.ReportCount = 0

	for i := 0; i < constant.ITEMCOUNT; i++ {
		itemName = helper.SetItemName(i)
		itemPrice = helper.SetInitialPrice(i)
		player.Item[i] = Item(itemName, itemPrice, 0)
	}

	return player
}

// Broker constructor
func Broker(name string) recordtype.Broker {
	var broker recordtype.Broker
	var itemName string
	var itemPrice int
	var itemQuantity int

	broker.Name = name

	for i := 0; i < constant.ITEMCOUNT; i++ {
		itemName = helper.SetItemName(i)
		itemPrice = helper.SetInitialPrice(i)
		itemQuantity = helper.GenerateRandomInteger(5, 20)
		broker.Item[i] = Item(itemName, itemPrice, itemQuantity)
	}

	return broker
}

// MenuCommand constructor
func MenuCommand(command string, showCommand string, itemName string, quantity int) recordtype.MenuCommand {
	var menuCommand recordtype.MenuCommand

	menuCommand.Command = command
	menuCommand.ShowCommand = showCommand
	menuCommand.ItemName = itemName
	menuCommand.Quantity = quantity

	return menuCommand
}
