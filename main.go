package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Users struct {
	UserID    int    `json:"UserID"`
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Password  string `json:"Password"`
	Followers []int  `json:"Followers"`
	Following []any  `json:"Following"`
	Posts     []any  `json:"Posts"`
	Login     bool   `json:"Login"`
}

type Posts struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Likes    []int  `json:"likes"`
	Comments []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"comments"`
	UserID int `json:"userId"`
}

func main() {
	// cliForUser(10, 100)
}

func addOrUnLike(userId int, postId int) { // like yoki unlike eventlarni bajarish uchun
	var postStruct []Posts
	var userStruct []Users
	getAllDataFromJson("./users.json", &userStruct)
	getAllDataFromJson("./posts.json", &postStruct)

	var getPost *Posts
	var checkLikes []int

	for idx := range postStruct {
		if postStruct[idx].ID == postId {
			getPost = &postStruct[idx]
			break
		}
	}

	flag := true
	for _, i := range getPost.Likes {
		if i != userId {
			checkLikes = append(checkLikes, i)
		} else {
			flag = false
		}
	}
	if flag {
		checkLikes = append(checkLikes, userId)
	}
	getPost.Likes = checkLikes
	writeToJson("./posts.json", postStruct)
}

func followOrUnFollow(firstUserId int, lastUserId int) { // follow yoki unfollow eventlarni bajarish uchun
	var userStruct []Users
	var getUser *Users
	var checkFollowers []int
	getAllDataFromJson("./users.json", &userStruct)

	for idx := range userStruct {
		if userStruct[idx].UserID == firstUserId {
			getUser = &userStruct[idx]
			break
		}
	}

	flag := true
	for _, i := range getUser.Followers {
		if i != lastUserId {
			checkFollowers = append(checkFollowers, i)
		} else {
			flag = false
		}
	}
	if flag {
		checkFollowers = append(checkFollowers, lastUserId)
	}
	getUser.Followers = checkFollowers
	writeToJson("./users.json", userStruct)
}

func writeToJson(path string, structWriteForJson any) { // json file ichiga yozish uchun
	a, err := json.MarshalIndent(structWriteForJson, " ", "\t")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(path, a, 0644)
}

func getAllDataFromJson(path string, structForData any) { // json file ichidan structga malumot olish uchun
	fl, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	readFile, err := io.ReadAll(fl)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(readFile, &structForData)
}

func cliForUser(commentCount int, likeCount int) { // ekranga clik yaratish uchun
	var choice int
	fmt.Println("======== About Posts ==========")
	for {
		fmt.Printf("[1] Barcha Kommentlar <%d>\n", commentCount)
		fmt.Printf("[2] Barcha Likelar <%d>\n", likeCount)
		fmt.Printf("[3] Chiqish\n")
		fmt.Print("Tanlang: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			getAllComments()
		case 2:
			getAllLikes()
		case 3:
			return
		default:
			fmt.Println("Xato tanlov")
		}
	}
}

func getAllComments() { // postdagi barcha kamentlar ekranga chiqarish uchun
	var postStruct []Posts
	var userStruct []Users
	getAllDataFromJson("./users.json", &userStruct)
	getAllDataFromJson("./posts.json", &postStruct)

	for _, i := range postStruct {
		for _, j := range i.Comments {
			for _, k := range userStruct {
				if k.UserID == i.UserID {
					fmt.Printf("\nUsername: %s\n", k.Username)
					fmt.Printf("Comment: %s\n", j.Name)
					fmt.Println()
				}
			}
		}
	}
}

func getAllLikes() { // postdagi barcha likelar ekranga chiqarish uchun
	var postStruct []Posts
	var userStruct []Users
	getAllDataFromJson("./users.json", &userStruct)
	getAllDataFromJson("./posts.json", &postStruct)

	for _, i := range postStruct {
		for _, j := range i.Likes {
			for _, user := range userStruct {
				if user.UserID == j {
					fmt.Printf("\nPostID: %d\n", i.ID)
					fmt.Printf("Liked by: %s\n", user.Username)
					fmt.Println()
				}
			}
		}
	}
}











