// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
	
	"time"
)

type Slice_FtpUserGroup []*FtpUserGroup

func (s Slice_FtpUserGroup) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_FtpUserGroup) RangeRaw(fn func(m *FtpUserGroup) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// FtpUserGroup FTP用户组
type FtpUserGroup struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FtpUserGroup
	namer   func(string) string
	connID  int
	context echo.Context
	
	Id          	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Name        	string  	`db:"name" bson:"name" comment:"组名称" json:"name" xml:"name"`
	Created     	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	Disabled    	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Banned      	string  	`db:"banned" bson:"banned" comment:"是否禁止组内用户连接" json:"banned" xml:"banned"`
	Directory   	string  	`db:"directory" bson:"directory" comment:"授权目录" json:"directory" xml:"directory"`
	IpWhitelist 	string  	`db:"ip_whitelist" bson:"ip_whitelist" comment:"IP白名单(一行一个)" json:"ip_whitelist" xml:"ip_whitelist"`
	IpBlacklist 	string  	`db:"ip_blacklist" bson:"ip_blacklist" comment:"IP黑名单(一行一个)" json:"ip_blacklist" xml:"ip_blacklist"`
}

func (this *FtpUserGroup) Trans() *factory.Transaction {
	return this.trans
}

func (this *FtpUserGroup) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FtpUserGroup) SetContext(ctx echo.Context) factory.Model {
	this.context = ctx
	return this
}

func (this *FtpUserGroup) Context() echo.Context {
	return this.context
}

func (this *FtpUserGroup) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *FtpUserGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *FtpUserGroup) Objects() []*FtpUserGroup {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FtpUserGroup) NewObjects() factory.Ranger {
	return &Slice_FtpUserGroup{}
}

func (this *FtpUserGroup) InitObjects() *[]*FtpUserGroup {
	this.objects = []*FtpUserGroup{}
	return &this.objects
}

func (this *FtpUserGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *FtpUserGroup) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *FtpUserGroup) Short_() string {
	return "ftp_user_group"
}

func (this *FtpUserGroup) Struct_() string {
	return "FtpUserGroup"
}

func (this *FtpUserGroup) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *FtpUserGroup) Namer() func(string) string {
	return this.namer
}

func (this *FtpUserGroup) CPAFrom(source factory.Model) factory.Model {
	this.SetContext(source.Context())
	this.Use(source.Trans())
	this.SetNamer(source.Namer())
	return this
}

func (this *FtpUserGroup) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FtpUserGroup) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FtpUserGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FtpUserGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FtpUserGroup) GroupBy(keyField string, inputRows ...[]*FtpUserGroup) map[string][]*FtpUserGroup {
	var rows []*FtpUserGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*FtpUserGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*FtpUserGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *FtpUserGroup) KeyBy(keyField string, inputRows ...[]*FtpUserGroup) map[string]*FtpUserGroup {
	var rows []*FtpUserGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*FtpUserGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *FtpUserGroup) AsKV(keyField string, valueField string, inputRows ...[]*FtpUserGroup) map[string]interface{} {
	var rows []*FtpUserGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *FtpUserGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FtpUserGroup) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Banned) == 0 { this.Banned = "N" }
	err = DBI.Fire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", this, nil)
	}
	return
}

func (this *FtpUserGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Banned) == 0 { this.Banned = "N" }
	if err = DBI.Fire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", this, mw, args...)
}

func (this *FtpUserGroup) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *FtpUserGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *FtpUserGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	if val, ok := kvset["banned"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["banned"] = "N" } }
	m := *this
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (this *FtpUserGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Banned) == 0 { this.Banned = "N" }
		return DBI.Fire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Banned) == 0 { this.Banned = "N" }
		return DBI.Fire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.Fire("updated", this, mw, args...)
		} else {
			err = DBI.Fire("created", this, nil)
		}
	} 
	return 
}

func (this *FtpUserGroup) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.Fire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", this, mw, args...)
}

func (this *FtpUserGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *FtpUserGroup) Reset() *FtpUserGroup {
	this.Id = 0
	this.Name = ``
	this.Created = 0
	this.Updated = 0
	this.Disabled = ``
	this.Banned = ``
	this.Directory = ``
	this.IpWhitelist = ``
	this.IpBlacklist = ``
	return this
}

func (this *FtpUserGroup) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["Disabled"] = this.Disabled
	r["Banned"] = this.Banned
	r["Directory"] = this.Directory
	r["IpWhitelist"] = this.IpWhitelist
	r["IpBlacklist"] = this.IpBlacklist
	return r
}

func (this *FtpUserGroup) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
			case "id": this.Id = param.AsUint(value)
			case "name": this.Name = param.AsString(value)
			case "created": this.Created = param.AsUint(value)
			case "updated": this.Updated = param.AsUint(value)
			case "disabled": this.Disabled = param.AsString(value)
			case "banned": this.Banned = param.AsString(value)
			case "directory": this.Directory = param.AsString(value)
			case "ip_whitelist": this.IpWhitelist = param.AsString(value)
			case "ip_blacklist": this.IpBlacklist = param.AsString(value)
		}
	}
}

func (this *FtpUserGroup) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Name": this.Name = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
				case "Disabled": this.Disabled = param.AsString(vv)
				case "Banned": this.Banned = param.AsString(vv)
				case "Directory": this.Directory = param.AsString(vv)
				case "IpWhitelist": this.IpWhitelist = param.AsString(vv)
				case "IpBlacklist": this.IpBlacklist = param.AsString(vv)
			}
	}
}

func (this *FtpUserGroup) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["name"] = this.Name
	r["created"] = this.Created
	r["updated"] = this.Updated
	r["disabled"] = this.Disabled
	r["banned"] = this.Banned
	r["directory"] = this.Directory
	r["ip_whitelist"] = this.IpWhitelist
	r["ip_blacklist"] = this.IpBlacklist
	return r
}

func (this *FtpUserGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *FtpUserGroup) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

