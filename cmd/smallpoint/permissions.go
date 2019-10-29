package main

import (
	"log"
	"sort"
)

var checkPermissionStmt = map[string]string{
	"sqlite":   "select groupname from permissions where (resource=? or resource='*') and resource_type=? and (permission&?=?);",
	"postgres": "select groupname from permissions where (resource=$1 or resource='*') and resource_type=$2 and (permission&$3=$4);",
}

const (
	permCreate = 1 << iota
	permUpdate
	permDelete
)

const (
	resourceGroup = 1 << iota
	resourceSVC
)

func checkPermission(resources string, resource_type, permission int, state *RuntimeState) []string {
	stmtText := checkPermissionStmt[state.dbType]
	stmt, err := state.db.Prepare(stmtText)
	if err != nil {
		log.Println("Error prepare statement " + stmtText)
	}
	defer stmt.Close()

	var groupnames []string
	rows, err := stmt.Query(resources, resource_type, permission, permission)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			log.Println("No rows found")
			return nil
		} else {
			log.Println(err)
			return nil

		}
	}
	defer rows.Close()

	for rows.Next() {
		var groupName string
		err = rows.Scan(&groupName)
		groupnames = append(groupnames, groupName)
	}
	return groupnames
}

func (state *RuntimeState) canPerformAction(username, resources string, resource_type, permission int) (bool, error) {
	groups := checkPermission(resources, resource_type, permission, state)
	if len(groups) < 1 {
		return false, nil
	}
	groupsOfUser, err := state.Userinfo.GetgroupsofUser(username)
	if err != nil {
		return false, err
	}
	sort.Strings(groupsOfUser)

	adminGroup := state.Config.TargetLDAP.AdminGroup
	adminIndex := sort.SearchStrings(groupsOfUser, adminGroup)
	if adminIndex < len(groupsOfUser) && groupsOfUser[adminIndex] == adminGroup {
		return true, nil
	}

	for _, group := range groups {
		var index int
		index = sort.SearchStrings(groupsOfUser, group)
		if index < len(groupsOfUser) && groupsOfUser[index] == group {
			return true, nil
		}
		continue
	}

	return false, nil
}
