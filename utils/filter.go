package utils

import (
	"strings"
)

type (
	// Filter Represents a filter
	// (which will allow to specify the data the user want to replicate or not on the database, schema, table or column level.
	Filter struct {
		FilterPolicy string                 `json:"filter_policy" mapstructure:"filter_policy"`
		Filter       map[string]interface{} `json:"filter"`
	}
)

// isAccept returns true if policy is accept
func (f *Filter) isAccept() bool {
	return f.FilterPolicy == "accept"
}

// IsFilteredDatabase check if database is filtered
func (f *Filter) IsFilteredDatabase(database string) bool {
	found := false
	if tableFilter, ok := f.Filter[strings.ToLower(database)]; ok {
		if tableFilter == nil {
			found = true
		} else {
			return false
		}
	}
	if found {
		return f.isAccept()
	}
	return !f.isAccept()
}

// IsFilteredTable check if table is filtered
func (f *Filter) IsFilteredTable(database string, table string) bool {
	found := false
	if tableFilter, ok := f.Filter[strings.ToLower(database)].(map[string]interface{}); ok {
		if columnFilter, ok := tableFilter[strings.ToLower(table)]; ok {
			if columnFilter == nil {
				found = true
			} else {
				return false
			}
		}
	}
	if found {
		return f.isAccept()
	}
	return !f.isAccept()
}

// IsFilteredColumn check if column is filtered
func (f *Filter) IsFilteredColumn(database string, table string, column string) bool {
	found := false
	if tableFilter, ok := f.Filter[strings.ToLower(database)].(map[string]interface{}); ok {
		columnFilter := tableFilter[strings.ToLower(table)]
		if columnFilter != nil {
			found = contains(columnFilter.([]interface{}), column)
		}
	}
	if found {
		return f.isAccept()
	}
	return !f.isAccept()
}

// contains check if e is contain in s
func contains(s []interface{}, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
