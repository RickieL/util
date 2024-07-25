package sql

import "strings"

// SplitSQL 分割 sql 语句
func SplitSQL(sql string) []string {
	var (
		startIndex int
		sqls       []string
	)
	for i := 0; i < len(sql); i++ {
		// 判断 comment
		if sql[i] == '\'' {
			i += strings.Index(sql[i+1:], "'") + 1
		} else if sql[i] == '"' {
			i += strings.Index(sql[i+1:], "\"") + 1
		} else if sql[i] == '`' {
			i += strings.Index(sql[i+1:], "`") + 1
			// 判断注释
		} else if sql[i] == '/' && i+1 < len(sql) && sql[i+1] == '*' {
			i += strings.Index(sql[i+1:], "*/") + 1
		} else if sql[i] == '-' && i+1 < len(sql) && sql[i+1] == '-' {
			i += strings.Index(sql[i+1:], "\n") + 1
		} else if sql[i] == '#' {
			i += strings.Index(sql[i+1:], "\n") + 1
		} else if sql[i] == ';' {
			sqls = append(sqls, strings.TrimSpace(sql[startIndex:i+1]))
			startIndex = i + 1
		}
	}
	// 防止最后一个 sql 不以分号结尾
	if startIndex != len(sql) {
		if strings.TrimSpace(sql[startIndex:]) != "" {
			sqls = append(sqls, strings.TrimSpace(sql[startIndex:]))
		}
	}
	return sqls
}
