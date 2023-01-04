package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// This example opens https://github.com/, searches for "git",
// and then gets the header element which gives the description for Git.
func main() {
	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()

	// Create a new page
	defer browser.MustClose()
	page := browser.MustPage("https://discord.com/login?redirect_to=%2Fchannels%2F%40me")

	var signinusername string
	fmt.Print("Enter your discord username or email to sign in: ")
	fmt.Scanln(&signinusername)

	var password string
	fmt.Print("Enter your discord password: ")
	fmt.Scanln(&password)

	var blacklistname string
	fmt.Print("Enter your discord username without the numbers: ")
	fmt.Scanln(&blacklistname)

	var channelurl string
	fmt.Print("Enter the name of the channel: ")
	fmt.Scanln(&channelurl)

	var message string
	fmt.Print("Enter the message you would like to send to the members: ")
	fmt.Scanln(&message)

	//sign in
	page.MustElement("input")
	inputs := page.MustElements("input")
	inputs[0].MustInput(signinusername).MustType(input.Enter)
	inputs[1].MustInput(password).MustType(input.Enter)
	time.Sleep(1 * time.Second)
	page.MustElement("button[type='submit']").MustClick()

	// wait for page to load
	time.Sleep(5 * time.Second)

	//select channel
	page.MustElement(fmt.Sprintf(`[data-dnd-name="TOKENS"] .wrapper-3kah-n`)).MustClick()
	// wait for channel page load
	time.Sleep(5 * time.Second)
	fmt.Println("Found channel page")

	// show members list
	page.MustElement("[aria-label='Show Member List']").MustClick()
	//wait for members list to load
	time.Sleep(2 * time.Second)

	//get all members
	members := page.MustElements(`[data-list-item-id^="members-1023984313873735711___"]`)
	fmt.Println("All found members: ")
	for _, i := range members {
		username := i.MustElement(".username-u-ebrn").MustText()
		fmt.Print(username, ", ")
	}
	fmt.Println()
	sendMembersMessage(page, members, blacklistname, message)

}

func sendMembersMessage(page *rod.Page, members rod.Elements, blacklistname string, message string) {
	for index := range members {
		element := page.MustElements(`[data-list-item-id^="members-1023984313873735711___"]`)[index]
		username := element.MustElement(".username-u-ebrn").MustText()
		if username != blacklistname {
			element.MustClick()
			// wait for dialogue
			time.Sleep(2 * time.Second)
			fmt.Println("Username", username)

			page.MustElement(fmt.Sprintf("input[placeholder='Message @%v']", username)).MustInput(message).MustType(input.Enter)
			fmt.Println("Sent text to", username)
			page.MustElement(fmt.Sprintf(`[data-dnd-name="TOKENS"] .wrapper-3kah-n`)).MustClick()
			//wait for channel page load
			time.Sleep(5 * time.Second)
			fmt.Println("Back on the channel page...")
		}
	}
}
