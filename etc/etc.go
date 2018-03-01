package etc

import (
	"fmt"
	"strconv"

	"stock/etc/conf"
)

var (
	globalConfPath = `/stock/etc/default.toml`
)

var (
	envDev     bool
	globalConf map[string]interface{}
)

// Dev tells us that current environment is for develop or not.
func Dev() bool {
	return envDev
}

func Int(section, key string) int64 {
	value := getValue(section, key)
	rs, err := strconv.ParseInt(fmt.Sprint(value), 10, 64)
	if err != nil {
		panic("etc: parse bool err:" + err.Error())
	}
	return rs
}

func Bool(section, key string) bool {
	value := getValue(section, key)
	rs, err := strconv.ParseBool(fmt.Sprint(value))
	if err != nil {
		panic("etc: parse bool err:" + err.Error())
	}
	return rs
}

func Float(section, key string) float64 {
	value := getValue(section, key)
	rs, err := strconv.ParseFloat(fmt.Sprint(value), 64)
	if err != nil {
		panic("etc: parse bool err:" + err.Error())
	}
	return rs
}

func String(section, key string) string {
	return fmt.Sprint(getValue(section, key))
}

func getValue(section, key string) interface{} {
	sectionInterface, ok := globalConf[section]
	if !ok {
		panic("etc: section " + section + " not exists.")
	}
	sectionMap, ok := sectionInterface.(map[string]interface{})
	if !ok {
		panic("etc: " + section + " is not a section.")
	}
	if value, ok := sectionMap[key]; ok {
		return value
	}
	panic("etc: can not find key: " + key)
}

func init() {
	globalConf = make(map[string]interface{})
	if err := conf.ReadTOML(globalConfPath, &globalConf); err != nil {
		panic(err)
	}
	envDev = Bool("system", "dev")
}
