package payload

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/lowply/webhook/logger"
)

type Payload struct {
	After   string      `json:"after"`
	BaseRef interface{} `json:"base_ref"`
	Before  string      `json:"before"`
	Commits []struct {
		Added  []string `json:"added"`
		Author struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"committer"`
		Distinct  bool          `json:"distinct"`
		ID        string        `json:"id"`
		Message   string        `json:"message"`
		Modified  []string      `json:"modified"`
		Removed   []interface{} `json:"removed"`
		Timestamp string        `json:"timestamp"`
		URL       string        `json:"url"`
	} `json:"commits"`
	Compare    string `json:"compare"`
	Created    bool   `json:"created"`
	Deleted    bool   `json:"deleted"`
	Forced     bool   `json:"forced"`
	HeadCommit struct {
		Added  []string `json:"added"`
		Author struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"committer"`
		Distinct  bool          `json:"distinct"`
		ID        string        `json:"id"`
		Message   string        `json:"message"`
		Modified  []string      `json:"modified"`
		Removed   []interface{} `json:"removed"`
		Timestamp string        `json:"timestamp"`
		URL       string        `json:"url"`
	} `json:"head_commit"`
	Pusher struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"pusher"`
	Ref        string `json:"ref"`
	Repository struct {
		ArchiveURL       string      `json:"archive_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BlobsURL         string      `json:"blobs_url"`
		BranchesURL      string      `json:"branches_url"`
		CloneURL         string      `json:"clone_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		CommentsURL      string      `json:"comments_url"`
		CommitsURL       string      `json:"commits_url"`
		CompareURL       string      `json:"compare_url"`
		ContentsURL      string      `json:"contents_url"`
		ContributorsURL  string      `json:"contributors_url"`
		CreatedAt        float64     `json:"created_at"`
		DefaultBranch    string      `json:"default_branch"`
		Description      string      `json:"description"`
		DownloadsURL     string      `json:"downloads_url"`
		EventsURL        string      `json:"events_url"`
		Fork             bool        `json:"fork"`
		Forks            float64     `json:"forks"`
		ForksCount       float64     `json:"forks_count"`
		ForksURL         string      `json:"forks_url"`
		FullName         string      `json:"full_name"`
		GitCommitsURL    string      `json:"git_commits_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitURL           string      `json:"git_url"`
		HasDownloads     bool        `json:"has_downloads"`
		HasIssues        bool        `json:"has_issues"`
		HasPages         bool        `json:"has_pages"`
		HasWiki          bool        `json:"has_wiki"`
		Homepage         interface{} `json:"homepage"`
		HooksURL         string      `json:"hooks_url"`
		HtmlURL          string      `json:"html_url"`
		ID               float64     `json:"id"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		IssuesURL        string      `json:"issues_url"`
		KeysURL          string      `json:"keys_url"`
		LabelsURL        string      `json:"labels_url"`
		Language         string      `json:"language"`
		LanguagesURL     string      `json:"languages_url"`
		MasterBranch     string      `json:"master_branch"`
		MergesURL        string      `json:"merges_url"`
		MilestonesURL    string      `json:"milestones_url"`
		MirrorURL        interface{} `json:"mirror_url"`
		Name             string      `json:"name"`
		NotificationsURL string      `json:"notifications_url"`
		OpenIssues       float64     `json:"open_issues"`
		OpenIssuesCount  float64     `json:"open_issues_count"`
		Owner            struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"owner"`
		Private         bool    `json:"private"`
		PullsURL        string  `json:"pulls_url"`
		PushedAt        float64 `json:"pushed_at"`
		ReleasesURL     string  `json:"releases_url"`
		Size            float64 `json:"size"`
		SshURL          string  `json:"ssh_url"`
		Stargazers      float64 `json:"stargazers"`
		StargazersCount float64 `json:"stargazers_count"`
		StargazersURL   string  `json:"stargazers_url"`
		StatusesURL     string  `json:"statuses_url"`
		SubscribersURL  string  `json:"subscribers_url"`
		SubscriptionURL string  `json:"subscription_url"`
		SvnURL          string  `json:"svn_url"`
		TagsURL         string  `json:"tags_url"`
		TeamsURL        string  `json:"teams_url"`
		TreesURL        string  `json:"trees_url"`
		UpdatedAt       string  `json:"updated_at"`
		URL             string  `json:"url"`
		Watchers        float64 `json:"watchers"`
		WatchersCount   float64 `json:"watchers_count"`
	} `json:"repository"`
	Sender struct {
		AvatarURL         string  `json:"avatar_url"`
		EventsURL         string  `json:"events_url"`
		FollowersURL      string  `json:"followers_url"`
		FollowingURL      string  `json:"following_url"`
		GistsURL          string  `json:"gists_url"`
		GravatarID        string  `json:"gravatar_id"`
		HtmlURL           string  `json:"html_url"`
		ID                float64 `json:"id"`
		Login             string  `json:"login"`
		OrganizationsURL  string  `json:"organizations_url"`
		ReceivedEventsURL string  `json:"received_events_url"`
		ReposURL          string  `json:"repos_url"`
		SiteAdmin         bool    `json:"site_admin"`
		StarredURL        string  `json:"starred_url"`
		SubscriptionsURL  string  `json:"subscriptions_url"`
		Type              string  `json:"type"`
		URL               string  `json:"url"`
	} `json:"sender"`
}

func Buffer2payload(b *bytes.Buffer) Payload {

	decoder := json.NewDecoder(strings.NewReader(b.String()))

	var p Payload
	err := decoder.Decode(&p)
	if err != nil {
		logger.Log("Failed to decode payload")
	}
	return p
}
