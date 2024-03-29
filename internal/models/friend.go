package models

//type Friend struct {
//	ID     uint `gorm:"primaryKey"`
//	UserID uint
//	User   User   `gorm:"foreignKey:UserID;references:ID"`
//	Friend []User `gorm:"many2many:friendship"`
//	DefaultModel
//}
//
//func (fd *Friend) TableName() string {
//	return "friend"
//}
//
//func (fd *Friend) FindOneFriend(db *gorm.DB, ctx context.Context, friendID uint) (*User, error) {
//	var friend User
//	if err := db.WithContext(ctx).Debug().Model(&fd).Where(User{
//		ID: friendID,
//	}).Association("Friend").Find(&friend); err != nil {
//		return nil, err
//	}
//
//	return &friend, nil
//}
////
////func (fd *Friend) GetUserFriend(db *gorm.DB, ctx context.Context) error {
////	return db.WithContext(ctx).Debug().Where("user_id = ?", fd.UserID).Find(&fd).Error
////}
//
//func (fd *Friend) CountFriend(db *gorm.DB, ctx context.Context) int64 {
//	return db.WithContext(ctx).Debug().Model(&fd).Association("Friend").Count()
//}
//
//func (fd *Friend) GetFriendsList(db *gorm.DB, ctx context.Context) ([]*User, error) {
//	var friends []*User
//	if err := db.WithContext(ctx).Debug().Model(&fd).Association("Friend").Find(&friends); err != nil {
//		return nil, err
//	}
//	return friends, nil
//}
//
////InsertOne - Call it after account has been created.
//func (fd *Friend) InsertOne(db *gorm.DB, ctx context.Context) error {
//	//Add a new Friend
//	return db.WithContext(ctx).Debug().Create(&fd).Error
//}
//
//func (fd *Friend) RemoveOne(db *gorm.DB, ctx context.Context, userID, friendID uint) error {
//	//Remove an existing Friend
//	return db.WithContext(ctx).Debug().Transaction(func(tx *gorm.DB) error {
//		//Friendship : A -> B
//		if err := tx.WithContext(ctx).Debug().Model(&Friend{UserID: userID}).Association("Friend").Delete(&User{ID: friendID}); err != nil {
//			return err
//		}
//
//		//Friendship : B -> A
//		if err := tx.WithContext(ctx).Debug().Model(&Friend{UserID: friendID}).Association("Friend").Delete(&User{ID: userID}); err != nil {
//			return err
//		}
//		return nil
//	})
//}
//
//func (fd *Friend) IsFriend(db *gorm.DB, ctx context.Context, friendID uint) (bool, error) {
//	var friend User
//
//	if err := db.WithContext(ctx).Debug().Where("user_id = ?", fd.UserID).First(&fd).Error; err != nil {
//		return false, err
//	}
//
//	err := db.WithContext(ctx).Debug().Model(&fd).Where("user_id = ?", friendID).Association("Friend").Find(&friend)
//	if err != nil {
//		return false, err
//	}
//
//	if friend.ID == 0 {
//		return false, nil
//	}
//
//	return true, nil
//}
