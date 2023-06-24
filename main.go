package main

import (
	"fmt"

	"github.com/4lch4/glogger"
	"github.com/go-resty/resty/v2"
	// "github.com/sirupsen/logrus"
)

// #region Structs

// A string representing the possible values for the Type field of an Item.
type ItemType string // "job", "story", "comment", "poll", or "pollopt"

// The possible values for the Type field of an Item.
const (
	Job     ItemType = "job"
	Story   ItemType = "story"
	Comment ItemType = "comment"
	Poll    ItemType = "poll"
	PollOpt ItemType = "pollopt"
)

type StoryList string

const (
	TopStories  StoryList = "topstories"
	NewStories  StoryList = "newstories"
	BestStories StoryList = "beststories"
	AskStories  StoryList = "askstories"
	ShowStories StoryList = "showstories"
	JobStories  StoryList = "jobstories"
)

type Item struct {
	// The item's unique id (this is the only required property).
	ID string `json:"id"`

	// `true` if the item is deleted.
	Deleted bool `json:"deleted"`

	// A string representing the type of item. One of: "job", "story", "comment", "poll", or "pollopt"
	Type ItemType `json:"type"`

	// The username of the item's author.
	By string `json:"by"`

	// Creation date of the item, in Unix Time.
	Time string `json:"time"`

	// The comment, story or poll text. HTML.
	Text string `json:"text"`

	// `true` if the item is dead.
	Dead bool `json:"dead"`

	// The comment's parent: either another comment or the relevant story.
	Parent string `json:"parent"`

	// The pollopt's associated poll.
	Poll string `json:"poll"`

	// The ids of the item's comments, in ranked display order.
	Kids []string `json:"kids"`

	// The URL of the story.
	URL string `json:"url"`

	// The story's score, or the votes for a pollopt.
	Score string `json:"score"`

	// The title of the story, poll or job. HTML.
	Title string `json:"title"`

	// A list of related pollopts, in display order.
	Parts []string `json:"parts"`

	// In the case of stories or polls, the total comment count.
	Descendants string `json:"descendants"`
}

type User struct {
	// The user's unique username. Case-sensitive. Required.
	ID string `json:"id"`

	// Creation date of the user, in Unix Time.
	Created int `json:"created"`

	// The user's karma.
	Karma int `json:"karma"`

	// The user's optional self-description. HTML.
	About string `json:"about"`

	// List of the user's stories, polls and comments.
	Submitted []int `json:"submitted"`
}

type Updates struct {
	// An array of item ids for items that have changed.
	Items []int `json:"items"`

	// An array of usernames for profiles that have changed.
	Profiles []string `json:"profiles"`
}

// #endregion Structs

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
