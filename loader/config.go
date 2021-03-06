package loader

import (
	"reflect"
	"sync"
)

//Config assembler config struct
type Config struct {
	//Checkers registered type checkers
	Checkers *TypeCheckers
	//Unifiers registered unifiers
	Unifiers *Unifiers
	//TagName tag name used when parsed.
	//Tag will not be parsed if set to empty string
	//Default value is config
	TagName string
	//TagLazyLoad tag for lazyload
	//Default value is lazyload
	TagLazyLoad string
	//TagAnonymous tag for anonymous
	//Default value is anonymous
	TagAnonymous string
	//TagParser func which parses tags with given config
	TagParser func(c *Config, value string) (*Tag, error)
	//CaseSensitive convert struct field in case sensitive mode.
	CaseSensitive bool
	//DisableConvertStringInterface disable convert String interface to string field
	DisableConvertStringInterface bool
	//DisableConvertNumber disable numver converting
	DisableConvertNumber bool
	//CachedTags cached struct field tags
	CachedTags sync.Map
}

//GetTag get tags for given reflect type and struct field.
//Return tag and any error if raised
func (c *Config) GetTag(structType reflect.Type, field reflect.StructField) (*Tag, error) {
	if c.TagName == "" {
		return nil, nil
	}
	sname := structType.Name()
	if sname != "" {
		key := structType.PkgPath() + "." + structType.Name() + "." + field.Name
		t, ok := c.CachedTags.Load(key)
		if ok {
			return t.(*Tag), nil
		}
		tag, err := c.TagParser(c, field.Tag.Get(c.TagName))
		if err != nil {
			return nil, err
		}
		c.CachedTags.Store(key, tag)
		return tag, nil
	}
	return c.TagParser(c, field.Tag.Get(c.TagName))
}

//NewConfig create new config.
func NewConfig() *Config {
	return &Config{
		TagParser:    ParseTag,
		Unifiers:     NewUnifiers(),
		Checkers:     NewTypeCheckers(),
		TagLazyLoad:  "lazyload",
		TagAnonymous: "anonymous",
	}
}

//NewCommonConfig create new common config
func NewCommonConfig() *Config {
	c := NewConfig()
	c.TagName = "config"
	c.Checkers.Append(CommonTypeCheckers)
	c.Unifiers.Append(CommonUnifiers)
	return c
}

func InitCommon() {
	CommonUnifiers = DefaultCommonUnifiers()
	CommonTypeCheckers = DefaultCommonTypeCheckers()
}
