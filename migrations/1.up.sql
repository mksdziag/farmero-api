CREATE TABLE articles (
  id UUID PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  description TEXT NOT NULL,
  cover TEXT NOT NULL
);

CREATE TABLE categories (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  key TEXT NOT NULL
);

CREATE TABLE tags (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  key TEXT NOT NULL
);

CREATE TABLE articles_tags (
  id UUID PRIMARY KEY,
  article_id UUID NOT NULL,
  tag_id UUID NOT NULL,
  FOREIGN KEY (article_id) REFERENCES articles(id),
  FOREIGN KEY (tag_id) REFERENCES tags(id)
);

CREATE TABLE articles_categories (
  id UUID PRIMARY KEY,
  article_id UUID NOT NULL,
  category_id UUID NOT NULL,
  FOREIGN KEY (article_id) REFERENCES articles(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);