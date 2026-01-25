INSERT INTO users (email, name, password) VALUES 
('user1@test.com', 'Амбар Грозный', 'password123'),
('user2@test.com', 'Мария Петрова', 'password123'),
('user3@test.com', 'Петя Мариновна', 'password123'),
('user4@test.com', 'Иван Иваныч', 'password123'),
('user5@test.com', 'Бегемот Красавыч', 'password123');

DO $$
BEGIN
    FOR i IN 1..30 LOOP
        INSERT INTO resources (name, type, is_available) 
        VALUES ('Столик №' || i, 'table', true);
    END LOOP;
    
    FOR i IN 1..30 LOOP
        INSERT INTO resources (name, type, is_available) 
        VALUES ('Доктор ' || i, 'doctor', true);
    END LOOP;
    
    FOR i IN 1..40 LOOP
        INSERT INTO resources (name, type, is_available) 
        VALUES ('Переговорка №' || i, 'meeting_room', true);
    END LOOP;
END $$;