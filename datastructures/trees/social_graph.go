package trees

type SocialGraph struct {
	Size  int
	Links [][]Link
}

type Link struct {
	Vertex1 Record
	Vertex2 Record
	Weight  int
}

type Record struct {
	ID   int
	data string
}

func NewSocialGraph(size int) *SocialGraph {
	return &SocialGraph{
		Size:  size,
		Links: make([][]Link, size),
	}
}
