/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Package file 上传文件管理
package file

import (
	"time"

	"github.com/webx-top/echo"
	"github.com/webx-top/echo/code"

	"github.com/admpub/nging/application/handler"
	"github.com/admpub/nging/application/model/file"
	"github.com/admpub/nging/application/registry/upload/checker"
	uploadSubdir "github.com/admpub/nging/application/registry/upload/subdir"
)

func setUploadURL(ctx echo.Context) error {
	//uploadType := `nging_user`
	uploadType := ctx.Form(`uploadtype`, `nging_user`)
	params := uploadSubdir.ParseUploadType(uploadType)
	if !params.IsAllowed() {
		return ctx.NewError(code.Failure, ctx.T(`不支持的上传类型: %v`, uploadType))
	}
	ctx.Set(`uploadType`, uploadType)
	ctx.Set(`uploadURL`, checker.BackendUploadURL(uploadType))
	return nil
}

func FileList(ctx echo.Context) error {
	if err := setUploadURL(ctx); err != nil {
		return err
	}
	err := List(ctx, ``, 0)
	ctx.Set(`dialog`, false)
	ctx.Set(`multiple`, true)
	partial := ctx.Formx(`partial`).Bool()
	if partial {
		return ctx.Render(`manager/file/list.main.content`, err)
	}
	return ctx.Render(`manager/file/list`, err)
}

func FileDelete(ctx echo.Context) (err error) {
	user := handler.User(ctx)
	id := ctx.Paramx("id").Uint64()
	fileM := file.NewFile(ctx)
	ownerID := uint64(user.Id)
	if id == 0 {
		ids := ctx.FormxValues(`id`).Uint64()
		for _, id := range ids {
			err = fileM.DeleteByID(id, `user`, ownerID)
			if err != nil {
				return err
			}
		}
		goto END
	}
	err = fileM.DeleteByID(id, `user`, ownerID)
	if err != nil {
		return err
	}

END:
	return ctx.Redirect(handler.URLFor(`/manager/file/list`))
}

// FileClean 删除未使用文件
func FileClean(ctx echo.Context) (err error) {
	fileM := file.NewFile(ctx)
	ago := ctx.Form(`ago`)
	var seconds int64 = 86400 * 365
	if len(ago) > 0 {
		t, e := time.ParseDuration(ago)
		if e != nil {
			return e
		}
		seconds = int64(t.Seconds())
	}
	err = fileM.RemoveUnused(seconds, ``, 0)
	if err != nil {
		return err
	}

	return ctx.Redirect(handler.URLFor(`/manager/file/list`))
}
