package shop

import (
	"user-command-module/internal/application/commands/add_shop_address"
	"user-command-module/internal/application/commands/assign_member"
	"user-command-module/internal/application/commands/create_shop"
	"user-command-module/internal/application/queries/check_permission"
	"user-command-module/internal/domain/shared"
	domain_shop "user-command-module/internal/domain/shop"

	"github.com/iamKienb/api-contract/gen/shop"
)

func ToCreateShopCommand(userID string, userName string, req *shop.CreateShopRequest) (create_shop.Command, error) {
	var profile *create_shop.Profile
	if req.Profile != nil {
		profile = &create_shop.Profile{
			Description: req.Profile.Description,
			LogoUrl:     req.Profile.LogoUrl,
			BannerUrl:   req.Profile.BannerUrl,
		}
	}
	userParse, err := shared.ParseToRawID[shared.UserID](userID)
	if err != nil {
		return create_shop.Command{}, shared.ErrUserInvalid
	}

	return create_shop.Command{
		User: create_shop.User{
			ID:   userParse,
			Name: userName,
		},
		Name:    req.GetName(),
		Slug:    req.GetSlug(),
		Profile: profile,
	}, nil
}

func ToCreateShopResponse(result *create_shop.Result) *shop.CreateShopResponse {
	return &shop.CreateShopResponse{
		ShopId: result.ShopID,
	}
}

func ToAssignMemberCommand(userID string, userName string, req *shop.AssignMemberRolesRequest) (assign_member.Command, error) {
	shopID, err := shared.ParseToRawID[shared.ShopID](req.GetShopId())
	if err != nil {
		return assign_member.Command{}, domain_shop.ErrShopInvalid
	}
	userParse, err := shared.ParseToRawID[shared.UserID](userID)
	if err != nil {
		return assign_member.Command{}, shared.ErrUserInvalid
	}

	reqMemberRoles := req.GetMemberRoles()
	memberRoles := make([]assign_member.MemberRole, 0, len(reqMemberRoles))

	for _, memberRole := range reqMemberRoles {
		memberID, err := shared.ParseToRawID[shared.UserID](memberRole.GetId())
		if err != nil {
			return assign_member.Command{}, shared.ErrUserInvalid
		}

		reqRoleIDs := memberRole.GetRoleIDs()
		roleIDs := make([]shared.RoleID, 0, len(reqRoleIDs))

		for _, roleID := range reqRoleIDs {
			roleIDs = append(roleIDs, shared.RoleID(roleID.GetId()))
		}

		memberRoles = append(memberRoles, assign_member.MemberRole{
			MemberID: memberID,
			Name:     memberRole.GetName(),
			RoleIDs:  roleIDs,
		})
	}

	return assign_member.Command{
		User: assign_member.User{
			ID:   userParse,
			Name: userName,
		},
		ShopID:      shopID,
		MemberRoles: memberRoles,
		Action:      req.GetAction(),
	}, nil
}

func ToAssignMemberResponse(result *assign_member.Result) *shop.AssignMemberRolesResponse {
	return &shop.AssignMemberRolesResponse{
		Success: result.Success,
	}
}

func ToAddAddressCommand(userID string, req *shop.AddShopAddressRequest) (add_shop_address.Command, error) {
	parsedUserID, err := shared.ParseToRawID[shared.UserID](userID)
	if err != nil {
		return add_shop_address.Command{}, shared.ErrUserInvalid
	}
	shopID, err := shared.ParseToRawID[shared.ShopID](req.GetShopId())
	if err != nil {
		return add_shop_address.Command{}, domain_shop.ErrShopInvalid
	}

	return add_shop_address.Command{
		UserID: parsedUserID,
		ShopID: shopID,

		Country:  toShopLocationInfo(req.GetCountry()),
		City:     toShopLocationInfo(req.GetCity()),
		District: toShopLocationInfo(req.GetDistrict()),
		Ward:     toShopLocationInfo(req.GetWard()),

		AddressLine: req.GetAddressLine(),
		ContactName: req.GetContactName(),
		PhoneNumber: req.GetPhoneNumber(),
		Type:        req.GetType(),
	}, nil
}

func ToAddAddressResponse(result *add_shop_address.Result) *shop.AddShopAddressResponse {
	return &shop.AddShopAddressResponse{
		AddressId: result.ShopAddressID,
	}
}

func ToCheckPermissionQuery(userID string, req *shop.CheckPermissionRequest) (check_permission.Query, error) {
	parsedUserID, err := shared.ParseToRawID[shared.UserID](userID)
	if err != nil {
		return check_permission.Query{}, shared.ErrUserInvalid
	}

	shopID, err := shared.ParseToRawID[shared.ShopID](req.GetShopId())
	if err != nil {
		return check_permission.Query{}, domain_shop.ErrShopInvalid
	}

	return check_permission.Query{
		ShopID: shopID,
		UserID: parsedUserID,
		Action: req.GetAction(),
	}, nil

}

func ToCheckPermissionResponse(result *check_permission.Result) *shop.CheckPermissionResponse {
	return &shop.CheckPermissionResponse{
		IsAllowed: result.IsAllowed,
		Message:   result.Message,
	}
}

type shopLocationSource interface {
	GetId() int64
	GetName() string
}

func toShopLocationInfo(src shopLocationSource) add_shop_address.LocationInfo {
	if src == nil {
		return add_shop_address.LocationInfo{}
	}

	return add_shop_address.LocationInfo{
		ID:   int(src.GetId()),
		Name: src.GetName(),
	}
}
