package forum

// NewUnPost is an action undoing a post that is active
func NewUnPost(poster flon.AccountName, postUUID string) *flon.Action {
	a := &flon.Action{
		Account: ForumAN,
		Name:    ActN("unpost"),
		Authorization: []flon.PermissionLevel{
			{Actor: poster, Permission: flon.PermissionName("active")},
		},
		ActionData: flon.NewActionData(UnPost{
			Poster:   poster,
			PostUUID: postUUID,
		}),
	}
	return a
}

// UnPost represents the `eosio.forum::unpost` action.
type UnPost struct {
	Poster   flon.AccountName `json:"poster"`
	PostUUID string           `json:"post_uuid"`
}
