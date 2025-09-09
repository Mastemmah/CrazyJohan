package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func shredFile(filePath string) error {
	fmt.Printf("🤪 Crazy Johan found a target: %s\n", filePath)

	file, err := os.OpenFile(filePath, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("Johan couldn't get to the file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Johan doesn't know how big this file is: %w", err)
	}
	fileSize := fileInfo.Size()
	fmt.Printf("...file is %d bytes. Time for destruction!\n", fileSize)

	fmt.Println("...overwriting with zeroes (boring)...")
	zeros := make([]byte, fileSize)
	_, err = file.WriteAt(zeros, 0)
	if err != nil {
		return fmt.Errorf("Johan failed to scribble over the file with zeroes: %w", err)
	}
	file.Sync()

	fmt.Println("...scribbling with random data (CHAOS!)...")
	randomData := make([]byte, fileSize)
	if _, err := rand.Read(randomData); err != nil {
		return fmt.Errorf("Johan ran out of random ideas: %w", err)
	}
	if _, err := file.WriteAt(randomData, 0); err != nil {
		return fmt.Errorf("Johan failed to scribble on the file: %w", err)
	}
	file.Sync()

	file.Close()

	fmt.Println("...and now for the FINAL DELETION! KABOOM!")
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("the operating system is protecting the file from Johan: %w", err)
	}

	fmt.Println("✅ File destroyed. But is it enough...?")
	return nil
}

// MUMBLING MODE - to prevent data restore on SSD's due to wear leveling mechanism
func mumblingMode(filePath string) error {
	fmt.Println("\n incoherent mumbling... Johan is activating 'MUMBLING MODE'...")

	fmt.Println("\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("!!!                                                                     !!!")
	fmt.Println("!!!  'Mumbling Mode' is about to fill ALL FREE SPACE on this drive      !!!")
	fmt.Println("!!!  with random data to securely erase file traces on SSDs.            !!!")
	fmt.Println("!!!                                                                     !!!")
	fmt.Println("!!!  THIS WILL BE SLOW AND WILL MAKE YOUR SYSTEM UNRESPONSIVE.          !!!")
	fmt.Println("!!!  This is an EXTREMELY INTENSIVE operation for your drive.           !!!")
	fmt.Println("!!!                                                                     !!!")
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	for i := 5; i > 0; i-- {
		fmt.Printf("...starting in %d seconds. Press CTRL+C to abort.\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Print("\n--> Press ENTER to continue or CTRL+C to abort immediately. ")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Println("\nAlright, you asked for it... Let the mumbling begin.")

	dir := filepath.Dir(filePath)

	tempFile, err := os.CreateTemp(dir, "johan_is_mumbling_*.tmp")
	if err != nil {
		return fmt.Errorf("Johan couldn't create his giant mumbling file: %w", err)
	}
	fileName := tempFile.Name()
	defer os.Remove(fileName)
	defer tempFile.Close()

	fmt.Printf("Creating monster file: %s\n", fileName)

	chunk := make([]byte, 1024*1024)

	for {
		_, err := rand.Read(chunk)
		if err != nil {
			return fmt.Errorf("Johan ran out of garbage to write: %w", err)
		}

		_, err = tempFile.Write(chunk)
		if err != nil {
			if strings.Contains(err.Error(), "no space left on device") {
				break
			}
			return fmt.Errorf("Johan encountered an unexpected problem while filling the disk: %w RAGE QUITING!", err)
		}
	}

	fmt.Println("✅ Disk is full! The SSD controller has no choice but to clean up.")
	fmt.Println("✅ Johan do his job. The data is now truly gone.")
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: crazy_johan [-mumbling] <path_to_file_to_destroy>")
		return
	}

	useMumblingMode := false
	filePath := ""

	if len(os.Args) > 2 && os.Args[1] == "-mumbling" {
		useMumblingMode = true
		filePath = os.Args[2]
	} else {
		filePath = os.Args[1]
	}

	if err := shredFile(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	if useMumblingMode {
		if err := mumblingMode(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR DURING MUMBLING MODE: %v\n", err)
			os.Exit(1)
		}
	}
}
