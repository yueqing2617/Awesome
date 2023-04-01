package helper

import "os"

// GetEnv 获取根目录下 .env 文件中的环境变量
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// SetEnv 设置 .env 文件中的环境变量
func SetEnv(key string, value string) error {
	err := os.Setenv(key, value)
	if err != nil {
		return err
	}
	return nil
}

// HasEnv 判断 .env 文件中是否存在环境变量
func HasEnv(key string) bool {
	value := os.Getenv(key)
	if value == "" {
		return false
	}
	return true
}
