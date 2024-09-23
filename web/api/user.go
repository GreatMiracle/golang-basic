package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	db "simpleBank/db/sqlc"
	"simpleBank/util"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"fullName"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName, //i can see pwd
		Email:          req.Email,
	}

	_, err = server.store.GetUser(ctx, req.Username)
	if err == nil {
		// Nếu không có lỗi, có nghĩa là người dùng đã tồn tại
		ctx.JSON(http.StatusConflict, errorResponseStr("User already exists"))
		return
	}
	if !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		return
	}

	res := userResponse{
		Username:          user.Username,
		FullName:          user.FullName, //i can see pwd
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, res)
}

/*pq.Error là một struct đại diện cho lỗi từ driver PostgreSQL (pq) trong Go. Nó chứa thông tin chi tiết về lỗi, bao gồm mã lỗi (error code),
mô tả lỗi và một số thông tin khác liên quan đến lỗi.
2. Sử Dụng errors.As
Trong đoạn mã, errors.As(err, &pqErr) được sử dụng để kiểm tra xem biến err có phải là một lỗi của loại *pq.Error hay không.
Nếu có, pqErr sẽ chứa thông tin của lỗi đó.
3. Kiểm Tra Mã Lỗi
Sau khi xác định rằng err là một lỗi PostgreSQL, bạn có thể kiểm tra mã lỗi cụ thể thông qua thuộc tính pqErr.Code.Name().
Mã lỗi này cho phép bạn xác định loại lỗi cụ thể mà bạn gặp phải.
4. Xử Lý Lỗi Đặc Biệt: unique_violation
Trong đoạn mã, có một phần kiểm tra mã lỗi unique_violation:
go

Copy
if errors.As(err, &pqErr) {
switch pqErr.Code.Name() {
case "unique_violation":
ctx.JSON(http.StatusForbidden, errorResponse(err))
return
}
}
Giải thích:
Nếu mã lỗi là unique_violation, điều này có nghĩa là bạn đang cố gắng chèn một bản ghi mà vi phạm ràng buộc duy nhất
(ví dụ: username hoặc email đã tồn tại trong cơ sở dữ liệu).
Trong trường hợp này, phản hồi trả về cho client sẽ có mã trạng thái HTTP 403 (Forbidden) cùng với thông tin lỗi.*/
