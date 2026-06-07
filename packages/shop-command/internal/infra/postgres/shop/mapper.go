package shop

import (
	"time"

	"shop-command-module/db/repository"
	"shop-command-module/internal/domain/shared"
	domain_shop "shop-command-module/internal/domain/shop"

	"github.com/iamKienb/go-core/postgres/conv"
	"github.com/jackc/pgx/v5/pgtype"
)

func toInfraShop(shop *domain_shop.Shop) repository.CreateShopParams {
	return repository.CreateShopParams{
		ID:        conv.UUID(shop.ID),
		OwnerID:   conv.UUID(shop.OwnerID),
		Name:      shop.Name,
		Slug:      shop.Slug,
		Status:    string(shop.Status),
		CreatedBy: conv.UUID(shop.CreatedBy),
		UpdatedBy: nullableUUID(shop.UpdatedBy),
		CreatedAt: conv.TimeStampZ(&shop.CreatedAt),
		UpdatedAt: conv.TimeStampZ(shop.UpdatedAt),
	}
}

func toInfraProfile(profile *domain_shop.Profile) repository.CreateShopProfileParams {
	return repository.CreateShopProfileParams{
		ShopID:      conv.UUID(profile.ShopID),
		Description: conv.Text(profile.Description),
		LogoUrl:     conv.Text(profile.LogoUrl),
		BannerUrl:   conv.Text(profile.BannerUrl),
		CreatedBy:   conv.UUID(profile.CreatedBy),
		UpdatedBy:   nullableUUID(profile.UpdatedBy),
		CreatedAt:   conv.TimeStampZ(&profile.CreatedAt),
		UpdatedAt:   conv.TimeStampZ(profile.UpdatedAt),
	}
}

func toInfraShopAddress(address *domain_shop.ShopAddress) repository.CreateShopAddressParams {
	return repository.CreateShopAddressParams{
		ID:           conv.UUID(address.ID),
		ShopID:       conv.UUID(address.ShopID),
		CountryID:    address.CountryID,
		CountryName:  address.CountryName,
		ProvinceID:   address.ProvinceID,
		ProvinceName: address.ProvinceName,
		WardID:       address.WardID,
		WardName:     address.WardName,
		AddressLine:  address.AddressLine,
		ContactName:  address.ContactName,
		PhoneNumber:  address.PhoneNumber,
		Type:         string(address.Type),
		CreatedAt:    conv.TimeStampZ(&address.CreatedAt),
		UpdatedAt:    conv.TimeStampZ(address.UpdatedAt),
		CreatedBy:    conv.UUID(address.CreatedBy),
		UpdatedBy:    nullableUUID(address.UpdatedBy),
	}
}

func toInfraShopStatus(shop *domain_shop.Shop) repository.UpdateShopStatusParams {
	return repository.UpdateShopStatusParams{
		ID:        conv.UUID(shop.ID),
		Status:    string(shop.Status),
		UpdatedBy: nullableUUID(shop.UpdatedBy),
		UpdatedAt: conv.TimeStampZ(shop.UpdatedAt),
	}
}

func toDomainShop(model repository.Shop, profile repository.ShopProfile, addresses []repository.ShopAddress) *domain_shop.Shop {
	shop := &domain_shop.Shop{
		ID:        shared.ShopID(model.ID.Bytes),
		OwnerID:   shared.UserID(model.OwnerID.Bytes),
		Name:      model.Name,
		Slug:      model.Slug,
		Status:    domain_shop.ShopStatus(model.Status),
		CreatedBy: shared.UserID(model.CreatedBy.Bytes),
		UpdatedBy: toDomainUserID(model.UpdatedBy),
		CreatedAt: model.CreatedAt.Time,
		UpdatedAt: toTimePointer(model.UpdatedAt),
		DeletedAt: toTimePointer(model.DeletedAt),
		Profile: domain_shop.Profile{
			ShopID:      shared.ShopID(profile.ShopID.Bytes),
			Description: toStringPointer(profile.Description),
			LogoUrl:     toStringPointer(profile.LogoUrl),
			BannerUrl:   toStringPointer(profile.BannerUrl),
			CreatedBy:   shared.UserID(profile.CreatedBy.Bytes),
			UpdatedBy:   toDomainUserID(profile.UpdatedBy),
			CreatedAt:   profile.CreatedAt.Time,
			UpdatedAt:   toTimePointer(profile.UpdatedAt),
		},
		Addresses: make([]domain_shop.ShopAddress, 0, len(addresses)),
	}

	for _, address := range addresses {
		shop.Addresses = append(shop.Addresses, domain_shop.ShopAddress{
			ID:           shared.ShopAddressID(address.ID.Bytes),
			ShopID:       shared.ShopID(address.ShopID.Bytes),
			CountryID:    address.CountryID,
			CountryName:  address.CountryName,
			ProvinceID:   address.ProvinceID,
			ProvinceName: address.ProvinceName,
			WardID:       address.WardID,
			WardName:     address.WardName,
			ContactName:  address.ContactName,
			PhoneNumber:  address.PhoneNumber,
			AddressLine:  address.AddressLine,
			Type:         domain_shop.AddressTypeEnum(address.Type),
			CreatedBy:    shared.UserID(address.CreatedBy.Bytes),
			UpdatedBy:    toDomainUserID(address.UpdatedBy),
			CreatedAt:    address.CreatedAt.Time,
			UpdatedAt:    toTimePointer(address.UpdatedAt),
		})
	}

	return shop
}

func nullableUUID(id *shared.UserID) pgtype.UUID {
	if id == nil {
		return pgtype.UUID{Valid: false}
	}

	return conv.UUID(*id)
}

func toStringPointer(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}

	result := value.String
	return &result
}

func toDomainUserID(value pgtype.UUID) *shared.UserID {
	if !value.Valid {
		return nil
	}

	id := shared.UserID(value.Bytes)
	return &id
}

func toTimePointer(value pgtype.Timestamptz) *time.Time {
	if !value.Valid {
		return nil
	}

	result := value.Time
	return &result
}
