package foradmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPackagesTable(ctx *context.Context) (postsTable table.Table) {
	postsTable = table.NewDefaultTable(table.DefaultConfig().SetExportable(true))

	info := postsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Description", "description", db.Text)
	info.AddField("AuthorID", "author_id", db.Int)

	formList := postsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("AuthorID", "author_id", db.Int, form.Number)
	formList.AddField("Git URL", "git_url", db.Varchar, form.Url)

	formList.EnableAjax("Request Success", "Request Failed")

	formList.SetTable("package_models").SetTitle("Package").SetDescription("Package")

	return
}
