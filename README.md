# TYCOON

### Problem:
You are going to build an interactive game called tycoon. The game follows the idea of Taipan (https://en.wikipedia.org/wiki/Taipan!), which is a very popular game in 1980s. The player will take the role as a trader; buying, selling, travelling, looking for a better profit, etc.

### Basic Tasks :
1. The price should not change if the command “Show” is called.
1. The price change should not be fixed at 25% up or 15% down, but in a range from 0-25% when up and 0-15% when down.
1. Every simulation day, there is cost of living of 1 currency unit. i.e. The account balance is deducted by 1 unit currency.

### Unfortunate Tasks:
1. Sort the account balance from the lowest to the highest value using Insertion sort
1. Ask the trader for particular account value:
	1. Search using binary search
	1. Print the day that account value is found/reached, or the last value checked if not found
1. Print the maximum and the minimum account value.
1. Print the difference final to initial account, and print WIN/LOST correspondingly

### Task Options:
1. Broker can’t always accept all selling request, depending on the broker warehouse availability. The warehouse availability for each item is set to be a random number between 1 to 100. It is set at the same time during price change.
1. Two broker price list. The price changes differently. To change to either broker A or B use command: “Visit A” or “Visit B”

### Coding Rules:
1. Follow the usual coding rules for labs and programming assignments!
1. Folder/filename: tycoon####/tycoon####.go, where #### is the last 4 digits of your SID. eg. tycoon7563/tycoon7563.go
P.S. File and folder naming here is N/A due to the multiple file structuring!

### Basic data and processes in the game are:
1. Available commodities to trade are: Bronze, Silver, Gold, and Diamond.
1. There is initial price of your choice, as long as initial bronze price < silver price < gold price < diamond price. The trader can buy or sell any item at the given price.
1. Initial account value should cover to buy at least 1 unit of bronze and not more than 2 units of silver.
1. The trader can keep unlimited number of cargo in his truck. The trader can only sell any item that he has in his truck. Initially his truck is empty.
1. Ask and keep the player’s name. Time to time, the game should mention the player’s name in the comment.
1. Except “Quit”, every command should be implemented as function or procedure.
1. After calling the function/procedure, the game should change simulation day. By default it is one command or iteration per day. (It is as simple integer counter representing day, starting from day 1). See Basic Task Variations for more detail.
1. At each simulation day, the price may change 50% of the time. When the price change, the probability of price is going up is 70% and is going down 30%. When the price going up, it is up by exactly 25% and if it is going down, it is down by exactly 15%. See Basic Task Variations and Task options for more detail.
1. Available initial commands are (More commands should be available, depending on your task):
	1. “Buy” units item, buy that many units of item at current price. Broker has as many items as the trader wants. 
	1. “Sell” units item, sell that many units of item at current price (if there enough item in his truck)
	1. “Show catalog” to show the current price list
	1. “Show account” to show the current account and content of his truck
	1. “Quit”, stop the program
	1. The program ends when the player enter “Quit” or it has accepted 2018 buy/sell commands.

### Interaction Idea

#### Some ideas of how interaction should look like (italic texts are input):

**Welcome sir/mam, your name is?** _Rocky_
<br>
**Rocky, as a startup, you’ve been given account 500. Trade wisely.**
<br>
**[Day 1] What’s your command?** _show catalog_
<br>
**Current commodity prices:**
<br>
**Bronze 100**
<br>
**Silver 250**
<br>
**Gold 500**
<br>
**Diamond 1000**
<br>
**[Day 2] What’s your command?** _buy 1 gold_
<br>
**Buying 1 gold bought at 500?** _yes_
<br>
**[Day 3] What’s your command?** _buy 1 silver_
<br>
**I am sorry, Rocky. You don’t have enough money.**
<br>
**[Day 4] What’s your command?** _show account_
<br>
**Your balance is 0**
<br>
**You have in your storage:**
<br>
**0 silk**
<br>
**0 silver**
<br>
**1 gold, bought at 500**
<br>
**0 diamond**
<br>
**Thank you for your business. New price list has been issued.**
<br>
**[Day 5] What’s your command?** _show catalog_
<br>
**Current commodity prices (bid/offer):**
<br>
**silk 115**
<br>
**silver 225**
<br>
**gold 575**
<br>
**diamond 600**
<br>
**[Day 6] What’s your command?** _quit_
<br>
**Rocky, we liquidate all your asset.**
<br>
**Your final asset value is 575.**
<br>
**Apparently you win.**
<br>
