CREATE TABLE IF NOT EXISTS Product (
    ID VARCHAR(255) PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Price INTEGER NOT NULL,
    Unit_ID INTEGER NOT NULL,
    Category_ID INTEGER NOT NULL,
    Created_At TIMESTAMP NOT NULL,
    Updated_At TIMESTAMP NOT NULL,
    Deleted_At TIMESTAMP NULL,
    FOREIGN KEY (Unit_ID) REFERENCES Unit(ID),
    FOREIGN KEY (Category_ID) REFERENCES Category(ID)
);

INSERT INTO Product (ID, Name, Price, Unit_ID, Category_ID, Created_At, Updated_At)
VALUES 
    ('01H7D8E4K5TX7P9ZP6G3V0A1W2', 'Leite UHT 1L', 3, 1, 1, NOW(), NOW()),
    ('01H7D8E4L6JH0Q1ZP7G4V1B2X3', 'Arroz 5kg', 20, 2, 2, NOW(), NOW()),
    ('01H7D8E4M7TX2R3ZQ8G5V2C3Y4', 'Feijão Preto 1kg', 8, 2, 2, NOW(), NOW()),
    ('01H7D8E4N8CWR4RZP9G6V3A4Z5', 'Macarrão Espaguete 500g', 4, 4, 4, NOW(), NOW()),
    ('01H7D8E4P9HX7Q5ZP0G7V4B6X7', 'Óleo de Soja 900ml', 6, 1, 1, NOW(), NOW()),
    ('01H7D8E4Q0PT8R6ZP1G8V5C7Y8', 'Açúcar Cristal 1kg', 5, 2, 4, NOW(), NOW()),
    ('01H7D8E4R1HX9P7ZP2G9V6B8X9', 'Café em Pó 500g', 7, 2, 4, NOW(), NOW()),
    ('01H7D8E4S2JH0Q8ZP3G0V7C9Y0', 'Papel Toalha 2 Rolos', 8, 3, 3, NOW(), NOW()),
    ('01H7D8E4T3TX1R9ZP4G1V8B0X1', 'Sabão em Pó 1kg', 6, 3, 3, NOW(), NOW()),
    ('01H7D8E4U4CWR2PZP5G2V9C1Y2', 'Desinfetante 500ml', 4, 3, 3, NOW(), NOW()),
    ('01H7D8E4V5HX3Q1ZP6G3V0B2X3', 'Molho de Tomate 340g', 3, 2, 1, NOW(), NOW()),
    ('01H7D8E4W6PT4R2ZP7G4V1C3Y4', 'Sardinha em Lata 180g', 2, 5, 1, NOW(), NOW()),
    ('01H7D8E4X7JH5Q3ZP8G5V2B4X5', 'Atum em Lata 170g', 2, 5, 1, NOW(), NOW()),
    ('01H7D8E4Y8TX6R4ZP9G6V3C5Y6', 'Maionese 500g', 5, 4, 1, NOW(), NOW()),
    ('01H7D8E4Z9CWR7PZP0G7V4B6X7', 'Ketchup 400g', 4, 4, 1, NOW(), NOW()),
    ('01H7D8F500HX8Q9ZP1G8V5C7Y8', 'Biscoito Cream Cracker 400g', 3, 4, 1, NOW(), NOW()),
    ('01H7D8F511PT9R0ZP2G9V6B8X1', 'Queijo Mussarela 200g', 8, 6, 6, NOW(), NOW()),
    ('01H7D8F522JH1Q1ZP3G0V7C9Y0', 'Presunto 150g', 6, 6, 6, NOW(), NOW()),
    ('01H7D8F533TX2R2ZP4G1V8B0X1', 'Iogurte Natural 170g', 3, 6, 6, NOW(), NOW()),
    ('01H7D8F544CWR3PZP5G2V9C1Y2', 'Manteiga 200g', 7, 6, 6, NOW(), NOW()),
    ('01H7D8F555HX4Q4ZP6G3V0B2X3', 'Refrigerante 2L', 5, 1, 1, NOW(), NOW());