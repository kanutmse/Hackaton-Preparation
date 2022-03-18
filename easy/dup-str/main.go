package main

/*
First, we store each character's last index appeared in s. (= hashMap)

Then we have an array res and iterate through s again and perform this:

Push a character s[i] to res if res is empty.
If res is non-empty, check res's last element.
If res[last] is greater than s[i] and hashmap[res[last]] >= i, which means res[last] appears again in s during following iteration, so to keep res small in lexicographical order, we can discard res[last]. After discarding, again we repeat this step 2 until we don't met the condition.
But if we only follow the steps above, we will pop all characters if we have res = "abcde" and say s[i] = 'a'. We need to manage so that once the lexicographical order is fixed, we won't touch the same character again. In this purpose, we have visited map.
And once it's added we record true in the map and false when popped.


*/
func removeDuplicateLetters(s string) string {

	hashMap := [26]int{}
	visited := [26]bool{}

	for i := range s {
		hashMap[int(s[i])-int('a')] = i
	}

	var res []byte
	for i := range s {
		ch := s[i]
		for len(res) > 0 {
			lastCh := res[len(res)-1]
			if ch <= lastCh && hashMap[int(lastCh)-int('a')] >= i && !visited[int(ch)-int('a')] {
				visited[int(lastCh)-int('a')] = false
				res = res[:len(res)-1]
				continue
			}
			break
		}

		if !visited[int(ch)-int('a')] {
			res = append(res, ch)
			visited[int(ch)-int('a')] = true
		}
	}
	return string(res)
}
