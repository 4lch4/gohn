package main

import (
	"fmt"

	"github.com/4lch4/glogger"
	"github.com/go-resty/resty/v2"
	// "github.com/sirupsen/logrus"
)

// #region Internals

// The base URL for v0 of the Hacker News API.
const APIBaseURL = "https://hacker-news.firebaseio.com/v0"

func getUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s.json", APIBaseURL, endpoint)
}

var AppName = "gohn"

var logger = glogger.NewLogger(&glogger.NewLoggerInput{
	AppName: &AppName,
})

var client = resty.New()

// #endregion Internals

// #region Standard Endpoint Methods

func GetItem(id string) (Item, error) {
	var item Item

	_, err := client.R().SetResult(&item).Get(getUrl(fmt.Sprintf("item/%s", id)))
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

func GetUser(username string) (User, error) {
	var user User
	reqUrl := getUrl(fmt.Sprintf("user/%s", username))

	fmt.Println("reqUrl =", reqUrl)

	_, err := client.R().SetResult(&user).Get(reqUrl)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetMaxItemID() (string, error) {
	// resp, err := client.R().Get(fmt.Sprintf("%s/maxitem.json", APIBaseURL))
	resp, err := client.R().Get(getUrl("maxitem"))
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func GetStories(list StoryList) ([]int, error) {
	storyList := string(list)
	// url := fmt.Sprintf("%s/%s.json", APIBaseURL, list)

	var stories []int

	_, err := client.R().SetResult(&stories).Get(getUrl(storyList))
	if err != nil {
		return []int{}, err
	}

	return stories, nil
}

func GetUpdates() (Updates, error) {
	var updates Updates

	_, err := client.R().SetResult(&updates).Get(getUrl("updates"))
	if err != nil {
		return Updates{}, err
	}

	return updates, nil
}

// #endregion Standard Endpoint Methods

func main() {
	res, err := GetStories(TopStories)
	if err != nil {
		fmt.Println(err)
		// logger.Error(err, nil)
	}

	fmt.Println("[main]: Top Story IDs:")
	fmt.Println(res)

	user, err := GetUser("Alcha")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("[main]: User:")
	fmt.Println("About: ", user.About)
	fmt.Println("Created: ", user.Created)
	fmt.Println("ID: ", user.ID)
	fmt.Println("Karma: ", user.Karma)
	fmt.Println("Submitted: ", user.Submitted)
}
