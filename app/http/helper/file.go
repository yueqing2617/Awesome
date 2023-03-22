package helper

import "strings"

// IsImage Determine if the file is an image.
func IsImage(name string) bool {
	// 获取文件后缀
	ext := GetFileExt(name)
	// 图片后缀
	imageExt := []string{"jpg", "jpeg", "png", "gif", "bmp", "webp"}
	for _, v := range imageExt {
		if v == ext {
			return true
		}
	}
	return false
}

// GetFileExt Get the file extension.
func GetFileExt(name string) string {
	// 获取文件后缀
	ext := strings.Split(name, ".")
	return ext[len(ext)-1]
}

// IsVideo Determine if the file is a video.
func IsVideo(name string) bool {
	// 获取文件后缀
	ext := GetFileExt(name)
	// 视频后缀
	videoExt := []string{"mp4", "avi", "rmvb", "rm", "flv", "3gp", "mkv", "mov", "wmv", "webm", "ogg"}
	for _, v := range videoExt {
		if v == ext {
			return true
		}
	}
	return false
}

// IsAudio Determine if the file is an audio.
func IsAudio(name string) bool {
	// 获取文件后缀
	ext := GetFileExt(name)
	// 音频后缀
	audioExt := []string{"mp3", "wav", "wma", "ogg", "ape", "flac", "aac"}
	for _, v := range audioExt {
		if v == ext {
			return true
		}
	}
	return false
}

// IsDoc Determine if the file is a document.
func IsDoc(name string) bool {
	// 获取文件后缀
	ext := GetFileExt(name)
	// 文档后缀
	docExt := []string{"doc", "docx", "xls", "xlsx", "ppt", "pptx", "pdf", "txt"}
	for _, v := range docExt {
		if v == ext {
			return true
		}
	}
	return false
}

// IsZip Determine if the file is a zip.
func IsZip(name string) bool {
	// 获取文件后缀
	ext := GetFileExt(name)
	// 压缩包后缀
	zipExt := []string{"zip", "rar", "7z", "tar", "gz"}
	for _, v := range zipExt {
		if v == ext {
			return true
		}
	}
	return false
}

// IsCode Determine if the file is a code.
func IsCode(name string) bool {
	// 获取文件后缀
	ext := GetFileExt(name)
	// 代码后缀
	codeExt := []string{"go", "php", "java", "js", "css", "html", "py", "sh", "c", "cpp", "h", "hpp", "cs"}
	for _, v := range codeExt {
		if v == ext {
			return true
		}
	}
	return false
}

// IsOther Determine if the file is other.
func IsOther(name string) bool {
	// 获取文件后缀
	ext := GetFileExt(name)
	// 其他后缀
	otherExt := []string{"exe", "apk", "ipa", "iso", "msi", "dmg"}
	for _, v := range otherExt {
		if v == ext {
			return true
		}
	}
	return false
}
