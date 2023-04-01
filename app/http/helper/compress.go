package helper

import (
	"archive/zip"
	"io"
	"os"
)

// ZipFiles 压缩文件
func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		fileToZip, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		// 添加到zip文件
		fileInfo, err := fileToZip.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}

		header.Name = fileInfo.Name()

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return err
		}
	}

	return nil
}

// UnzipFile 解压文件
func UnzipFile(filename string, dest string) error {
	// 读取zip文件
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 解压文件
	for _, file := range reader.File {
		// 解压文件
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		// 创建解压文件
		targetFile := dest + "/" + file.Name
		if file.FileInfo().IsDir() {
			os.MkdirAll(targetFile, os.ModePerm)
		} else {
			if err := os.MkdirAll(dest, os.ModePerm); err != nil {
				return err
			}
			targetFile, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer targetFile.Close()

			// 写入文件
			if _, err := io.Copy(targetFile, fileReader); err != nil {
				return err
			}
		}
	}

	return nil
}
