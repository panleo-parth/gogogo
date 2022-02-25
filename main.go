package main

import (
	"fmt"
	"strings"
	"regexp"
	"booking_app/helper"
)


const conference_tickets uint = 50
var conference_name = "Go Con"
var remaining_tickets uint = conference_tickets
var bookings = []string{}
var first_names = []string{}


func main(){

	
	greetings()

	var first_name string
	var last_name string
	var usertickets uint
	var	email string

	var choice int = 1

	for i:=remaining_tickets; i>0; i=remaining_tickets{
		if choice == 0{
			break
		}
		fmt.Println("Enter your first name:")
		fmt.Scan(&first_name)
		exp, _ := regexp.Compile("[A-Z][a-z]+")
		if !helper.Check_exp(exp, first_name){
			fmt.Printf("*********************Please check the name...**********************\n")
			continue
		}
		fmt.Println("Enter your last name:")
		fmt.Scan(&last_name)
		exp, _ = regexp.Compile("[A-Z][a-z]+")
		if !helper.Check_exp(exp, last_name){
			fmt.Printf("*********************Please check the name...**********************\n")
			continue
		}
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		exp, _ = regexp.Compile("[a-z][0-9]+@[a-z]+.com")
		if !helper.helper.Check_exp(exp, email){
			fmt.Printf("*********************Please check the email...**********************\n")
			continue
		}
		fmt.Println("Enter total no. of tickets:")
		fmt.Scan(&usertickets)
		isnotvalid := usertickets <= 0 || usertickets > remaining_tickets
		if isnotvalid{
			fmt.Printf("Sorry we have %v tickets available...\n\n", remaining_tickets)
			continue
		}
		remaining_tickets = remaining_tickets - usertickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will be receiving a confirmation at %v.\n", first_name, last_name, usertickets, email)
		fmt.Printf("\nRemaining tickets: %v\n", remaining_tickets)

		bookings = append(bookings, first_name + " " + last_name)

		fmt.Printf("\nWanna book more : (either a '0' for a no or anything else for a yes)...\n")
		fmt.Scan(&choice)
	}
	
	bookies()
}


func greetings(){

	fmt.Println("Welcome to our ",conference_name, " booking application....")
	fmt.Printf("We have a total of %v tickets and %v are available...\n", conference_tickets, remaining_tickets)
	fmt.Println("Book your ticket here to attend...")
}

func bookies(){
	for _, booking := range bookings{
		names := strings.Fields(booking)
		first_names = append(first_names, names[0])
	}
	fmt.Printf("\nThe bookies are :\n")
	fmt.Printf("%v", first_names)
}