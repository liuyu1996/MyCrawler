package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil{
		panic(err)
	}

	result := ParseCityList(contents)
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	if len(result.Requests) != resultSize{
		t.Error("Requests size is wrong")
	}
	for i, url := range expectedUrls{
		if result.Requests[i].Url != url{
			t.Error("Requests is not expected url")
		}
	}
	
	if len(result.Items) != resultSize{
		t.Error("items size is wrong")
	}
	for i, city := range expectedCities{
		if result.Items[i].(string) != city{
			t.Error("Requests.item is not expected city")
		}
	}
}
