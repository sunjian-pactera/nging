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

type Slice_FrpGroup []*FrpGroup

func (s Slice_FrpGroup) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_FrpGroup) RangeRaw(fn func(m *FrpGroup) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// FrpGroup FRP服务组
type FrpGroup struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FrpGroup
	namer   func(string) string
	connID  int
	context echo.Context
	
	Id         	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid        	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Name       	string  	`db:"name" bson:"name" comment:"组名" json:"name" xml:"name"`
	Description	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated    	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
}

func (this *FrpGroup) Trans() *factory.Transaction {
	return this.trans
}

func (this *FrpGroup) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FrpGroup) SetContext(ctx echo.Context) factory.Model {
	this.context = ctx
	return this
}

func (this *FrpGroup) Context() echo.Context {
	return this.context
}

func (this *FrpGroup) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *FrpGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *FrpGroup) Objects() []*FrpGroup {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FrpGroup) NewObjects() factory.Ranger {
	return &Slice_FrpGroup{}
}

func (this *FrpGroup) InitObjects() *[]*FrpGroup {
	this.objects = []*FrpGroup{}
	return &this.objects
}

func (this *FrpGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *FrpGroup) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *FrpGroup) Short_() string {
	return "frp_group"
}

func (this *FrpGroup) Struct_() string {
	return "FrpGroup"
}

func (this *FrpGroup) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *FrpGroup) Namer() func(string) string {
	return this.namer
}

func (this *FrpGroup) CPAFrom(source factory.Model) factory.Model {
	this.SetContext(source.Context())
	this.Use(source.Trans())
	this.SetNamer(source.Namer())
	return this
}

func (this *FrpGroup) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FrpGroup) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FrpGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FrpGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FrpGroup) GroupBy(keyField string, inputRows ...[]*FrpGroup) map[string][]*FrpGroup {
	var rows []*FrpGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*FrpGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*FrpGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *FrpGroup) KeyBy(keyField string, inputRows ...[]*FrpGroup) map[string]*FrpGroup {
	var rows []*FrpGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*FrpGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *FrpGroup) AsKV(keyField string, valueField string, inputRows ...[]*FrpGroup) map[string]interface{} {
	var rows []*FrpGroup
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

func (this *FrpGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FrpGroup) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
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

func (this *FrpGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	this.Updated = uint(time.Now().Unix())
	if err = DBI.Fire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", this, mw, args...)
}

func (this *FrpGroup) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *FrpGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *FrpGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	kvset["updated"] = uint(time.Now().Unix())
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

func (this *FrpGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { this.Updated = uint(time.Now().Unix())
		return DBI.Fire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
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

func (this *FrpGroup) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.Fire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", this, mw, args...)
}

func (this *FrpGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *FrpGroup) Reset() *FrpGroup {
	this.Id = 0
	this.Uid = 0
	this.Name = ``
	this.Description = ``
	this.Created = 0
	this.Updated = 0
	return this
}

func (this *FrpGroup) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["Name"] = this.Name
	r["Description"] = this.Description
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	return r
}

func (this *FrpGroup) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
			case "id": this.Id = param.AsUint(value)
			case "uid": this.Uid = param.AsUint(value)
			case "name": this.Name = param.AsString(value)
			case "description": this.Description = param.AsString(value)
			case "created": this.Created = param.AsUint(value)
			case "updated": this.Updated = param.AsUint(value)
		}
	}
}

func (this *FrpGroup) Set(key interface{}, value ...interface{}) {
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
				case "Uid": this.Uid = param.AsUint(vv)
				case "Name": this.Name = param.AsString(vv)
				case "Description": this.Description = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsUint(vv)
			}
	}
}

func (this *FrpGroup) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["uid"] = this.Uid
	r["name"] = this.Name
	r["description"] = this.Description
	r["created"] = this.Created
	r["updated"] = this.Updated
	return r
}

func (this *FrpGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *FrpGroup) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

