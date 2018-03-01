package tushare

import (
	"errors"
	"stock/etc"

	python "github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

var (
	sitPkg = etc.String("py_path", "sitepkg")
)

var PyStr = python.PyString_FromString
var GoStr = python.PyString_AS_STRING

func GetProfit(year, season string) string {
	//InsertBeforeSysPath("")
	stock := ImportModule("", "stock")
	basics := stock.GetAttrString("profit")
	//defer basics.Clear()

	bArgs := python.PyTuple_New(1)
	python.PyTuple_SetItem(bArgs, 0, PyStr(year))
	python.PyTuple_SetItem(bArgs, 1, PyStr(season))
	res := basics.Call(bArgs, python.Py_None)

	return GoStr(res)
}

func Tushare() (string, error) {
	// import stock.py
	InsertBeforeSysPath(sitPkg)
	stock := ImportModule("tushare", "stock")
	if stock == nil {
		return "", errors.New("stock == nil")
	}
	basics := stock.GetAttrString("stock_basics")
	if basics == nil {
		return "", errors.New("basic == nil")
	}
	bArgs := python.PyTuple_New(1)
	python.PyTuple_SetItem(bArgs, 0, PyStr("sybmol"))

	res := basics.Call(bArgs, python.Py_None)

	return GoStr(res), nil
}

// InsertBeforeSysPath will add given dir to python import path
func InsertBeforeSysPath(p string) string {
	sysModule := python.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")
	python.PyList_Insert(path, 0, PyStr(p))
	return GoStr(path.Repr())
}

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python.PyObject {
	sysModule := python.PyImport_ImportModule("sys") // import sys
	path := sysModule.GetAttrString("path")          // path = sys.path
	python.PyList_Insert(path, 0, PyStr(dir))        // path.insert(0, dir)
	return python.PyImport_ImportModule(name)        // return __import__(name)
}
