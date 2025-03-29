CREATE TABLE messages (
  id BIGSERIAL PRIMARY KEY,
  conversation_id UUID NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
  sender_id INT NOT NULL REFERENCES student(id) ON DELETE CASCADE,
  content TEXT NOT NULL,
  status message_status DEFAULT 'unseen',
  sent_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);