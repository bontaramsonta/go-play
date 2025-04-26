package main

import "slices"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	directFriends := friendships[username]
	ignoreList := append([]string{username}, directFriends...)
	fm := make(map[string]bool)
	for _, df := range directFriends {
		fodf := friendships[df]
		for _, sf := range fodf {
			if !slices.Contains(ignoreList, sf) {
				fm[sf] = true
			}
		}
	}
	var sflist []string
	for k := range fm {
		sflist = append(sflist, k)
	}
	return sflist
}
