package packages

import (
	"errors"

	"github.com/zueffc/mirage-node/internal/data"
	"github.com/zueffc/mirage-node/pkg/checker"
)

func Create(name, description, gitUrl string, authorId uint) error {
	if checker.SizeCheck(name, 3, 128) && checker.SizeCheck(description, 10, 16128) && checker.SizeCheck(gitUrl, 21, 255) {
		newPackage := data.PackageModel{
			Package: data.Package{
				Name:        name,
				Description: description,
				AuthorID:    authorId,
				GitURL:      gitUrl,
			},
		}

		data.Database().Create(&newPackage)
		return nil
	}

	return errors.New("SIZE IS UNCOMPATIBLE")
}

func FindByID(queryType string, id uint) *[]data.PackageModel {
	if queryType == "all" {
		var packages []data.PackageModel

		data.Database().Where(&data.PackageModel{
			Package: data.Package{
				AuthorID: id,
			},
		}).Find(&packages)

		return &packages
	} else {
		return nil
	}
}

func FindByName(name string) *data.PackageModel {
	var pkg data.PackageModel

	data.Database().Where(&data.PackageModel{
		Package: data.Package{
			Name: name,
		},
	}).Find(&pkg)

	return &pkg
}

func Delete() {}
func Edit()   {}
