package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}
	pasetoMaker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return pasetoMaker, nil
}

// CreateToken tạo một PASETO token cho username với thời gian hiệu lực
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		return "", err
	}
	return token, err
}

// VerifyToken kiểm tra tính hợp lệ của PASETO token
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	if err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil); err != nil {
		return nil, err
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil

}

/*
Trong đoạn mã token, err := maker.paseto.Encrypt(maker.symmetricKey, data, nil),
tham số được truyền vào hàm Encrypt có thể được giải thích như sau:
	-maker.symmetricKey: Đây là khóa đối xứng mà bạn sẽ sử dụng để mã hóa dữ liệu.
		Khóa này cần được bảo mật và không được công khai, vì nó cho phép giải mã lại dữ liệu sau này.

	-data: Đây là dữ liệu mà bạn muốn mã hóa.
		Dữ liệu này có thể là bất kỳ thông tin nào mà bạn muốn bảo vệ, chẳng hạn như thông tin người dùng hoặc các thông tin nhạy cảm khác.

	-nil: Tham số cuối cùng là nil, đại diện cho "optional" context hoặc metadata.
		Trong nhiều thư viện mã hóa, tham số này thường được sử dụng để truyền thêm thông tin tùy chọn cho quá trình mã hóa,
		chẳng hạn như các thông tin liên quan đến tiêu đề hoặc các tùy chọn mã hóa khác.
		Khi bạn truyền nil, điều này có nghĩa là bạn không muốn thêm bất kỳ thông tin tùy chọn nào vào quá trình mã hóa.
		Một số thư viện mã hóa có thể cho phép bạn thêm metadata vào token, nhưng nếu không có nhu cầu đó, bạn có thể để nó là nil.

Tóm lại, bạn đang sử dụng hàm Encrypt để mã hóa dữ liệu với khóa đối xứng, và không cung cấp thêm thông tin tùy chọn nào trong tham số cuối cùng.

You said: thêm metadata vào token là như thế nào
ChatGPT said:
ChatGPT
Thêm metadata vào token là việc bạn đưa thêm thông tin phụ trợ vào token để lưu trữ các dữ liệu bổ sung cùng với payload chính.
Metadata thường là những thông tin mà không phải là phần chính yếu của token, nhưng lại cần thiết cho các tác vụ liên quan đến xác thực hoặc kiểm soát quyền truy cập.
Trong trường hợp này, bạn có thể thêm metadata vào một token JWT hoặc PASETO. Ví dụ, trong JWT, bạn có thể thêm các trường bổ sung vào phần claims của token.

Ví dụ thêm metadata vào token JWT với MapClaims:

	go
	Copy code
	claims := jwt.MapClaims{
	"id":         payload.ID,
	"username":   payload.Username,
	"issued_at":  payload.IssuedAt,
	"expired_at": payload.ExpiredAt,
	"role":       "admin",           // Thêm metadata 'role'
	"ip_address": "192.168.1.1",     // Thêm metadata 'ip_address'
	}
Trong ví dụ trên:

"role": là metadata chỉ rõ vai trò của người dùng (admin, editor, viewer...).
"ip_address": là metadata chứa địa chỉ IP của người dùng khi tạo token.
Các trường metadata này có thể được tùy ý thêm vào và truy xuất lại khi cần thiết trong quá trình xác thực token. Thông tin này không ảnh hưởng đến payload chính mà chỉ cung cấp thông tin bổ trợ.*/
