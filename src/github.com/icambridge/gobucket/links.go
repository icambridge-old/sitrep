package gobucket

type BitbucketAuthor struct {
	Username    string               `json:"username"`
	DisplayName string               `json:"display_name"`
	Links       BitbucketAvatarLinks `json:"links"`
}

type BitbucketAvatarLinks struct {
	Self   BitbucketLink `json:"self"`
	Avatar BitbucketLink `json:"avatar"`
}

type BitbucketLink struct {
	Href string `json:"href"`
}

type SelfLinks struct {
	Self BitbucketLink `json:"self"`
}
