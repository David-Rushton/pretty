package pretty

type Link string

func WithLink(url string) Link {
	return Link(url)
}
