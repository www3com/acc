package email

import "testing"

func TestSendMail(t *testing.T) {
	body := `
			<html>
				<body>
					<h1>您的验证码是1010</h1>
				</body>
			</html>
	`

	if err := SendMail("jason.wang@vdx.com", "注册成功通知", body); err != nil {
		t.Errorf("邮件发送失败, 原因： %v", err)
	}

}
