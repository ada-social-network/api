package handler

// ListUsersByPromo respond a list of Users from a promo
/* func TestListUsersByPromo(t *testing.T) {
	return func(c *gin.Context) {
		users := &[]models.Promo{}

		result := db.Find(&users)
		if result.Error != nil {
			httpError.Internal(c, result.Error)
			return
		}

		c.JSON(200, users)
	}
} */

/* func TestListPromoUsers(t *testing.T) {
	db := commonTesting.InitDB(&models.Promo{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.Promo{})

	ListPromoUsers(db)(ctx)

	got := &[]models.Promo{}
	_ = json.Unmarshal(res.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("ListPromoUsers response should not be empty")
	}
}
*/
