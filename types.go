package main

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

// Item struct represents a single item returned by the HN API.
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

// User struct represents a single user returned by the HN API.
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

// Updates struct represents the response from the HN API for the updates endpoint.
type Updates struct {
	// An array of item ids for items that have changed.
	Items []int `json:"items"`

	// An array of usernames for profiles that have changed.
	Profiles []string `json:"profiles"`
}
