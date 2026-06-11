package models

import "time"

type Page struct {
	Size  int
	Token string
}

type Shop struct {
	ID        string        `json:"id"`
	OwnerID   string        `json:"owner_id"`
	Name      string        `json:"name"`
	Slug      string        `json:"slug"`
	Status    string        `json:"status"`
	Profile   *ShopProfile  `json:"profile"`
	Addresses []ShopAddress `json:"addresses"`
	Members   []ShopMember  `json:"members"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`

	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ShopProfile struct {
	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`
	BannerURL   string `json:"banner_url"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`

	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LocalRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ShopAddress struct {
	ID     string `json:"id"`
	ShopID string `json:"shop_id"`

	Country  LocalRef `json:"country"`
	Province LocalRef `json:"province"`
	Ward     LocalRef `json:"ward"`

	AddressLine string `json:"address_line"`
	FullAddress string `json:"full_address"`
	ContactName string `json:"contact_name"`
	PhoneNumber string `json:"phone_number"`
	Type        string `json:"type"`
}

type Role struct {
	ID   int32  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ShopMember struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	RoleIDs     []int32   `json:"role_ids"`
	AddedBy     string    `json:"added_by"`
	NameAddedBy string    `json:"name_added_by"`
	JoinedAt    time.Time `json:"joined_at"`
	Roles       []Role    `json:"roles"`
}

type ShopPage struct {
	Items         []Shop
	Total         int64
	NextPageToken string
}
