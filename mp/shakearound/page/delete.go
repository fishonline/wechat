// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com), magicshui(shuiyuzhe@gmail.com), Harry Rong(harrykobe@gmail.com)

package page

import (
	"gopkg.in/chanxuehong/wechat.v1/mp"
)

// 删除页面
func Delete(clt *mp.Client, pageIds []int64) (err error) {
	request := struct {
		PageIds []int64 `json:"page_ids,omitempty"`
	}{
		PageIds: pageIds,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
