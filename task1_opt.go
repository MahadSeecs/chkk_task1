package main

import "fmt"

/* this func splits the relative path into its components by removing /, this will make them as stacks as end elements
will be removed*/
func splitPath(path string) []string {
	var objects []string 
	var obj string
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if obj != "" {
				objects = append(objects, obj)
				obj = ""
			}
		} else {
			obj += string(path[i])
		}
	}
	if obj != "" {
		objects = append(objects, obj)
	}
	return objects
}

/*here the cleaned and split paths will be compared. any .. found in relParts will lead to popping the currParts.
for ., it remain the same and be appended*/
func GetAbsolutePath(currentDir, relativePath string) string {
	currParts := splitPath(currentDir)
	relParts := splitPath(relativePath)

	for i := 0; i < len(relParts); i++ {
		if relParts[i] == ".." {
			if len(currParts) > 0 {
				currParts = currParts[:len(currParts)-1]
			}
		} else if relParts[i] != "." {
			currParts = append(currParts, relParts[i])
		}
	}

	result := "/"
	for i := 0; i < len(currParts); i++ {
		if i > 0 {
			result += "/"
		}
		result += currParts[i]
	}
	return result
}

func main() {
	currentDir := "/home///folder1"
	relativePath := "../../folder2//hi.txt"
	absolutePath := GetAbsolutePath(currentDir, relativePath)
	fmt.Println("Absolute Path:", absolutePath)

}
