-- Insert some test data
INSERT INTO users (username, name, email)
VALUES
  ('ary82','aryan goyal','ary82@mail.com'),
  ('test', 'test account', 'test@mail.com'),
  ('demo', 'demo throwaway', 'demo@mail.com');

INSERT INTO stashes (title, body, owner_id, is_public)
VALUES
  ('Top places to learn golang', 'The best places to learn golang', '1', 'true'),
  ('Postgresql resources', 'This contains resources to learn PostgreSQL database', '2', 'false'),
  ('Top places to learn hexagonal pattern', 'If I knew those places, I wouldve learnt it', '1', 'true');

INSERT INTO links (url, stash_id)
VALUES
  ('https://go.dev', '1'),
  ('https://gobyexample.com/', '1'),
  ('https://www.postgresql.org/', '2'),
  ('https://www.postgresql.org/docs/', '2');

INSERT INTO comments (author, body, stash_id)
VALUES
  ('1', 'Go.dev has everything', '1'),
  ('2', 'great list 10/10', '1'),
  ('3', 'Good one', '1'),
  ('1', 'Postgresql official docs the best huh', '2');

INSERT INTO stars (user_id, stash_id)
VALUES
  ('1', '1'),
  ('2', '1'),
  ('2', '2');
