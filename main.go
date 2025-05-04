package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pwman/vault"
)

const vaultFile = "store/store.dat"

func main() {
	os.MkdirAll("store", 0700)

	if len(os.Args) < 2 {
		fmt.Println("Usage: pwman <init|add|get|list>")
		return
	}

	command := os.Args[1]
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter master password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	switch command {
	case "init":
		v := &vault.Vault{Entries: []vault.Entry{}}
		err := vault.SaveVaultToFile(v, password, vaultFile)
		if err != nil {
			fmt.Println("Error saving vault:", err)
			return
		}
		fmt.Println("Vault initialized and saved.")

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: pwman add <label>")
			return
		}
		label := os.Args[2]

		fmt.Print("Enter username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		fmt.Print("Enter password: ")
		pass, _ := reader.ReadString('\n')
		pass = strings.TrimSpace(pass)

		v, err := vault.LoadVaultFromFile(password, vaultFile)
		if err != nil {
			fmt.Println("Error loading vault:", err)
			return
		}

		v.Entries = append(v.Entries, vault.Entry{
			Label:    label,
			Username: username,
			Password: pass,
		})

		err = vault.SaveVaultToFile(v, password, vaultFile)
		if err != nil {
			fmt.Println("Error saving vault:", err)
			return
		}

		fmt.Println("Entry added.")

	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Usage: pwman get <label>")
			return
		}
		label := os.Args[2]

		v, err := vault.LoadVaultFromFile(password, vaultFile)
		if err != nil {
			fmt.Println("Error loading vault:", err)
			return
		}

		for _, e := range v.Entries {
			if e.Label == label {
				fmt.Println("Username:", e.Username)
				fmt.Println("Password:", e.Password)
				return
			}
		}
		fmt.Println("Entry not found.")

	case "list":
		v, err := vault.LoadVaultFromFile(password, vaultFile)
		if err != nil {
			fmt.Println("Error loading vault:", err)
			return
		}

		for _, e := range v.Entries {
			fmt.Println("-", e.Label)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
