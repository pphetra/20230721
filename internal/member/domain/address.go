package member_domain

type Address struct {
	Line1      string `json:"line_1"`
	Line2      string `json:"line_2"`
	PostalCode string `json:"postal_code"`
}

func NewAddress(
	line1 string,
	line2 string,
	postalCode string,
) (Address, error) {
	return Address{
		Line1:      line1,
		Line2:      line2,
		PostalCode: postalCode,
	}, nil
}
