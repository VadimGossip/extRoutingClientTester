package model

type Response struct {
	Ok   bool         `json:"ok"`
	Body ResponseBody `json:"result"`
}

type ResponseBody struct {
	Anumber  string   `json:"anumber"`
	Bnumber  string   `json:"bnumber"`
	ErrCode  int      `json:"errcode"`
	Flags    int      `json:"flags"`
	Mark     int      `json:"mark"`
	RChecked int      `json:"rchecked"`
	Res      int      `json:"res"`
	RMatched int      `json:"rmatched"`
	Status   string   `json:"status"`
	Tags     []string `json:"tags"`
}
