package gorocket

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type ChannelMessagesRequest struct {
	RoomId string
	Query  string
	Sort   string
	Count  int
	Offset int
}

type ChannelMessage struct {
	ID          string `json:"_id"`
	Alias       string `json:"alias"`
	Msg         string `json:"msg"`
	Attachments []struct {
		Color             string      `json:"color"`
		Text              string      `json:"text"`
		Ts                time.Time   `json:"ts"`
		AuthorName        interface{} `json:"author_name"`
		AuthorLink        interface{} `json:"author_link"`
		AuthorIcon        interface{} `json:"author_icon"`
		Title             interface{} `json:"title"`
		TitleLink         interface{} `json:"title_link"`
		TitleLinkDownload interface{} `json:"title_link_download"`
		Fields            []struct {
			Title string `json:"title"`
			Value string `json:"value"`
			Short bool   `json:"short"`
		} `json:"fields"`
	} `json:"attachments"`
	ParseUrls bool `json:"parseUrls"`
	Bot       struct {
		I string `json:"i"`
	} `json:"bot"`
	Groupable bool      `json:"groupable"`
	Avatar    string    `json:"avatar"`
	Ts        time.Time `json:"ts"`
	U         struct {
		Id       string `json:"_id"`
		Username string `json:"username"`
		Name     string `json:"name"`
	} `json:"u"`
	Rid       string        `json:"rid"`
	Unread    bool          `json:"unread"`
	UpdatedAt time.Time     `json:"_updatedAt"`
	Mentions  []interface{} `json:"mentions"`
	Channels  []interface{} `json:"channels"`
	Md        []struct {
		Type string `json:"type"`
	} `json:"md"`
}

type ChannelMessagesResponse struct {
	Messages []ChannelMessage `json:"messages"`
	Count    int              `json:"count"`
	Offset   int              `json:"offset"`
	Total    int              `json:"total"`
	Success  bool             `json:"success"`
}

// GetChannelMessages retrieves the messages from a channel by a query.
func (c *Client) GetChannelMessages(param *ChannelMessagesRequest) (*ChannelMessagesResponse, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/%s/channels.messages", c.baseURL, c.apiVersion),
		nil)

	if param.RoomId == "" {
		return nil, fmt.Errorf("wrong parameters")
	}

	url := req.URL.Query()
	if param.RoomId != "" {
		url.Add("roomId", param.RoomId)
	}
	if param.Query != "" {
		url.Add("query", param.Query)
	}
	if param.Sort != "" {
		url.Add("sort", param.Sort)
	}
	if param.Count != 0 {
		url.Add("count", strconv.Itoa(param.Count))
	}
	if param.Offset != 0 {
		url.Add("offset", strconv.Itoa(param.Offset))
	}
	req.URL.RawQuery = url.Encode()

	if err != nil {
		return nil, err
	}

	res := ChannelMessagesResponse{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
