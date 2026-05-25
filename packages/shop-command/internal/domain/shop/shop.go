package shop

import (
	"errors"
	"time"
	"user-command-module/internal/domain/shared"
)

type ShopStatus string

const (
	StatusPending  ShopStatus = "PENDING"
	StatusActive   ShopStatus = "ACTIVE"
	StatusInactive ShopStatus = "INACTIVE"
	StatusBanned   ShopStatus = "BANNED"
	StatusDeleted  ShopStatus = "DELETED"
)

type Shop struct {
	ID      shared.ShopID
	OwnerID shared.UserID
	Name    string
	Slug    string
	Status  ShopStatus

	CreatedBy shared.UserID
	UpdatedBy *shared.UserID

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

	Profile   Profile
	Addresses []ShopAddress
	shared.EventEntity
}

func NewShop(params NewShopParams) *Shop {
	now := time.Now().UTC()
	shopID := shared.NewID[shared.ShopID]()

	profile := Profile{
		ShopID:      shopID,
		Description: params.Description,
		LogoUrl:     params.LogoUrl,
		BannerUrl:   params.BannerUrl,

		CreatedBy: params.UserID,
		UpdatedBy: nil,

		CreatedAt: now,
		UpdatedAt: nil,
	}

	shop := &Shop{
		ID:      shopID,
		OwnerID: params.UserID,
		Name:    params.Name,
		Slug:    params.Slug,
		Status:  StatusPending,

		Profile:   profile,
		Addresses: []ShopAddress{},

		CreatedBy: params.UserID,
		UpdatedBy: nil,

		CreatedAt: now,
		UpdatedAt: nil,
		DeletedAt: nil,
	}

	shop.AddEvent(ShopCreatedEvent{
		ShopID:    shopID,
		OwnerID:   shop.OwnerID,
		Name:      shop.Name,
		Slug:      shop.Slug,
		Status:    shop.Status,
		CreatedBy: shop.CreatedBy,
		CreatedAt: now,
	})

	shop.AddEvent(ShopProfileCreatedEvent{
		ShopID:      shopID,
		Description: profile.Description,
		LogoUrl:     profile.LogoUrl,
		BannerUrl:   profile.BannerUrl,
		CreatedBy:   profile.CreatedBy,
		CreatedAt:   now,
	})

	params.Address.ShopID = shopID
	shop.AddAddress(params.Address)

	return shop
}

func (s *Shop) FlushEvents() []shared.DomainEvent {
	var events []shared.DomainEvent
	events = append(events, s.Flush()...)
	s.ClearEvent()

	return events
}

func (s *Shop) AddAddress(params NewShopAddressParams) (*ShopAddress, error) {
	if s.Status == StatusBanned || s.Status == StatusInactive {
		return nil, ErrShopNotAllowed
	}

	addressID := shared.NewID[shared.ShopAddressID]()
	now := time.Now().UTC()

	address := ShopAddress{
		ID:     addressID,
		ShopID: params.ShopID,

		CountryID:  params.CountryID,
		CityID:     params.CityID,
		DistrictID: params.DistrictID,
		WardID:     params.WardID,

		Type:        params.Type,
		ContactName: params.ContactName,
		PhoneNumber: params.PhoneNumber,
		AddressLine: params.AddressLine,

		CreatedBy: params.UserID,
		UpdatedBy: nil,

		CreatedAt: now,
		UpdatedAt: nil,
	}

	if s.Addresses != nil {
		s.Addresses = append(s.Addresses, address)
	}

	s.AddEvent(ShopAddressAddedEvent{
		ShopID:        address.ShopID,
		ShopAddressID: addressID,

		CountryID:   params.CountryID,
		CountryName: params.CountryName,

		CityID:   params.CityID,
		CityName: params.CityName,

		DistrictID:   params.DistrictID,
		DistrictName: params.DistrictName,

		WardID:   params.WardID,
		WardName: params.WardName,

		AddressLine: address.AddressLine,
		ContactName: address.ContactName,
		PhoneNumber: address.PhoneNumber,
		Type:        address.Type,

		CreatedBy: address.CreatedBy,
		CreatedAt: now,
	})

	return &address, nil
}

func (s *Shop) ActivateIfNeeded(userID shared.UserID) error {
	if s.Status == StatusActive {
		return nil
	}

	if s.Status == StatusBanned {
		return errors.New("cannot activate a banned shop")
	}

	if s.Name == "" || s.Profile.Description == nil {
		return errors.New("shop name and description are required for activation")
	}

	if s.Profile.LogoUrl == nil {
		return errors.New("shop logo is required for activation")
	}

	hasPickupAddress := false
	for _, addr := range s.Addresses {
		if addr.Type == TypePickup || addr.Type == TypeReturn {
			hasPickupAddress = true
			break
		}
	}
	if !hasPickupAddress {
		return errors.New("shop must have at least one pickup address before activation")
	}

	s.Status = StatusActive
	now := time.Now().UTC()
	s.UpdatedAt = &now
	s.UpdatedBy = &userID

	s.AddEvent(ShopActivatedEvent{
		ShopID:    s.ID,
		Status:    s.Status,
		UpdatedBy: *s.UpdatedBy,
		UpdatedAt: *s.UpdatedAt,
	})

	return nil
}

func (s *Shop) Type() string {
	return "SHOP"
}

func (s *Shop) MatchesSlug(slug string) bool {
	return s.Slug == slug
}

func (s *Shop) IsActive() bool {
	return s.Status == StatusActive
}
