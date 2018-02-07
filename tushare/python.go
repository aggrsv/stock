package tushare

import (
	"fmt"

	python "github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

var PyStr = python.PyString_FromString
var GoStr = python.PyString_AS_STRING

func Tushare() error {
	// import stock.py
	InsertBeforeSysPath("/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/site-packages")
	stock := ImportModule("/Users/youmy/go/src/stock/tushare", "stock")

	basics := stock.GetAttrString("stock_basics")
	bArgs := python.PyTuple_New(1)
	python.PyTuple_SetItem(bArgs, 0, PyStr("sybmol"))

	res := basics.Call(bArgs, python.Py_None)
	fmt.Printf("[CALL] tushare('sybmol') = %s\n", GoStr(res))
	fmt.Println(res.Type())
	fmt.Println(res.Repr())
	//fmt.Println(res.Bytes())
	return nil
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


                                                                                             