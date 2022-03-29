package database

import (
	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
    "rest-api-test/entity"
)

// データインサート処理
func DBInsert(title string, content string) {
    db := GormConnect()

    defer db.Close()

    sample := entity.SampleEntity{Title: title, Content: content}
    // Insert処理
    db.Create(&sample)
}

//DB更新
// func dbUpdate(id int, tweetText string) {
//     db := gormConnect()
//     var tweet Tweet
//     db.First(&tweet, id)
//     tweet.Content = tweetText
//     db.Save(&tweet)
//     db.Close()
// }

// 全件取得
// func dbGetAll() []Tweet {
//     db := gormConnect()

//     defer db.Close()
//     var tweets []Tweet
//     // FindでDB名を指定して取得した後、orderで登録順に並び替え
//     db.Order("created_at desc").Find(&tweets)
//     return tweets
// }

//DB一つ取得
func dbGetOne(id int) entity.SampleEntity {
    db := GormConnect()
    var sample entity.SampleEntity
    db.First(&sample, id)
    db.Close()
    return sample
}

//DB削除
func dbDelete(id int) {
    db := GormConnect()
    var sample entity.SampleEntity
    db.First(&sample, id)
    db.Delete(&sample)
    db.Close()
}