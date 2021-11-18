package repositiory

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"voting-app/ent"
	"voting-app/ent/poll"
	"voting-app/ent/result"
	"voting-app/ent/user"

	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	_ "github.com/go-sql-driver/mysql"
)

//VoteCount contains the data for each
type VoteCount struct {
	PollId int    `json:"pollid,omitempty"`
	Option string `json:"option,omitempty"`
	Count  int    `json:"count,omitempty"`
}

type MysqlRepo struct {
	client *ent.Client
}

var client *ent.Client

func init() {
	db, err := sql.Open("mysql", "root:Abcd!@34@tcp(localhost:3306)/VotingDb?parseTime=true")
	if err != nil {
		fmt.Println("Error in connecting with DB", err)
		// return nil, err
	}
	if db.Ping() != nil {
		fmt.Println("Success in pinging")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)
	client = ent.NewClient(ent.Driver(drv))
	// defer client.Close()
}

// NewEntClient .. return an ORM client to the caller
func NewEntClient() Repositiory {
	return &MysqlRepo{client: client}
}

//	CreateUser ...
func (repo *MysqlRepo) CreateUser(ctx context.Context, name string, email string, password string) (id int) {
	// ctx := context.Background()
	user, err := repo.client.User.
		Create().
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		fmt.Println("error in user create", err)
		return -1
	}
	fmt.Println("Created user ", user)
	return user.ID
}

//	CreateUser ...
func (repo *MysqlRepo) AuthenticateUser(ctx context.Context, email string, password string) (id int, name string) {
	var v []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	err := repo.client.User.
		Query().
		Where(user.And(
			user.Email(email),
			user.Password(password),
		)).
		Select(user.FieldID, user.FieldName).
		Scan(ctx, &v)

	if err != nil {
		fmt.Println("error in authenricate user ", err)
		return -1, ""
	}
	return v[0].Id, v[0].Name
}

func (repo *MysqlRepo) CreatePoll(ctx context.Context, ownerId int, topic string, options []string) (int, bool) {
	// tx, err := client.Tx(ctx)
	// if err != nil {
	// 	return -1, true
	// }

	poll, err := repo.client.Poll.
		Create().
		SetOwnerid(ownerId).
		SetTopic(topic).
		SetOptions(options).
		Save(ctx)

	if err != nil {
		fmt.Println("error in user create", err)
		// rerr := tx.Rollback()
		// if rerr != nil {
		// 	fmt.Println("error in rollback of transaction", rerr)
		// }
		return -1, true
	}

	// //This step can be avoided later
	// adminId := 0
	// for _, op := range options {
	// 	repo.CastVote(ctx, adminId, poll.ID, op)

	// }

	fmt.Println("Created user ", poll)
	return poll.ID, false
}

func (repo *MysqlRepo) CastVote(ctx context.Context, voterId int, pollId int, option string) (int, bool) {
	result, err := repo.client.Result.
		Create().
		SetUserid(voterId).
		SetPollid(pollId).
		SetOption(option).
		Save(ctx)

	if err != nil {
		fmt.Println("error in casting vote", err)
		return -1, true
	}
	fmt.Println("Created user ", result)
	return result.ID, false
}

// // rollback calls to tx.Rollback and wraps the given error
// // with the rollback error if occurred.
// func rollback(tx *ent.Tx, err error) error {
// 	if rerr := tx.Rollback(); rerr != nil {
// 		err = fmt.Errorf("%w: %v", err, rerr)
// 	}
// 	return err
// }

// func (repo *MysqlRepo) Test(ctx context.Context, voterId int, pollId int, option string) {
func (repo *MysqlRepo) ValidateOption(ctx context.Context, pollId int, option string) bool {
	result, err := repo.client.Poll.
		Query().
		Where(poll.And(
			func(s *entsql.Selector) {
				s.Where(sqljson.ValueContains(poll.FieldOptions, option))
			},
			poll.ID(pollId),
		)).Count(ctx)

	if err != nil {
		fmt.Println("error in user create", err)
		//return -1, true
		return false
	}
	fmt.Println("Option found ", result)
	return result == 1
}

