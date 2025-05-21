package api

import (
	"fmt"
)

func GetSupergroupUsernamesByID(id int64) ([]string, error) {
	var usernames []string

	item, err := GetSupergroup(id)
	if err != nil {
		return nil, fmt.Errorf("tdlibClient.GetSupergroup(%d) %v", id, err)
	}
	uMap := make(map[string]string, 0)
	if item != nil && item.Usernames != nil {
		for _, v := range item.Usernames.ActiveUsernames {
			if _, ok := uMap[v]; !ok {
				uMap[v] = v
				usernames = append(usernames, v)
			}
		}
		//for _, v := range item.Usernames.DisabledUsernames {
		//	if _, ok := uMap[v]; !ok {
		//		uMap[v] = v
		//		usernames = append(usernames, v)
		//	}
		//}
		if len(item.Usernames.EditableUsername) > 0 {
			v := item.Usernames.EditableUsername
			if _, ok := uMap[v]; !ok {
				uMap[v] = v
				usernames = append(usernames, v)
			}
		}
	}
	if len(usernames) > 0 {
		return usernames, nil
	}
	return nil, fmt.Errorf("the Supergroup ID(%d) have no username", id)
}
