package input

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	op "cli_prot/operations"
)

func Initialize() {
	lineReader := bufio.NewReader(os.Stdin)

	fmt.Println("Golang CLI Notebook by TRPLX v0.0.1")
	fmt.Println("Type 'help' for guide")
	// while loop
	for {
		fmt.Println("Which operation do you want to perform")
		input, err := lineReader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input or error in processing",err)
			continue
		}
		input = strings.TrimSpace(input)
		if input != "end" {
			inputs := strings.Fields(input)
			if len(inputs) > 0 {
				analyseParams(inputs)
			} else {
				fmt.Println("no command provided")
			}
		} else {
			fmt.Println("Program ended successfully")
			break
		}
	}
}


func analyseParams(cmds []string) {
	command := cmds[0]
	args := cmds[1:]

	fmt.Println(args)
	switch command {
	case "hi":
		op.Respond()
	case "help":
		listBaseOperations()
// -------------add notes
	case "add-note":
		if checkArgs(args,1) {
			op.AddNotesToDefault(args[0])
		} else if checkArgs(args,2) {
			// for specified notes
			fmt.Println("specific notes feat not added yet")
		} else if checkEmptyArgs(args) {
			fmt.Println("missing args")
		} else {
			fmt.Println("excess args!")
		}
	case "a-n":
		if checkArgs(args,1) {
			op.AddNotesToDefault(args[0])
		} else if checkArgs(args,2) {
			op.AddNotesToFile(args[0],args[1])
		} else if checkEmptyArgs(args) {
			fmt.Println("missing args")
		} else {
			fmt.Println("excess args!")
		}
// --------------creating notes
	case "new-n":
		if checkArgs(args,1) {
			op.CreateNote(args[0])
		} else if checkEmptyArgs(args) {
			fmt.Println("missing args")
		} else {
			fmt.Println("excesss args!")
		}
// -----------------reading notes
	case "r-n":
		if checkArgs(args,1) {
			op.ReadNote(args[0])
		} else if checkEmptyArgs(args) {
			fmt.Println("missing args")
		} else {
			fmt.Println("excesss args!")
		}
// ------------------deleting notes
	case "d-n":
		if checkArgs(args,1) {
			op.DeleteNote(args[0])
		} else if checkEmptyArgs(args) {
			fmt.Println("missing args")
		} else {
			fmt.Println("excesss args!")
		}
// -------------list notes
	case "n-ls":
		op.ListNotes()
// --------------clearing note ops
	case "clr-n-ls":
		op.ClearNotesList()
	case "clr-n":
		op.ClearNotesList()
	default:
		fmt.Println("Unregistered command")
	}
}

func checkArgs(argSlice []string, length int) bool {
	if len(argSlice) == length {
		return true
	}
	return false
}

func checkEmptyArgs(argSlice []string) bool {
	return checkArgs(argSlice,0)
}

func listBaseOperations() {
	fmt.Println("hi -> casual greeting interaction")
	fmt.Println("notes-ls -> listing all notes created alt=n-ls")
	fmt.Println("add-note <text> ?<note_name> -> adds notes to the default note or to a specified note file alt=a-n")
	fmt.Println("<text> must have tokens or words joined by the '-' symbol")
	fmt.Println("<note_name> must be a name or word not separated by whitespace for valid processing")
	fmt.Println("read-note ?<note_name> -> reads te contents of the specified note file or the default note file if unspecified alt=r-n")
	fmt.Println("new-note <note_name> -> creates a new note for note operations alt=new-n")
	fmt.Println("clear-notes <note_name>?=default -> clears the content of the specified notes file or the default one if not provided alt=clr-n")
	fmt.Println("clear-notes-ls -> clears the content of the notes list as well as the note files alt=clr-n-ls")
	fmt.Println("delete-note -f <note_name> -> deletes the specified note file or clears the default file if the -f flag is present alt=d-n")
	fmt.Println("help -> to show this menu again")
}
