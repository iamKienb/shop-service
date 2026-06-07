package shop

import (
	"shop-command-module/internal/domain/shared"
	"time"
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

		CountryID:   params.CountryID,
		CountryName: params.CountryName,

		ProvinceID:   params.ProvinceID,
		ProvinceName: params.ProvinceName,

		WardID:   params.WardID,
		WardName: params.WardName,

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

		ProvinceID:   params.ProvinceID,
		ProvinceName: params.ProvinceName,

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

func (s *Shop) TryActivate(userID shared.UserID) (bool, error) {
	if s.Status == StatusActive {
		return false, nil
	}

	if s.Status == StatusBanned {
		return false, ErrShopNotAllowed
	}

	if !s.isActivationReady() {
		return false, nil
	}

	now := time.Now().UTC()
	s.Status = StatusActive
	s.UpdatedAt = &now
	s.UpdatedBy = &userID

	s.AddEvent(ShopActivatedEvent{
		ShopID:    s.ID,
		Status:    s.Status,
		UpdatedBy: *s.UpdatedBy,
		UpdatedAt: *s.UpdatedAt,
	})

	return true, nil
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

func (s *Shop) isActivationReady() bool {
	if s.Name == "" || isBlankPointer(s.Profile.Description) || isBlankPointer(s.Profile.LogoUrl) {
		return false
	}

	return s.hasAddressType(TypePickup) && s.hasAddressType(TypeReturn)
}

func (s *Shop) hasAddressType(addressType AddressTypeEnum) bool {
	for _, addr := range s.Addresses {
		if addr.Type == addressType {
			return true
		}
	}

	return false
}

func isBlankPointer(value *string) bool {
	return value == nil || *value == ""
}
