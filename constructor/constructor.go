package constructor

import (
	"../constant"
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

	player.Name = name
	player.Balance = 0
	player.ReportCount = 0

	for i := 0; i < constant.ITEMCOUNT; i++ {
		switch i {
		case 0:
			itemName = "Bronze"
		case 1:
			itemName = "Silver"
		case 2:
			itemName = "Gold"
		case 3:
			itemName = "Platinum"
		}
		player.Item[i] = Item(itemName, 0, 0)
	}

	return player
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
