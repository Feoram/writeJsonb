CREATE TABLE data (
                      id SERIAL PRIMARY KEY,
                      payload JSONB,
                      created_at TIMESTAMP DEFAULT now()
);
