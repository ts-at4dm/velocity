package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func deleteTempFiles(tempDirs []string) error { 
	// Read all contents of the directory
	for _, tempDir := range tempDirs {	
		contents, err := os.ReadDir(tempDir)
		if err != nil {
			return fmt.Errorf("failed to read directory %s: %v", tempDir, err)
		}

		// Loop through each item in the directory
		for _, entry := range contents {
			entryPath := filepath.Join(tempDir, entry.Name())

			// Remove file or directory
			fmt.Printf("Deleting: %s\n", entryPath)
			err = os.RemoveAll(entryPath)
			if err != nil {
				return fmt.Errorf("failed to delete %s: %v", entryPath, err)
			}
		}
	}
	return nil
}

func optimizeServices() {

}

func main() {
	// Path to the temporary files directory
	tempDir := []string{
		`c:/users/terran.stone/downloads/a_testfiles`,
		`c:/users/terran.stone/downloads/b_testfiles`,
		`c:/users/terran.stone/downloads/c_testfiles`,
		`c:/users/terran.stone/downloads/d_testfiles`,


}

	err := deleteTempFiles(tempDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("System Cleaned.")
	}
}
