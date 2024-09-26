package dao

import (
	"context"
	"price-subscriber/internal/database/gorms"
	datetable "price-subscriber/internal/datatable"
)

type ListPriceDao struct {
	Conn *gorms.GormConn
}

func NewListPriceDao() *ListPriceDao {
	return &ListPriceDao{
		Conn: gorms.New(),
	}
}

func (l ListPriceDao) GetListPrice(c context.Context, product string) (datetable.ListPriceItem, error) {
	listPrice := datetable.ListPriceItem{}
	err := l.Conn.DB.Where("product=?",product).First(&listPrice).Error
	return listPrice, err
}

func (l ListPriceDao) CreateListPrice(c context.Context, listPrice datetable.ListPriceItem) (bool, error) {
	result := l.Conn.DB.Create(&listPrice)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// func (m MemberDao) GetEmailFromMember(c context.Context, email string) (bool, error) {
// 	var count int64
// 	err := m.Conn.DB.Model(&datatable.Member{}).Where("email=?", email).Count(&count).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	if count > 0 {
// 		return true, nil
// 	}
// 	return false, nil
// }

// func (m MemberDao) GetPhoneFromMember(c context.Context, phone string) (bool, error) {
// 	var count int64
// 	//err := m.Conn.DB.Model(&datatable.Member{}).Where("email=?", phone).Count(&count).Error()
// 	err := m.Conn.DB.Raw("select count(*) as count from ms_member where mobile = ?", phone).Scan(&count).Error
// 	if err != nil {
// 		return false, err
// 	}

// 	if count > 0 {
// 		return true, nil
// 	}
// 	return false, nil

// }

// func (m MemberDao) GetAccountFromMember(c context.Context, account string) (bool, error) {
// 	var count int64
// 	err := m.Conn.DB.Model(&datatable.Member{}).Where("email=?", account).Count(&count).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	if count > 0 {
// 		return true, nil
// 	}
// 	return false, nil
// }

// func (m MemberDao) InsertUserTOMember(conn database.DbConn, c context.Context, member datatable.Member) (bool, error) {
// 	m.Conn = conn.(*gorms.GormConn)
// 	result := m.Conn.Tx(c).Create(&member)
// 	err := result.Error
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// func (m MemberDao) GetUserByAccount(ctx context.Context, account string) (datatable.Member, error) {
// 	var member datatable.Member
// 	err := m.Conn.DB.Where("account = ?", account).First(&member).Error
// 	return member, err
// }
