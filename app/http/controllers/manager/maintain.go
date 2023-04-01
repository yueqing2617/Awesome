package manager

import (
	"Awesome/app/http/helper"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"os/exec"
	"time"
)

type Maintain struct {
}

func NewMaintain() *Maintain {
	return &Maintain{
		//Inject services
	}
}

// Backups
func (m *Maintain) Backups(ctx http.Context) {
	backupFileName := facades.Config.Env("DB_CONNECTION").(string) + facades.Config.Env("DB_DATABASE").(string) + time.Now().Format("20060102150405") + ".sql"
	backupFilePath := facades.Config.Env("DB_BACKUP_PATH").(string) + "/" + backupFileName
	command := "mysqldump -h" + facades.Config.Env("DB_HOST").(string) + " -P" + facades.Config.Env("DB_PORT").(string) + " -u" + facades.Config.Env("DB_USERNAME").(string) + " -p" + facades.Config.Env("DB_PASSWORD").(string) + " " + facades.Config.Env("DB_DATABASE").(string) + " > " + backupFilePath
	// 执行命令
	if err := exec.Command("bash", "-c", command).Run(); err != nil {
		helper.RestfulError(ctx, err.Error())
		return
	}
}
