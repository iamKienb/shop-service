package shop

import (
	"shop-query-module/internal/application/service/models"

	api "github.com/iamKienb/api-contract/gen/shop"
)

func ToShopView(shop *models.Shop) *api.ShopDetail {
	if shop == nil {
		return nil
	}
	return &api.ShopDetail{
		ShopId:    shop.ID,
		Name:      shop.Name,
		Slug:      shop.Slug,
		Status:    shop.Status,
		Profile:   ToShopProfileView(shop.Profile),
		Addresses: ToShopAddressViews(shop.Addresses),
		Members:   ToShopMemberViews(shop.Members),
	}
}

func ToShopViews(shops []models.Shop) []*api.ShopDetail {
	if len(shops) == 0 {
		return nil
	}

	views := make([]*api.ShopDetail, 0, len(shops))
	for i := range shops {
		views = append(views, ToShopView(&shops[i]))
	}
	return views
}

func ToShopProfileView(profile *models.ShopProfile) *api.ShopProfileDetail {
	if profile == nil {
		return nil
	}

	return &api.ShopProfileDetail{
		Description: profile.Description,
		LogoUrl:     profile.LogoURL,
		BannerUrl:   profile.BannerURL,
	}
}

func ToShopAddressViews(addresses []models.ShopAddress) []*api.ShopAddressDetail {
	if len(addresses) == 0 {
		return nil
	}

	views := make([]*api.ShopAddressDetail, 0, len(addresses))
	for _, address := range addresses {
		views = append(views, &api.ShopAddressDetail{
			Id:     address.ID,
			ShopId: address.ShopID,

			ProvinceId:   address.Province.ID,
			ProvinceName: address.Province.Name,
			WardId:       address.Ward.ID,
			WardName:     address.Ward.Name,
			AddressLine:  address.AddressLine,
			FullAddress:  address.FullAddress,
			ContactName:  address.ContactName,
			PhoneNumber:  address.PhoneNumber,
			Type:         address.Type,
		})
	}
	return views
}

func ToShopMemberViews(members []models.ShopMember) []*api.ShopMemberDetail {
	if len(members) == 0 {
		return nil
	}
	views := make([]*api.ShopMemberDetail, 0, len(members))

	for _, member := range members {
		memberRoles := make([]*api.Role, 0, len(member.Roles))
		for _, mbRole := range member.Roles {
			memberRoles = append(memberRoles, &api.Role{
				Id:   mbRole.ID,
				Code: mbRole.Code,
				Name: mbRole.Name,
			})
		}

		views = append(views, &api.ShopMemberDetail{
			Id:      member.ID,
			Name:    member.Name,
			RoleIds: member.RoleIDs,
			Roles:   memberRoles,
		})
	}
	return views
}
