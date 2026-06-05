package port

import "encoding/json"

type Page struct {
	Size  int
	Token string
}

type Shop struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	Status    string         `json:"status"`
	Profile   *ShopProfile   `json:"profile"`
	Addresses []ShopAddress  `json:"address"`
	Members   []ShopMember   `json:"members"`
	Extra     map[string]any `json:"-"`
}

func (s *Shop) UnmarshalJSON(data []byte) error {
	type shopAlias Shop
	var raw struct {
		shopAlias
		Address json.RawMessage `json:"address"`
		Members json.RawMessage `json:"members"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*s = Shop(raw.shopAlias)
	s.Addresses = decodeOneOrMany[ShopAddress](raw.Address)
	s.Members = decodeOneOrMany[ShopMember](raw.Members)
	return nil
}

type ShopProfile struct {
	Description string         `json:"description"`
	LogoURL     string         `json:"logo_url"`
	BannerURL   string         `json:"banner_url"`
	Extra       map[string]any `json:"-"`
}

type ShopAddress struct {
	ID          string         `json:"id"`
	ShopID      string         `json:"shop_id"`
	FullAddress string         `json:"full_address"`
	AddressLine string         `json:"address_line"`
	ContactName string         `json:"contact_name"`
	PhoneNumber string         `json:"phone_number"`
	Type        string         `json:"type"`
	Extra       map[string]any `json:"-"`
}

type ShopMember struct {
	ID      string         `json:"member_id"`
	Name    string         `json:"member_name"`
	RoleIDs []int32        `json:"role_ids"`
	Extra   map[string]any `json:"-"`
}

type ShopPage struct {
	Items         []Shop
	Total         int64
	NextPageToken string
}

func decodeOneOrMany[T any](raw json.RawMessage) []T {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var many []T
	if err := json.Unmarshal(raw, &many); err == nil {
		return many
	}
	var one T
	if err := json.Unmarshal(raw, &one); err != nil {
		return nil
	}
	return []T{one}
}
