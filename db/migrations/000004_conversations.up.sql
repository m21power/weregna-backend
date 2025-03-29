CREATE TABLE conversations (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user1_id INT NOT NULL REFERENCES student(id) ON DELETE CASCADE,
  user2_id INT NOT NULL REFERENCES student(id) ON DELETE CASCADE,
  last_message_text TEXT,
  last_message_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX conversations_user_pair_unique
ON conversations (LEAST(user1_id, user2_id), GREATEST(user1_id, user2_id));