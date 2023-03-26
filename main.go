package main

import (
	"fmt"
	"github.com/goravel/framework/facades"

	"Awesome/bootstrap"
)

func main() {
	agpc()
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Start http server by facades.Route.
	go func() {
		if err := facades.Route.Run(); err != nil {
			facades.Log.Errorf("Route run error: %v", err)
		}
	}()

	select {}
}

/***
 *
 *         /\
 *        /  \    __      __   ___   ___    ___    _ __ ___     ___
 *       / /\ \   \ \ /\ / /  / _ \ / __|  / _ \  | '_ ` _ \   / _ \
 *      / ____ \   \ V  V /  |  __/ \__ \ | (_) | | | | | | | |  __/
 *     /_/    \_\   \_/\_/    \___| |___/  \___/  |_| |_| |_|  \___|
 *
 *
 */
// agpc is a function that prints the ASCII art of the framework.
func agpc() {
	height := 9
	art := []string{
		"         /\\",
		"        /  \\    __      __   ___   ___    ___    _ __ ___     ___",
		"       / /\\ \\   \\ \\ /\\ / /  / _ \\ / __|  / _ \\  | '_ ` _ \\   / _ \\",
		"      / ____ \\   \\ V  V /  |  __/ \\__ \\ | (_) | | | | | | | |  __/",
		"     /_/    \\_\\   \\_/\\_/    \\___| |___/  \\___/  |_| |_| |_|  \\___|",
		"						- 服装生产ERP系统",
		"						- 作者：周东明",
		"						- 邮箱：empty@inzj.cn",
		"						- QQ：  804966813",
	}
	for i := 0; i < height; i++ {
		if i < len(art) {
			fmt.Println(art[i])
		} else {
			fmt.Println()
		}
	}
}
