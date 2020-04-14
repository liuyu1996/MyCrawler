package model

type Profile struct {
	Name string
	Gender string
	Age int
	Height int
	Weight int
	Income string
	Marriage string
	Education string
	Occupation string
	Hokou string
	Xinzuo string
	House string
	Car string
}

type BilibiliProfile struct {
	Title string
	DatePublished string
	CommentCount  int
	Author        string
	PlayCount     string
	DanMuCount    string
	Scores        int
}

type VideoInfo struct {
	Title string
	PlayCount  string
	DanMuCount string
	Scores     int
}