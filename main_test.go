package main

import (
	"os"
	"testing"
)

func TestShredFile_Success(t *testing.T) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "johan-test-*")
	if err != nil {
		t.Fatalf("Nie udało się stworzyć pliku tymczasowego: %v", err)
	}

	defer t.Cleanup(func() {
		os.Remove(tmpFile.Name())
	})

	_, err = tmpFile.WriteString("Johan was here")
	if err != nil {
		t.Fatalf("Nie udało się zapisać do pliku tymczasowego: %v", err)
	}
	tmpFile.Close()

	err = shredFile(tmpFile.Name())

	if err != nil {
		t.Errorf("shredFile() zwróciło niespodziewany błąd: %v", err)
	}

	_, err = os.Stat(tmpFile.Name())
	if err == nil {
		t.Error("Plik powinien zostać usunięty, ale wciąż istnieje.")
	} else if !os.IsNotExist(err) {
		t.Errorf("Oczekiwano błędu 'file does not exist', ale otrzymano: %v", err)
	}
}

func TestShredFile_FileNotFound(t *testing.T) {
	t.Helper()

	nonExistentFilePath := t.TempDir() + "/definitely-not-a-real-file.txt"

	err := shredFile(nonExistentFilePath)

	if err == nil {
		t.Error("Oczekiwano błędu dla nieistniejącego pliku, ale nie otrzymano go.")
	}
}
