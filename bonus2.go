package main

import (
	"fmt"
	"os"
	"strings"
	"archive/zip"
	"io"
)

// C:/Users/-/Desktop/files/file1.txt,C:/Users/-/Desktop/files/file2.txt,C:/Users/-/Desktop/files/file3.pdf 

func main() {

	fmt.Println("Enter File Paths: ")

	var paths string

	fmt.Scanln(&paths)
	
	s := strings.Split(paths, ",")
	fmt.Println()

	// first checks if the files exist
	for _, path := range s {
		if _ , err := os.Stat(path); !os.IsNotExist(err){
			fmt.Println(path + " --> Good! File exists")
		} 
	} 

	output := "bonus2.zip" //new zip file name

    if err := ZipFiles(output, s); err != nil {
        panic(err)
    }

    fmt.Println("Zipped File:", output)
}

func ZipFiles(filename string, files []string) error {

    newZipFile, err := os.Create(filename) //creates new zip file
    defer newZipFile.Close()

    zipWriter := zip.NewWriter(newZipFile)

    // Add files to zip
    for _, file := range files {
        if err = AddFiles(zipWriter, file); err != nil {
            return err
        }
    }
	
	defer zipWriter.Close()
    return nil
}

func AddFiles(zipWriter *zip.Writer, filename string) error {

    fileToZip, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer fileToZip.Close()

    // Get the file information
    info, _ := fileToZip.Stat()

    header, _ := zip.FileInfoHeader(info)

    // Getting only filename from the given path
	last := filename[strings.LastIndex(filename, "/") + 1 :]
	header.Name = last
    header.Method = zip.Deflate

    writer, _ := zipWriter.CreateHeader(header)

    _, err = io.Copy(writer, fileToZip) //doing copy of file to new destination
    return err
}