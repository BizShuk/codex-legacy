package i

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// [Question]: Covert int to base-7 string
func AmazingTalker_OA1(input int) string {
	var b7str string
	var minus string
	if input == 0 {
		return "0"
	}
	if input < 0 {
		minus = "-"
		input = input * -1
	}

	for input > 0 {
		q := input / 7
		r := input % 7
		input = q

		c := strconv.Itoa(r)
		b7str = c + b7str
	}

	return minus + b7str
}

// [Question]: Show list which two elements sum is maximum and than K
// [Assumption]: maxSum >= 0
// sorted array? If sorted, O(nlogn)
// negative number? assume only positive here
func AmazingTalker_OA2(list []int, k int) [][]int {
	maxSum, resultSet := -1, [][]int{}
	if len(list) <= 1 {
		return resultSet
	}

	sort.Ints(list)
	lp, rp := 0, len(list)-1

	for lp < rp {
		sum := list[lp] + list[rp]
		if sum >= k {
			rp--
			continue
		}

		if maxSum == sum {
			resultSet = append(resultSet, []int{list[lp], list[rp]})
		}

		if maxSum < sum {
			maxSum = sum
			resultSet = [][]int{{list[lp], list[rp]}}
		}

		lp++
	}

	return resultSet
}

// [Question]: Give test cases of one function. It can write without real code.
func AmazingTalker_OA3() {
	// 1. empty name
	// 2. name with symbol start
	// 3. name with digit start
	// 4. name with maximun length
	// 5. record exist in db
	// 6. mq is down by fake mq connection
	// 7. db failed connection.
}

// [Question]: Fix the cache set/get with context on API
func AmazingTalker_OA4() {
	r := gin.Default()
	r.GET("/admin/teachers", GetTeachers)
	_ = r.Run()
}

// ============= DO NOT EDIT =============

var db gorm.DB         // connected db client
var cache redis.Client // connected redis client

type Teacher struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func cacheMarshaller(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func cacheUnMarshaller(val string, container interface{}) error {
	return json.Unmarshal([]byte(val), container)
}

// =========== DO NOT EDIT END ===========
// =======================================

func GetTeachers(ctx *gin.Context) {

	cacheKey := GetTeachersCacheKey(ctx)
	cacheResult, err := cache.Get(cacheKey).Result()

	if err != nil && err != redis.Nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	var teachers []Teacher
	if err != redis.Nil {
		err := cacheUnMarshaller(cacheResult, &teachers)
		if err == nil {
			ctx.JSON(200, teachers)
			return // get teachers from cache

		}
		// Cache fetch failed, need to query from db. Alert required
		// ctx.JSON(500, gin.H{"message": err.Error()})
	}

	query := GetTeachersQuery(ctx)
	err = query.Find(&teachers).Error
	if err != nil { // fail to query db. system error
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	bs, err := cacheMarshaller(teachers)
	if err != nil { // Cache process failed, but get teachers already. Alert required
		// ctx.JSON(500, gin.H{"message": err.Error()})
	} else {
		err = cache.Set(cacheKey, bs, time.Minute*3).Err()
		if err != nil { // Insert key failed, but teachers are ready to response. alert required
			// ctx.JSON(500, gin.H{"message": err.Error()})
			logrus.Errorln(err)
		}
	}

	ctx.JSON(200, teachers)
}

func GetTeachersCacheKey(ctx *gin.Context) string {
	withDeleted := false
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	size, _ := strconv.ParseInt(ctx.Query("size"), 10, 64)

	if ctx.Query("deleted") == "true" {
		withDeleted = true
	}

	return fmt.Sprintf("%d-%d-%t", page, size, withDeleted)
}

func GetTeachersQuery(ctx *gin.Context) *gorm.DB {

	withDeleted := false
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 32)
	size, _ := strconv.ParseInt(ctx.Query("size"), 10, 32)
	offset := size * (page - 1)

	query := db.Where("").Limit(int(size)).Offset(int(offset))

	if !withDeleted {
		query.Where("deleted_at IN NOT NULL")
	}
	return query
}
