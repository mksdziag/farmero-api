package articles

func GetArticlesByCategory(category string) ([]Article, error) {
	var articles = make([]Article, 0)

	for _, article := range articlesList {
		for _, cat := range article.Categories {
			if cat == category {
				articles = append(articles, article)
				break
			}
		}
	}

	return articles, nil
}

func GetArticles() ([]Article, error) {
	articles := articlesList[0:6]

	return articles, nil
}

func GetArticle(id string) (Article, error) {
	var found = Article{}

	for _, article := range articlesList {
		if article.ID == id {
			found = article
		}
	}

	return found, nil
}
