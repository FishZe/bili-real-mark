package mark

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

// GetNum2Letter
//
//	@Description: 获取数字对应的字母
//	@param num
//	@return string
func GetNum2Letter(num int) string {
	var res string
	for num != 0 {
		if num%26 == 0 {
			res = "Z" + res
			num = num/26 - 1
		} else {
			res = string(rune(num%26+64)) + res
			num = num / 26
		}
	}
	return res
}

// Write2Excel
//
//	@Description: 写入Excel
//	@param fileName 文件名
//	@param data 数据
//	@param key 每个sheet的key的顺序
//	@return error
func Write2Excel(fileName string, data map[string]map[string][]string, key map[string][]string) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("close file failed: %v", err)
		}
	}()
	for i, v := range data {
		// 遍历sheet
		if now, err := f.NewSheet(i); err == nil {
			index := 1
			for _, k := range key[i] {
				// 对这个表的表头遍历
				_ = f.SetCellValue(i, GetNum2Letter(index)+"1", k)
				for l, m := range v[k] {
					_ = f.SetCellValue(i, GetNum2Letter(index)+strconv.Itoa(l+2), m)
				}
				index++
			}
			f.SetActiveSheet(now)
		}
	}
	if err := f.SaveAs(fileName); err != nil {
		return err
	}
	return nil
}
