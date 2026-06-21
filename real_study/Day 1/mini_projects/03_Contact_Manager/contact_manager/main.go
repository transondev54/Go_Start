package main

import (
	"encoding/json"
	"fmt"
	"os"
)
type Contact struct {
	ID int
	Name  string
	Phone string
	Email string
}
type ContactAddDTO struct {
	Name  string
	Phone string
	Email string
}
var contacts []Contact
var nextID int = 1
func viewAllContacts() {	
	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
	}
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Phone: %s, Email: %s\n", contact.ID, contact.Name, contact.Phone, contact.Email)
	}
}
func addContact(contact ContactAddDTO) {
	// Code to add contact to the database or in-memory storage
	contacts = append(contacts, Contact{
		ID:    nextID,
		Name:  contact.Name,
		Phone: contact.Phone,
		Email: contact.Email,
	})
	nextID++
	viewAllContacts()
}

func updateContact(id int, updatedContact Contact) {
	for i, contact := range contacts {
		if contact.ID == id {
			if updatedContact.Name != "" {
				contact.Name = updatedContact.Name
			}
			if updatedContact.Phone != "" {
				contact.Phone = updatedContact.Phone
			}
			if updatedContact.Email != "" {
				contact.Email = updatedContact.Email
			}
			contacts[i] = contact
			return
		}
	}
	fmt.Println("Contact not found.")
}
func deleteContact(id int) {
	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			fmt.Println("Contact deleted successfully.")
			return
		}
	}
	fmt.Println("Contact not found.")
}
func saveContactsToFile(contacts []Contact) {
	data, err := json.Marshal(contacts)
	if err != nil {
		fmt.Println("Error saving contacts to file:", err)
		return
	}
	err = os.WriteFile("contacts.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing contacts to file:", err)
	}
}
func main() {
	if contactsData, err := os.ReadFile("contacts.json"); err == nil {
		json.Unmarshal(contactsData, &contacts)
		if len(contacts) > 0 {
			nextID = contacts[len(contacts)-1].ID + 1
		}
	} else {
		fmt.Println("No existing contacts found. Starting with an empty contact list.")
	}
	fmt.Println("Welcome to the Contact Manager!")
	fmt.Println("╚════════════════════════════╝")		
	
	for {
		fmt.Println("\n1. Add Contact")
		fmt.Println("2. View All")
		fmt.Println("3. View Contact (by ID)")
		fmt.Println("4. Edit Contact")
		fmt.Println("5. Delete Contact")
		fmt.Println("6. Search Contact")
		fmt.Println("7. Exit")
		
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		
		switch choice {
		case 1:
			fmt.Println("Add Contact")
			fmt.Print("Enter Name: ")
			var name string
			fmt.Scan(&name)
			fmt.Print("Enter Phone: ")
			var phone string
			fmt.Scan(&phone)
			fmt.Print("Enter Email: ")
			var email string
			fmt.Scan(&email)
			contact := ContactAddDTO{
				Name:  name,
				Phone: phone,
				Email: email,
			}
			//convert item to json
			contactJSON, err := json.Marshal(contact)
			if err != nil {
				fmt.Println("Error converting contact to JSON:", err)
				continue
			}
			fmt.Println("Contact added (as JSON):", string(contactJSON))	
			addContact(contact)
		case 2:
			fmt.Println("View All Contacts")
			viewAllContacts()
		case 4:
			fmt.Println("Edit Contact")
			fmt.Print("Enter ID: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Enter New Name (leave blank to keep unchanged): ")
			var newName string
			fmt.Scan(&newName)
			fmt.Print("Enter New Phone (leave blank to keep unchanged): ")
			var newPhone string
			fmt.Scan(&newPhone)
			fmt.Print("Enter New Email (leave blank to keep unchanged): ")
			var newEmail string
			fmt.Scan(&newEmail)
			updateContact(id, Contact{
				Name:  newName,
				Phone: newPhone,
				Email: newEmail,
			})
		case 5:
			fmt.Println("Delete Contact")
			fmt.Print("Enter ID: ")
			var id int
			fmt.Scan(&id)
			deleteContact(id)
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		saveContactsToFile(contacts)
	}
}