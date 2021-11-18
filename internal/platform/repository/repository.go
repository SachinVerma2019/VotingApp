package repositiory

import (
	"context"
	"voting-app/ent"
)

// Repositiory interface for declaring all db contract
type Repositiory interface {
	AuthenticateUser(ctx context.Context, email string, password string) (id int, name string)
	CreateUser(ctx context.Context, name string, email string, password string) (id int)
	CreatePoll(ctx context.Context, ownerId int, topic string, options []string) (int, bool)
	CastVote(ctx context.Context, voterId int, pollId int, option string) (int, bool)
	GetAllPolls(ctx context.Context) []*ent.Poll
	GetVotedPolls(ctx context.Context, userId int) []*VoteCount
	GetNotVotedPolls(ctx context.Context, userId int) []*ent.Poll
	ValidateOption(ctx context.Context, pollId int, option string) bool
	GetUsersForVoteOption(ctx context.Context, pollId int, option string) []*string
}

// router.POST("/poll/vote", api.Vote(&dbConnection))
// router.GET("/poll/:pollId", api.Poll(&dbConnection))
// router.GET("/poll/:pollId/:voteOption", api.Voters(&dbConnection))

// CREATE TABLE `polls` (
// 	`id` int NOT NULL AUTO_INCREMENT,
// 	`ownerid` int NOT NULL,
// 	`topic` varchar(300) NOT NULL,
// 	`options` json NOT NULL,
// 	`createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	`modifyTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `pollid_UNIQUE` (`id`),
// 	UNIQUE KEY `topic_UNIQUE` (`topic`),
// 	KEY `userid_idx` (`ownerid`),
// 	CONSTRAINT `userid` FOREIGN KEY (`ownerid`) REFERENCES `users` (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

// CREATE TABLE `results` (
// 	`id` int NOT NULL AUTO_INCREMENT,
// 	`userId` int NOT NULL,
// 	`pollId` int NOT NULL,
// 	`option` varchar(20) NOT NULL,
// 	`createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	`modifyTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`id`),
// 	KEY `userId_idx` (`userId`),
// 	KEY `results_userId_idx` (`userId`),
// 	KEY `I_userId_idx` (`userId`),
// 	KEY `FK_pollId` (`pollId`),
// 	CONSTRAINT `FK_pollId` FOREIGN KEY (`pollId`) REFERENCES `polls` (`id`),
// 	CONSTRAINT `Fk_userId` FOREIGN KEY (`userId`) REFERENCES `users` (`id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
