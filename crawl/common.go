package crawl

import (
	"bytes"
	"github.com/bitly/go-simplejson"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"zhihu-spider/storage"
)

type Fetcher interface {
	Fetch()
}

func streamToJson(stream io.ReadCloser) (*simplejson.Json, error) {
	defer stream.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	js, err := simplejson.NewJson(buf.Bytes())
	if err != nil {
		log.Printf("parse json err:%v\n", err)
		return nil, err
	}
	if js.Get("error").Get("need_login").MustBool() == true {
		//log.Println("")
		//return nil, errors.New("fetch refuse")
		panic("fetch refuse by zhihu...")
	}
	return js, nil
}

//IsAnonymousUser检查指定用户ID是否是匿名
func IsAnonymousUser(id uuid.UUID) bool {
	if id == uuid.Nil {
		//log.Printf("id <%v> is anonymous", id)
		return true
	}
	return false
}

//AttentionQuestionOrSkip 关注问题
func AttentionQuestionOrSkip(userID uuid.UUID, questionID int) {
	if IsAnonymousUser(userID) {
		return
	}
	db := storage.GetDB()
	//检查用户是否已关注该问题
	//select count(1) from attentions where user_id = userID and question_refer = questionID
	//var count int64
	//db.Model(&storage.Attention{}).Where("user_id = ? AND question_refer = ?", userID, questionID).Count(&count)
	//if count == 0 {
	//	db.Create(&storage.Attention{
	//		UserID:        userID,
	//		QuestionRefer: questionID,
	//	})
	//}
	//return
	newAttentionItem := storage.Attention{
		UserID:        userID,
		QuestionRefer: questionID,
	}
	db.Where(storage.Attention{
		UserID:        userID,
		QuestionRefer: questionID,
	}).FirstOrCreate(&newAttentionItem)
}

//FollowUser 关注用户
func FollowUser(follower, followee uuid.UUID) {
	//处理匿名账号
	if IsAnonymousUser(follower) || IsAnonymousUser(followee) {
		return
	}
	if uuid.Equal(follower, followee) {
		return
	}
	//先检查follower是否关注了followee
	//select count(1) from follows where follower_id = follower AND followee = followee_id
	//var count int64
	//storage.GetDB().Transaction(func(tx *gorm.DB) error {
	//	tx.Model(&storage.Follow{}).Where("follower_id = ? AND followee_id = ?", follower, followee).Count(&count)
	//	if count == 0 {
	//		if err := tx.Create(&storage.Follow{
	//			FollowerID:  	follower,
	//			FolloweeID: 	followee,
	//		}).Error; err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
	FollowItem := storage.Follow{
		FollowerID: follower,
		FolloweeID: followee,
	}
	storage.GetDB().Where(storage.Follow{
		FollowerID: follower,
		FolloweeID: followee,
	}).FirstOrCreate(&FollowItem)
	//db.Model(&storage.Follow{}).Where("follower_id = ? AND followee_id = ?", follower, followee).Count(&count)
	//if count == 0 {
	//	db.Create(&storage.Follow{
	//		FollowerID:  	follower,
	//		FolloweeID: 	followee,
	//	})
	//}
	//return
}

//CreateUserOrSkip 创建用户
func CreateUserOrSkip(user *storage.User) {
	//事务
	//var count int64
	//storage.GetDB().Transaction(func(tx *gorm.DB) error {
	//	tx.Model(&storage.User{}).Where("id = ?", user.ID).Count(&count)
	//	if count == 0 {
	//		if err := tx.Create(user).Error; err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
	if IsAnonymousUser(user.ID) {
		//匿名用户，无序创建新账号
		//log.Printf("don't create userID<%v-%s-%s>\n", user.ID, user.Name, user.HeadLine)
		return
	}
	storage.GetDB().Where(storage.User{
		ID: user.ID,
	}).FirstOrCreate(user)
	return
}

//创建评论
func CreateCommentOrSkip(comment *storage.Comment) {
	//var count int64
	//storage.GetDB().Transaction(func(tx *gorm.DB) error {
	//	tx.Model(&storage.Comment{}).Where("comment_id = ?", comment.CommentID).Count(&count)
	//	if count == 0 {
	//		if err := tx.Create(comment).Error; err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
	//if count > 0 {
	//	return false
	//}
	//return true
	storage.GetDB().Where(storage.Comment{
		OldCommentID: comment.OldCommentID,
	}).FirstOrCreate(comment)
}

//UpdateAnswerCountOfUser 更新用户评论数
func UpdateAnswerCountOfUser(userID uuid.UUID) {
	if IsAnonymousUser(userID) {
		return
	}
	//原子更新
	db := storage.GetDB()
	db.Exec("UPDATE users SET answer_count = answer_count + 1 where id = ?", userID)
}

//随机选择一个用户ID作为问题发布人
func selectRandomUserID() uuid.UUID {
	db := storage.GetDB()
	var userItem storage.User
	result := db.Raw("select * from users order by random() limit 1").Scan(&userItem)
	if result.RowsAffected == 0 {
		userItem = storage.User{
			ID:       uuid.NewV4(),
			HeadLine: "测试",
			Name:     "JIONGUA",
		}
		db.Create(&userItem)
	}
	return userItem.ID
}
