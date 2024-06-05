-- +goose Up
-- +goose StatementBegin
CREATE TABLE Role(
    id SERIAL PRIMARY KEY,
    name TEXT
);

INSERT INTO Role(name)
VALUES ('User'),
('Moderator'),
('Admin');

CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    username TEXT,
    email TEXT,
    password TEXT,
    roleId int,
    createdAt TIMESTAMP,
    CONSTRAINT fk_role
        FOREIGN KEY(roleId)
            REFERENCES Role(id)
);

CREATE TABLE Game(
    id SERIAL PRIMARY KEY,
    name TEXT,
    createdAt TIMESTAMP
);

CREATE TABLE Score(
    id SERIAL PRIMARY KEY,
    playerId int,
    gameId int,
    approverId int,
    score int,
    createdAt TIMESTAMP,
    approvedAt TIMESTAMP,
    CONSTRAINT fk_player
        FOREIGN KEY(playerId)
            REFERENCES Users(id),
    CONSTRAINT fk_game
        FOREIGN KEY(gameId)
            REFERENCES Game(id),
    CONSTRAINT fk_approver
        FOREIGN KEY(approverId)
            REFERENCES Users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Role;
DROP TABLE User;
DROP TABLE Game;
DROP TABLE Score;
-- +goose StatementEnd
