package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func copyFile(src, dst string, wg *sync.WaitGroup) {
	defer wg.Done()

	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destFile.Close()

	fmt.Printf("Copying %s to %s...\n", src, dst)

	buf := make([]byte, 1024)
	var totalBytesCopied int64

	for {
		n, err := sourceFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		if n == 0 {
			break
		}

		if _, err := destFile.Write(buf[:n]); err != nil {
			fmt.Println(err)
			return
		}

		totalBytesCopied += int64(n)
		fmt.Printf("\rProgress: %d bytes copied", totalBytesCopied)
	}

	fmt.Printf("\nCopied %s to %s\n", src, dst)
}

func main() {
	sourceDirectory := "/mnt/media/shows"
	destinationDirectory := "/mnt/media2/shows/shows"
	numProcesses := 4

	var wg sync.WaitGroup

	files, err := os.ReadDir(sourceDirectory)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		sourcePath := filepath.Join(sourceDirectory, file.Name())
		destinationPath := filepath.Join(destinationDirectory, file.Name())

		wg.Add(1)
		go copyFile(sourcePath, destinationPath, &wg)

		// Limit the number of concurrent goroutines
		if (i+1)%numProcesses == 0 {
			wg.Wait()
		}
	}

	wg.Wait()
}
