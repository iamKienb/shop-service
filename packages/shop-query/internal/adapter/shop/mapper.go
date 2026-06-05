package shop

import (
	"shop-query-module/internal/application/port"

	api "github.com/iamKienb/api-contract/gen/shop"
)

func ToShopView(shop *port.Shop) *api.ShopView {
	if shop == nil {
		return nil
	}
	return &api.ShopView{
		ShopId:    shop.ID,
		Name:      shop.Name,
		Slug:      shop.Slug,
		Status:    shop.Status,
		Profile:   ToShopProfileView(shop.Profile),
		Addresses: ToShopAddressViews(shop.Addresses),
		Members:   ToShopMemberViews(shop.Members),
	}
}

func ToShopViews(shops []port.Shop) []*api.ShopView {
	views := make([]*api.ShopView, 0, len(shops))
	for i := range shops {
		views = append(views, ToShopView(&shops[i]))
	}
	return views
}

func ToShopProfileView(profile *port.ShopProfile) *api.ShopProfileView {
	if profile == nil {
		return nil
	}
	return &api.ShopProfileView{
		Description: profile.Description,
		LogoUrl:     profile.LogoURL,
		BannerUrl:   profile.BannerURL,
	}
}

func ToShopAddressViews(addresses []port.ShopAddress) []*api.ShopAddressView {
	views := make([]*api.ShopAddressView, 0, len(addresses))
	for _, address := range addresses {
		views = append(views, &api.ShopAddressView{
			Id:          address.ID,
			ShopId:      address.ShopID,
			FullAddress: address.FullAddress,
			AddressLine: address.AddressLine,
			ContactName: address.ContactName,
			PhoneNumber: address.PhoneNumber,
			Type:        address.Type,
		})
	}
	return views
}

func ToShopMemberViews(members []port.ShopMember) []*api.ShopMemberView {
	views := make([]*api.ShopMemberView, 0, len(members))
	for _, member := range members {
		views = append(views, &api.ShopMemberView{
			Id:      member.ID,
			Name:    member.Name,
			RoleIds: member.RoleIDs,
		})
	}
	return views
}
