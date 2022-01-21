package handler

/* func TestListPostHandler(t *testing.T) {
	db := commonTesting.InitDB(&models.Post{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.Post{})

	ListPostHandler(db)(ctx)

	got := &[]models.Post{}
	_ = json.Unmarshal(res.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("ListPostHandler response should not be empty")
	}
}

func TestCreatePostHandler(t *testing.T) {
	type args struct {
		post models.Post
		user *models.User
	}

	type want struct {
		count      int64
		statusCode int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid post",
			args: args{
				user: &models.User{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
				post: models.Post{
					Base:    models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b")},
					Title:   "foo bar",
					Content: "lorem ipsum",
				},
			},
			want: want{
				count:      1,
				statusCode: 200,
			},
		},
		{
			name: "invalid post",
			args: args{
				user: &models.User{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
				post: models.Post{Content: "l"},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
		{
			name: "missing user",
			args: args{
				post: models.Post{Content: "lorem ipsum"},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
		{
			name: "empty post",
			args: args{
				user: &models.User{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
				post: models.Post{},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := commonTesting.InitDB(&models.Post{})
			res, ctx, _ := commonTesting.InitHTTPTest()

			// set user in the current context
			ctx.Set(middleware.IdentityKey, tt.args.user)

			commonTesting.AddRequestWithBodyToContext(ctx, tt.args.post)

			CreatePostHandler(db)(ctx)

			post := &models.Post{}
			_ = json.Unmarshal(res.Body.Bytes(), post)

			if res.Code != tt.want.statusCode {
				t.Errorf("CreatePostHandler want:%d, got:%d", tt.want.statusCode, res.Code)
			}

			tx := db.First(&models.Post{}, "id = ?", post.ID)
			if tx.RowsAffected != tt.want.count {
				t.Errorf("CreatePostHandler want:%d, got:%d", tt.want.count, tx.RowsAffected)
			}
		})
	}
}

func TestDeletePostHandler(t *testing.T) {
	db := commonTesting.InitDB(&models.Post{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.Post{
		Base:    models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b")},
		Content: "lorem ipsum",
	})

	ctx.Params = gin.Params{
		{
			Key:   "id",
			Value: "80a08d36-cfea-4898-aee3-6902fa562f0b",
		},
	}

	DeletePostHandler(db)(ctx)

	if res.Code != 204 {
		t.Errorf("DeletePostHandler want:%d, got:%d", 204, res.Code)
	}

	tx := db.First(&models.Post{}, "id = ?", 123)
	if tx.RowsAffected != 0 {
		t.Errorf("DeletePostHandler Post should be deleted")
	}
}

func TestGetPostHandler(t *testing.T) {
	db := commonTesting.InitDB(&models.Post{})

	type args struct {
		post   *models.Post
		params gin.Params
	}

	type want struct {
		code int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "nominal",
			args: args{
				post: &models.Post{
					Base: models.Base{
						ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b"),
					},
				},
				params: gin.Params{
					{
						Key:   "id",
						Value: "80a08d36-cfea-4898-aee3-6902fa562f0b",
					},
				},
			},
			want: want{
				code: 200,
			},
		},
		{
			name: "not found",
			args: args{
				post: &models.Post{
					Base:    models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b")},
					Content: "125",
				},
				params: gin.Params{
					{
						Key:   "id",
						Value: "99999999-9999-9999-9999-999999999999",
					},
				},
			},
			want: want{
				code: 404,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, ctx, _ := commonTesting.InitHTTPTest()

			db.Create(tt.args.post)

			ctx.Params = tt.args.params

			GetPostHandler(db)(ctx)

			if res.Code != tt.want.code {
				t.Errorf("GetPostHandler want:%d, got:%d", tt.want.code, res.Code)
			}
		})
	}
}
*/
