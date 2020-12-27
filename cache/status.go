/*******************************************************************************
Method: 缓存状态
Author: Lemine
Langua: golang v1.14
Modify: 2020-12-12
*******************************************************************************/
package cache

import (
	"cacgo/utils"
)

type Status struct {
	Count    int64 //键值对数量
	KeyBytes int64 //键的总字节数
	ValBytes int64 //值的总字节数
	NowBytes int64 //缓存已使用的字节数
	MaxBytes int64 //缓存可使用的最大字节数

	UsageRate float64 //使用率(NowBytes/MaxBytes)
	UtiliRate float64 //利用率((ValBytes+KeyBytes)/NowBtyes)
}

func (self *Status) update() {
	self.NowBytes = 2*self.KeyBytes + self.ValBytes
	self.UsageRate = float64(self.NowBytes) / float64(self.MaxBytes)
	self.UtiliRate = float64(self.KeyBytes+self.ValBytes) / float64(self.NowBytes)
}

func (self *Status) Get() Status {
	return Status{
		Count:     self.Count,
		KeyBytes:  self.KeyBytes,
		ValBytes:  self.ValBytes,
		NowBytes:  self.NowBytes,
		MaxBytes:  self.MaxBytes,
		UsageRate: self.UsageRate,
		UtiliRate: self.UtiliRate,
	}
}

func (self *Status) Format() map[string]interface{} {
	return map[string]interface{}{
		"Count":     self.Count,
		"KeyBytes":  utils.FormatSizeInt64(self.KeyBytes),
		"ValBytes":  utils.FormatSizeInt64(self.ValBytes),
		"NowBytes":  utils.FormatSizeInt64(self.NowBytes),
		"MaxBytes":  utils.FormatSizeInt64(self.NowBytes),
		"UsageRate": utils.FormatPercentFloat64(self.UsageRate),
		"UiltiRate": utils.FormatPercentFloat64(self.UtiliRate),
	}
}
