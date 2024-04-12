type stash = {
  id: number;
  author: string;
  author_id: number;
  title: string;
  body: string;
  stars: number;
  created_at: string;
};

type stashDetail = {
  id: number;
  author: string;
  author_id: number;
  title: string;
  body: string;
  stars: number;
  is_public: boolean;
  created_at: string;
  links: link[];
  comments: comment[];
};

type link = {
  id: int;
  url: string;
  comment: string;
};

type comment = {
  id: int;
  author_id: int;
  author: string;
  body: string;
  created_at: string;
};

type user = {
  id: number;
  username: string;
  stars: number;
  picture: string;
  created_at: string;
  public_stashes: stash[];
};
