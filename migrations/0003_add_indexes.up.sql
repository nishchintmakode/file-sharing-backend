-- Index on user_id for faster retrieval of user-specific files
CREATE INDEX idx_files_user_id ON files(user_id);

-- Index on filename for search optimization
CREATE INDEX idx_files_filename ON files(filename);

-- Index on upload_date for filtering by date
CREATE INDEX idx_files_upload_date ON files(upload_date);