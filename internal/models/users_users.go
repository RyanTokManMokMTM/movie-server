package models

//type FriendTemp struct {
//	UserId   uint `gorm:"primaryKey"` // User is following FriendTemp
//	FriendID uint `gorm:"primaryKey"` // FriendTemp is followed by User
//	State    uint
//	DefaultModel
//}
//
//func (m *FriendTemp) TableName() string {
//	return "friends"
//}
//
//func (m *FriendTemp) BeforeCreate(db *gorm.DB) error {
//	m.State = 1 //when create set to follow
//	return nil
//}
//
////Following user
//func (m *FriendTemp) AddNewsFriend(ctx context.Context, db *gorm.DB) error {
//	//User ID Follow FriendID
//	return db.Debug().WithContext(ctx).Model(&m).Create(&m).Error
//}
//
////UnFollowing user
//func (m *FriendTemp) UpdateFriendState(ctx context.Context, db *gorm.DB) error {
//	return db.Debug().WithContext(ctx).Model(&m).Update("state", m.State).Error
//}
//
//// Is User Followed by owner
//func (m *FriendTemp) FindOneUserFromFriendList(ctx context.Context, db *gorm.DB) error {
//	// what should it return ???
//	return db.Debug().WithContext(ctx).Model(&m).First(&m).Error
//}
//
//func (m *FriendTemp) CountFollowingUser(ctx context.Context, db *gorm.DB) (int64, error) {
//	var count int64
//	//user -> friend
//	//example : userA -> B ,userA -> C,userA -> D
//	if err := db.Debug().WithContext(ctx).Model(&m).Where("user_id = ? AND state = 1", m.UserId).Count(&count).Error; err != nil {
//		return 0, err
//	}
//	return count, nil
//}
//
//func (m *FriendTemp) CountFollowedUser(ctx context.Context, db *gorm.DB) (int64, error) {
//	var count int64
//	//FriendA -> UserA
//	//FriendB -> UserA
//	//FriendC -> UserA
//	if err := db.Debug().WithContext(ctx).Model(&m).Where("friend_id = ? AND state = 1", m.UserId).Count(&count).Error; err != nil {
//		return 0, err
//	}
//	return count, nil
//}
