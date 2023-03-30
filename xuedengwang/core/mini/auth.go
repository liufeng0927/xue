package mini

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"xuedengwang/core/httpx"
	"xuedengwang/core/stringx"

	"net/url"
)

// Session 登录凭证校验
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
func (c *Client) Session(code string) (*ResponseCode2Session, error) {

	result := &ResponseCode2Session{}

	uri := stringx.Concat(c.BaseUri, "sns/jscode2session?")

	params := url.Values{
		"appid":      []string{c.AppId},
		"secret":     []string{c.AppSecret},
		"js_code":    []string{code},
		"grant_type": []string{"authorization_code"},
	}

	resp, err := httpx.Get(c.ctx, uri, params)
	logx.Infof("output:%s", string(resp))

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if result.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%d  ErrMsg:%s", result.ErrCode, result.ErrMSG))
	}

	return result, err
}
