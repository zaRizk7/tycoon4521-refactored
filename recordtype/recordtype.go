package recordtype

import "../constant"

// Item type
// Record type used to store item info,
// e.g. Bronze, Silver, Gold, Platinum
type Item struct {
	Name     string
	Price    int
	Quantity int
}

// DailyReport type,
// Record type used to store daily game info
type DailyReport struct {
	Balance int
	Day     int
}

// Player type,
// Record type used to store player info
type Player struct {
	Name        string
	Balance     int
	ReportCount int
	Item        [constant.ITEMCOUNT]Item
	DailyReport [constant.REPORTCOUNT]DailyReport
}

// Broker type,
// Record type used to store broker that sells and buys item
type Broker struct {
	Name string
	Item [constant.ITEMCOUNT]Item
}

// MenuCommand type,
// Record type used to define commands
type MenuCommand struct {
	Command     string
	ShowCommand string
	ItemName    string
	Quantity    int
}
