package shop

import (
	"user-command-module/db/repository"
	domain_shop "user-command-module/internal/domain/shop"

	"github.com/iamKienb/go-core/postgres/conv"
)

func toInfraShop(shop *domain_shop.Shop) repository.SaveShopParams {
	return repository.SaveShopParams{
		ID:        conv.UUID(shop.ID),
		OwnerID:   conv.UUID(shop.OwnerID),
		Name:      shop.Name,
		Slug:      shop.Slug,
		Status:    string(shop.Status),
		CreatedBy: conv.UUID(shop.CreatedBy),
		UpdatedBy: conv.UUID(*shop.UpdatedBy),
		CreatedAt: conv.TimeStampZ(&shop.CreatedAt),
		UpdatedAt: conv.TimeStampZ(shop.UpdatedAt),
	}
}

func toInfraProfile(profile *domain_shop.Profile) repository.SaveShopProfileParams {
	return repository.SaveShopProfileParams{
		ShopID:      conv.UUID(profile.ShopID),
		Description: conv.Text(profile.Description),
		LogoUrl:     conv.Text(profile.LogoUrl),
		BannerUrl:   conv.Text(profile.BannerUrl),
		CreatedBy:   conv.UUID(profile.CreatedBy),
		UpdatedBy:   conv.UUID(*profile.UpdatedBy),
		CreatedAt:   conv.TimeStampZ(&profile.CreatedAt),
		UpdatedAt:   conv.TimeStampZ(profile.UpdatedAt),
	}
}

func toInfraShopAddress(address *domain_shop.ShopAddress) repository.SaveShopAddressParams {
	return repository.SaveShopAddressParams{
		ID:          conv.UUID(address.ID),
		ShopID:      conv.UUID(address.ShopID),
		CountryID:   int32(address.CountryID),
		CityID:      int32(address.CityID),
		DistrictID:  int32(address.DistrictID),
		WardID:      int32(address.WardID),
		AddressLine: address.AddressLine,
		ContactName: address.ContactName,
		PhoneNumber: address.PhoneNumber,
		Type:        string(address.Type),
		CreatedAt:   conv.TimeStampZ(&address.CreatedAt),
		UpdatedAt:   conv.TimeStampZ(address.UpdatedAt),
		CreatedBy:   conv.UUID(address.CreatedBy),
		UpdatedBy:   conv.UUID(*address.UpdatedBy),
	}
}