func (repo *MysqlRepo) GetAllPolls(ctx context.Context) []*ent.Poll {
	polls, err := client.Poll.
		Query().
		All(ctx)

	if err != nil {
		return nil
	}
	return polls
}

func (repo *MysqlRepo) GetVotedPolls(ctx context.Context, userId int) []*VoteCount {
	/*
		-- count based on group by
			SELECT `pollId`, `option`, COUNT(pollId) as count from results
			WHERE pollId in (Select r.pollId from VotingDb.results r where r.userId = userId)
			Group BY `pollId`, `option`;
	*/
	var v []*VoteCount
	err := repo.client.Result.
		Query().
		Where(func(s *entsql.Selector) {
			t := entsql.Table(result.Table)
			s.Where(
				entsql.In(
					s.C(result.FieldPollid),
					entsql.Select(t.C(result.FieldPollid)).From(t).
						Where(entsql.EQ(t.C(result.FieldUserid), userId)),
				),
			)
		}).
		GroupBy(result.FieldPollid, result.FieldOption).
		Aggregate(ent.Count()).
		Scan(ctx, &v)

	if err != nil {
		return nil
	}
	return v
}

func (repo *MysqlRepo) GetNotVotedPolls(ctx context.Context, userId int) []*ent.Poll {
	/*
		Select * from VotingDb.polls p where p.id NOT IN
		(Select r.pollId from VotingDb.results r where r.userId = 1)
	*/
	polls, err := repo.client.Poll.
		Query().
		Where(func(s *entsql.Selector) {
			t := entsql.Table(result.Table)
			s.Where(
				entsql.NotIn(
					s.C(poll.FieldID),
					entsql.Select(t.C(result.FieldPollid)).From(t).Where(entsql.EQ(t.C(result.FieldUserid), userId)),
				),
			)
		}).
		All(ctx)

	if err != nil {
		return nil
	}
	return polls
}

func (repo *MysqlRepo) GetUsersForVoteOption(ctx context.Context, pollId int, option string) []*string {
	/*
		SELECT u.email from users u
		WHERE u.id in (SELECT userId from results where pollId = 1 AND `option` = 'java');
	*/

	var emails []*string
	err := repo.client.Debug().User.
		Query().
		Where(func(s *entsql.Selector) {
			t := entsql.Table(result.Table)
			p := entsql.And(
				entsql.EQ(t.C(result.FieldPollid), pollId),
				entsql.EQ(t.C(result.FieldOption), option),
			)
			s.Where(
				entsql.In(
					s.C(user.FieldID),
					entsql.Select(t.C(result.FieldUserid)).From(t).
						Where(p),
				),
			)

		}).
		Select(user.FieldEmail).
		Scan(ctx, &emails)

	if err != nil {
		fmt.Print(err)
		return nil
	}
	return emails
}

// func (repo *MysqlRepo) GetUsersForVoteOption_V2(ctx context.Context, pollId int, option string) []*string {
// 	/*
// 		SELECT u.email from users u
// 		WHERE u.id in (SELECT userId from results where pollId = 1 AND `option` = 'java');
// 	*/
// 	var emails []*string
// 	err := repo.client.User.
// 		Query().
// 		Where(func(s *entsql.Selector) {
// 			t := entsql.Table(result.Table)
// 			p := entsql.And(
// 				entsql.EQ(t.C(result.FieldPollid), pollId),
// 				entsql.EQ(t.C(result.FieldOption), option),
// 			)
// 			s.Where(entsql.Exists(entsql.Select().From(t).Where(p)))

// 			// s.Where(
// 			// 	entsql.In(
// 			// 		s.C(user.FieldID),
// 			// 		entsql.Select(t.C(result.FieldUserid)).From(t).
// 			// 			Where(entsql.EQ(t.C(result.FieldUserid), userId)),
// 			// 	),
// 			// )
// 		}).
// 		Select(user.FieldEmail).
// 		Scan(ctx, &emails)

// 	if err != nil {
// 		return nil
// 	}
// 	return emails
// }
