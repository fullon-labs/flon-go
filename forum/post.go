package forum

// NewPost is an action representing a simple message to be posted
// through the chain network.
func NewPost(poster flon.AccountName, postUUID, content string, replyToPoster flon.AccountName, replyToPostUUID string, certify bool, jsonMetadata string) *flon.Action {
	a := &flon.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []flon.PermissionLevel{
			{Actor: poster, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(Post{
			Poster:          poster,
			PostUUID:        postUUID,
			Content:         content,
			ReplyToPoster:   replyToPoster,
			ReplyToPostUUID: replyToPostUUID,
			Certify:         certify,
			JSONMetadata:    jsonMetadata,
		}),
	}
	return a
}

// Post represents the `eosio.forum::post` action.
type Post struct {
	Poster          flon.AccountName `json:"poster"`
	PostUUID        string           `json:"post_uuid"`
	Content         string           `json:"content"`
	ReplyToPoster   flon.AccountName `json:"reply_to_poster"`
	ReplyToPostUUID string           `json:"reply_to_post_uuid"`
	Certify         bool             `json:"certify"`
	JSONMetadata    string           `json:"json_metadata"`
}
