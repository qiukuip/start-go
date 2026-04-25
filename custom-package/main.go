// main.go
package main

import (
	"fmt"
	"custom-package/utils"  // 相对路径导入
)

func main() {
	result := utils.Add(1, 5)
	fmt.Printf("result is: %+v", result)
}

