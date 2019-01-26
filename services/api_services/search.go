/**
*@Author: haoxiongxiao
*@Date: 2019/1/26
*@Description: CREATE GO FILE api_services
*/
package api_services

type SearchRequestParams struct {
	CommonRequestParams
	KeyWord      string `json:"keyWord"`      //查询关键字，酒店名称、位置、品牌等
	Page         string `json:"page"`         //页码
	CityName     string `json:"cityName"`     //城市
	IData        string `json:"iData"`        //入住时间，格式为：YYYY-MM-DD（默认2天后）
	OutData      string `json:"outData"`      //离开时间，格式为：YYYY-MM-DD（默认3天后）
	SortCode     string `json:"sortCode"`     //排序规则(默认1.推荐值排序) 1、推荐值降序 2、起价升序 3、起价降序 6、装修时间排序
	ReturnFilter string `json:"returnFilter"` //是否返回聚合筛选条件,0:否,1:是。注意：returnFilter=1时搜索性能较差，尽量设置returnFilter=0
	Star         string `json:"star"`         //星级 TWO:二星级, THREE:三星级, FOUR:四星级, FIVE:五星级, BUDGET:经济型, CONFORT:舒适型, HIGHEND:高档型, LUXURY:豪华型【多个以逗号:‘,’分隔】
	Feature      string `json:"feature"`      //品牌:通过搜索结果反向聚合
	MinPrice     string `json:"minPrice"`     //房价最低价
	MaxPrice     string `json:"maxPrice"`     //房价最高价
	Facility     string `json:"facility"`     //设施:通过搜索结果反向聚合
	HotelLabels  string `json:"hotellablels"` //特色 1、温泉 3、休闲度假 4、购物便捷 5、客栈民宿 6、青年旅舍 7、精品酒店 8、亲子时刻
}

type SearchResponseParams struct {
	RetCode        string `json:"retCode"`
	Remark         string `json:"remark"`
	FilterInfo     string `json:"filterInfo"`
	Data           string `json:"data"`
	CityName       string `json:"cityName"`
	FilterName     string `json:"filterName"`
	FilterId       string `json:"filterId"`
	FilterProsList string `json:"filterProsList"`
	SubName        string `json:"subName"`
	SubId          string `json:"subId"`
	SubProsList    string `json:"subProsList"`
	Lat            string `json:"lat"`
	Lng            string `json:"lng"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	HotelList      string `json:"hotelList"`
	Count          string `json:"count"`
	CityCode       string `json:"cityCode"`
	HotelId        string `json:"hotelId"`
	StarName       string `json:"starName"`
	Address        string `json:"address"`
	Location       string `json:"location"`
	Facilities     string `json:"facilities"`
	StartPrice     string `json:"startPrice"`
	Pictures       string `json:"pictures"`
}


