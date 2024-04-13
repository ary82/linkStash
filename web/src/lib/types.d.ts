type stash = {
  id: number;
  author: string;
  author_id: number;
  title: string;
  body: string;
  stars: number;
  created_at: string;
};

type stashDetail = stash & {
  is_public: boolean;
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
  picture: string;
};

type userDetail = user & {
  stars: number;
  created_at: string;
  public_stashes: stash[];
};
