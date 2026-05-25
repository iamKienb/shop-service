package shop

import (
	"context"
	"user-command-module/internal/application/commands/add_shop_address"
)

func (s *shopService) AddAddress(ctx context.Context, cmd add_shop_address.Command) (*add_shop_address.Result, error) {
	// 1. Load Aggregate Root lên (Chỉ cần thông tin Shop gốc để check Status, chưa cần load mảng Addresses)
	// currentShop, err := s.shopRepo.FindByID(ctx, cmd.ShopID)
	// if err != nil {
	// 	return err
	// }
	// if currentShop == nil {
	// 	return shop.ErrShopNotFound
	// }

	// // 2. Thực thi nghiệp vụ bên trong Domain và nhận về Entity Address mới cùng error
	// // (Đoạn check Banned/Inactive tự chạy bên trong hàm này)
	// newAddress, err := currentShop.AddAddress(shop.NewShopAddressParams{
	// 	UserID:      cmd.UserID,
	// 	CountryID:   cmd.CountryID,
	// 	CityID:      cmd.CityID,
	// 	AddressLine: cmd.AddressLine,
	// })
	// if err != nil {
	// 	return err // Trả về lỗi nghiệp vụ (ví dụ: Shop bị Banned) cho phía Client biết
	// }

	// // 3. Thực hiện lưu xuống DB trong Transaction nếu cần (Unit of Work)
	// if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
	// 	// Lưu địa chỉ mới tạo
	// 	if err := s.addressRepo.CreateAddress(ctx, newAddress); err != nil {
	// 		return err
	// 	}

	// 	// Thu thập và publish Domain Event phát sinh (ShopAddressAddedEvent)
	// 	if events := currentShop.FlushEvents(); len(events) > 0 {
	// 		err := s.outboxService.Publish(ctx, events)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}

	// 	return nil
	// }); err != nil {
	// 	return err
	// }

	return nil, nil
}
