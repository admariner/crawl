package data

type Link struct {
	Address  *Address
	Anchor   string
	Href     string
	Nofollow bool
}

func MakeLink(base *Address, href string, anchor string, nofollow bool) *Link {
	link := &Link{
		Href:     href,
		Anchor:   anchor,
		Nofollow: nofollow,
		Address:  MakeAddressResolved(base, href),
	}
	return link
}
