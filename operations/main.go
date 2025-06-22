package operations

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"path/filepath"
)

const def_fileName string = "default.txt"
const ls_fileName string = "note-list.txt"
const notesDir string = "notes"
var notesList []string
var strBuilder strings.Builder
var dirPerm os.FileMode
var filePerm os.FileMode

func Respond() {
	fmt.Println("helloooo")
}

func initPerms() {
	filePerm = os.FileMode(0644)
	dirPerm = os.FileMode(0755)
}

func processInputText(text string) string {
	strChunks := strings.Split(text,"-")
	text = strings.Join(strChunks," ")
	return text
}

func formatInputText(text string) string {
	strBuilder.Reset()
	strBuilder.WriteString(text)
	strBuilder.WriteString("\n")
	return strBuilder.String()
}
	
func AddNotesToDefault(text string) {
	// string procesing
	// assuming '-' is the delimiter
	text = processInputText(text)
	initPerms()
	filePath := filepath.Join(notesDir,def_fileName)

	file,err := os.OpenFile(filePath,os.O_RDWR | os.O_CREATE | os.O_APPEND, filePerm)
	if err != nil {
		fmt.Println("err initializing the notes file")
		return
	}
	defer file.Close()

	text = formatInputText(text)

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("err adding notes to default file")
		return
	} else {
		fmt.Println("note added to default note file")
		return
	}

}

func AddNotesToFile(text string, filename string) {
	initPerms()
		
	targetFile := appendTextExt(filename)
	targetFile = filepath.Join(notesDir,targetFile)

	text = processInputText(text)
	text = formatInputText(text)

	file,err := os.OpenFile(targetFile,os.O_CREATE | os.O_APPEND | os.O_RDWR, filePerm)
	if err != nil {
		fmt.Println("err in creating target note file")
		return
	}

	_,err = file.WriteString(text)
	if err != nil {
		fmt.Println("err in adding note to target file")
		return
	} else {
		AddNoteToList(filename)

		strBuilder.Reset()
		strBuilder.WriteString("added notes to ")
		strBuilder.WriteString(targetFile)
		strBuilder.WriteString(" notes file")
		msg := strBuilder.String()
		fmt.Println(msg)
	}
}

func AddNoteToList(noteName string) {
	if len(notesList) > 0 {
		notesList = nil
	}

	filename := ListNotes()
	
	formattedText := strings.Join(notesList,"\n")

	file,err := os.OpenFile(filename,os.O_CREATE | os.O_TRUNC, filePerm)
	if err != nil {
		fmt.Println("err in adding new note to notes list",err)
		return 
	}
	defer file.Close()
	
	_,err = file.WriteString(formattedText)
	if err != nil {
		fmt.Println("err while writing to notes list",err)
		return
	}

	fmt.Printf("added %s to notes list\n",noteName)
}

func CreateNote(noteName string) {
	initPerms()

	targetFile := appendTextExt(noteName)
	targetFile = filepath.Join(notesDir,targetFile)

	file,err := os.OpenFile(targetFile,os.O_WRONLY | os.O_CREATE | os.O_TRUNC, filePerm)
	if err != nil {
		fmt.Println("err while creating new note file",err)
		return 
	}
	defer file.Close()

	AddNoteToList(noteName)

}

func appendTextExt(filename string) string {
	strBuilder.Reset()
	strBuilder.WriteString(filename)
	strBuilder.WriteString(".txt")
	targetFile := strBuilder.String()
	return targetFile
}

func ListNotes() string {
	initPerms()
	// reading dir for actual list
	filename := filepath.Join(notesDir,ls_fileName)

	file,err := os.OpenFile(filename,os.O_CREATE | os.O_APPEND | os.O_RDWR, filePerm)
	if err != nil {
		fmt.Println("err opening note list file")
	}
	defer file.Close()

	targetDir := filepath.Join("./",notesDir)

	notes,err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Println("err while getting notes list",err)
	}

	var fileList []string

	for _,note := range notes {
		if !note.IsDir() {
			if note.Name() != "default.txt" && note.Name() != "note-list.txt" {
				fname := note.Name()
				fname = strings.TrimSuffix(fname,".txt")
				fileList = append(fileList,fname)
			}
		}
	}

	formattedStr := strings.Join(fileList,"\n")
	notesList = strings.Split(formattedStr,"\n")

	_,err = file.WriteString(formattedStr)
	if err != nil {
		fmt.Println("err adding processed notes list to notes file",err)
	}

	fmt.Printf("There are %d notes in storage \n",len(notesList))
	fmt.Println(formattedStr)

	return filename

}

func ClearNotesList() {
	initPerms()
	ListNotes()

	filename := filepath.Join(notesDir,ls_fileName)
	file,err := os.OpenFile(filename,os.O_CREATE | os.O_TRUNC | os.O_WRONLY, filePerm)
	if err != nil {
		fmt.Println("err in setting up notes list for deletion",err)
		return
	}
	defer file.Close()

	for i := 0; i < len(notesList); i++ {
		note := appendTextExt(notesList[i])
		fname := filepath.Join(notesDir,note)
		err = os.Remove(fname)
		if err != nil {
			fmt.Println("err in clearing note files")
			return
		}


	}
	
	fmt.Println("notes list file cleared and the notes")
	
}

func getStringIndex(slice []string, targetStr string) int {
	if len(slice) > 0 {
		for i := 0; i < len(slice); i++ {
			if slice[i] == targetStr {
				return i
			}
		}
	}
	return -1
}

func DeleteNote(noteName string) {
	initPerms()
	
	// file ops for deletion
	targetFile := appendTextExt(noteName)
	targetFile = filepath.Join(notesDir,targetFile)

	err := os.Remove(targetFile)
	if err != nil {
		fmt.Println("err while deleting note file",err)
	}

	fmt.Printf("Deleted %s note file\n",noteName)

	// removing from slice buh realizing that the slice updates from the existing files and made the util func for nothing eitherways lesson learnt
}

func ReadNote(noteName string) {
	targetFile := appendTextExt(noteName)
	targetFile = filepath.Join(notesDir,targetFile)

	file,err := os.OpenFile(targetFile,os.O_RDWR | os.O_APPEND, filePerm)
	if err != nil {
		fmt.Println("err in preparing note file for reading",err)
	}
	defer file.Close()

	fileReader := bufio.NewScanner(file)
	for fileReader.Scan() {
		fmt.Println(fileReader.Text())
	}

	if err := fileReader.Err(); err != nil {
		fmt.Println("err in reading note file",err)
	}

}

