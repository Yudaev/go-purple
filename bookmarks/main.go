package main

import "fmt"

type bookmarkMap = map[string]string

func main() {
	bookmarks := bookmarkMap{}
	fmt.Println("Приложение для закладок")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выбери вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Print("Введи номер: ")
	_, _ = fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		println(key, ": ", value)
	}
}

func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Введите название: ")
	_, _ = fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите ссылку: ")
	_, _ = fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBookmark(bookmarks bookmarkMap) {
	var bookmarkKeyToDelete string
	fmt.Println("Введите название: ")
	_, _ = fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}
