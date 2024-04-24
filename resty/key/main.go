package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	resp, err := client.R().SetHeaders(map[string]string{
		"cookie":         getcookie(),
		"referer":        "https://setup.office.com/EnterPin?ctid=" + tid(),
		"content-type":   "application/json; charset=UTF-8",
		"origin":         "https://setup.office.com",
		"user-agent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		"x-ctid":         tid(),
		"sec-fetch-mode": "cors",
		"Content-Length": "158",
		"host":           "setup.office.com",
		"Connection":     "keep-alive",
		"sec-fetch-site": "same-origin",
		"sec-fetch-des":  "empty",
	}).SetBody(map[string]string{
		"Channel":             "",
		"ClientIp":            "36.110.222.130",
		"ClientTransactionId": tid(),
		"Token":               productKey(),
		"TokenHash":           "",
	}).Post("https://setup.office.com/home/submitpinmodern")

	if err != nil {
		fmt.Println("---请求错误:", err)
		return
	}

	s := resp.String()
	fmt.Println(s)

}

func getcookie() string {
	return "MUID=1BD1D1236487647137C3C157654F655D; omkt=zh-CN; MSCC=cid=pemqc5t7328yt4bn38grwvj1-c1=2-c2=2-c3=2; .AspNetCore.Cookies=CfDJ8Iny2OQBG0BDp31KFnvRAAU9gWe01qtvcqOyiDwgF4j6t2PmRHBpnt8CqZDayZm0cWycSIifPphedMO9bG8--JkXwCtlrKFpAqKddOX0I6UbPoKLgLPflC-M1lcE6t_nYUYKomL4eluy2yrK2TsvhlCPRJ6uE6WqaFdWkz-O7UxaLGWx4yB3UP0lMxkAFHh4_csKpX79UNQ1daKrsEw8OB32U_7tCHeqN1KGNbwwzTzwC4JoNVQGNgMahuJCLStiLbuCbYMUGPBgSiO4JyWf2AkkXy912kKJyqDfAd7nepU0d4ZAZfBsOlnuqN1JLCK_C74pSqcMlRJidIswFz8TyYpQ-7YrM55BdXJGURzTz1yFMrtyBUhzOs8d-x9Twlr5HN16986xWMDOs5SON4uUiQggeHYE0TwNQdwMAuYQJe9GVsZeTLicD3NupXZ3Pmbrj5j3OahG1NAmfW3QKe0pJqV3CXHsdc448cfR93Ui_d0gNIgs5c-_z2-eGZ4RCzo5WL30hxb3MkJAiV24IJ7WRxKKAR2CoAIfRrf-66ciBxbfSlmQ1TJkentCW0UIRUO8Hr1ktANCQBf9tjsDusLW6PcaL35dzkjvIryCDA3z81kp6A7i_oRqpUZvYVsNK257_vS_iw2-3S3P1p_xUjARJ77C2dNkKU3v-gZRBTrde4HxmIauWULN_GPQ63OCCIFapQXBgLAydcWcBqs4L7sbWTdmrU3J6qjTMpBnm97YXLdC4905WNqZaSN-fD8RzE9TfLvi_ma6WMOGSKmvkELDSIxRFC9EmzwXHW3T9ZKaThI2aT0KL_McjaJwqcP4Hn1sKM5TZVoYXSxus3AJO3i4dHhnR01-5FeWreD2Z_F7hoSNoKz2K8Q5VctIXCogQUaPf_YtcO4JF1E_N07mXO0xKXgY7MFMUeERAO2miY3V07c3wqG3ImaN_iJSWWS0gTHVbGqNfmmuXaA3LEt9YmWHTLQgph9xdV6GmsmmNfVZmega4eDkRQM2jqfSFyEEcHatOXggs4qxvXawfIo33lpU-laUdAetf5OXZVTBNkgH3NZPaVutN3XaISpJbI3SE5JBUfot_jEghs2bZQ1SMcTVTS0hP0yT0lhiBzJmFd64-Kkimka1lcNxlkgqE4MBmRgYoS_nbZGR9gRyPev1KZ4SskJLJcg8iRwx1A8SagI9jrRjKV-no5C3h-nJxPxsEjRUKQ-tPYLx_pLuoLCIrO_7tdFmoW_Hc0a_g6d8o332gf1YIsQT9801zDQ2K_yTHc6JIKJAT9PHN3fB-cGUkQyF10Awqj50dDQyZp53WoGZo7iATSYoKJgoX9To97h3Ii8wJqmWwiw1pmKfEP3cCn4OUD9usiiiIRvNPaxUe99DliVccuxqD9Fu-huad03Cp1ShU_GuEDDqd4srHRyA2IGcA_JDS51GZFcfO0Iy-EDjZp72olOrS_AS0FulJKgBbbQGxTwExV-Y0DAUzENBFNFUZQF0dYHsKz_8QxIvPQIpCujDnhNZ6mZQGPx-czqYvqdFuJMH-kADGf_BLfpRDBEf8Q7y6DslrHlMv17qWvWyoEYmgeQItvdMRnQISaM_RY7OnGsdLonc9JtY-vWzdyvFJuC2LNAmeozg2j5F-n6uNyS2BkoQFC5aTijexOq1-x8uJzpd91v8HKyQgLE3r3xFY1_ay2ceTP2zFk1ozrZxhsQUoWCbqcp4FBO91tZhIddqOru_GWlr7Rju-N_BjmDSRXMJ3I4Xn1j0_ybECZ31VFUBRbNbPFNkOo1yME2JAnWH-qf29kzCpUzqy1ZKw6YYhI_5sbgQ8gtsycFFVQYA5XDTBoyDt65879v7jWfbidLKda2eouQalAwRiKv-lImmrUmem3uOS1bT5KpoeY28LFCHP6J0Ddb-Ab2yLepKgoG5Yhwhk-filKQfo1u28mH1THj89SOuXIAwxa70sPpsOhIOIfFvVSOPIN2zMPof1Ro5QZOQqHO11_ZcPdvJM4IJCBuNXQUz-eyCurcbAfuVhloulXU7t0AKg88AuTrCVzZFvjBEBnSDgWFpFr3brtpkNpZ-hLndxWHdPBnBgqCGVwW3LZDv6iK6Z1A_5wWxpu373_XH1MlkGt7qssKw0j1HCh7JkPjc7QZ6ThCbLfKNlx1XLHXvMtktW8pr-JwT8GPZgEeN6ej0Gft1yvyO8odrkWGVxhTAsSYmnSDEPtW3odF0tukYJ3Zv24K3cwLMIXa0K2miw0wXOWKL-RyuaLYjgu5FOb9I9LENybWS2xD66WR1vuxk86aQLnP8p1zKrLITzMtH1NcEaoLecC4_FLYRYI8OCoWAouTBmX7CBW4BCkY0uisZ_vQz-xn3dRE2E0ir5ivlYPyo2lAzZ0gdHaJqihXAHza8VRakbOX7csJTn9VM8Nl7quMrRNtmvdbvV6ypoY2UlS4CYUG7Lv5lyz4Q7vwz903F6OeAfnkcF2SBPADOwtYmcwPWbl4ZsXqgAmo2hmsdLlXeF89ROIB2soxL4KNpKBsVPM71a1iy35vrpwUkWEyrv2FM1p0lNIlTkbN0qVeF9nTNtkqHci9QhuvEuwgfbG_6DYSlu_lFFguLZMG59MdJN1-e27VbC203ANZOwKrXRd81XRCQnEy8KK-v-_2iyXi_hJ9Bh2gl-Zko_xhRzOd85_x_I9UZnDAVfjUrKhvClWYgDVdLStOUNUJNbQfj-lujxjtbe4LPjSP3mDF4HWZRga2kU-j8bQELE_jk8I9_4dDdA_PgVxykRaOJSTdGV6xd6HdK6NhVzBzDnTs1; .AspNetCore.Session=CfDJ8Iny2OQBG0BDp31KFnvRAAU1cGPKuDXIL%2Bop%2BZkse7Q4swMuUxMmwCFN6ZOe4Mzsm5Bm6TfdgkedU1VfIll20rChFUEnE0tUJhryUmCX6K2WWtYnqDt61sIUWgrGk567qNIRFbOz86eloGF2jAW6HS5Lj%2B6QA%2BQ%2FN53XobRbCQ16"
}

func tid() string {
	return "b207419c-8d57-4b71-a310-d89755070e97"
}

func productKey() string {
	// TokenAlreadyRedeemed
	return "3QN83-W4KC6-T39D3-HG7C7-QPBYF"
}
