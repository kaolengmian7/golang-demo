package goquery

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	html := `<body>
				<div>DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

}

func TestMultiFind(t *testing.T) {
	html := `<body>
				<div>DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div,span").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func TestFind_IdSelector(t *testing.T) {
	html := `<body>
				<div id="div1">DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func TestFind_ClassSelector(t *testing.T) {
	html := `<body>
				<div>DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find(".name").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func TestFind_AttributeSelector(t *testing.T) {
	html := `<body>
				<div>DIV1</div>
				<div lang="zh">DIV2</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div[lang]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func TestFind_AttributeSelector_2(t *testing.T) {
	html := `<body>
				<div>DIV1</div>
				<div lang="zh">DIV2</div>
				<div lang="en">DIV3</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div[lang=zh]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func TestFind_ChildrenSelector(t *testing.T) {
	html := `<body>
				<div>DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("body>span").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func TestFind_ContentFilter_Contains(t *testing.T) {
	html := `<body>
				<div>DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div:contains(V2)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func TestFind_ContentFilter_Has(t *testing.T) {
	html := `<body>
				<span>SPAN1</span>
				<span>
					SPAN2
					<div>DIV</div>
				</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("span:has(div)").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

func TestAdd(t *testing.T) {
	html := `<body>
				<div>DIV</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(dom.Find("div").Text())
	log.Println(dom.Find("div").Add("span").Text())
}

func TestIs(t *testing.T) {
	html := `<body>
				<div>DIV</div>
				<span>SPAN</span>
			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(dom.Find("div").Add("span").Text())
}
