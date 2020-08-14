package repositories

import (
	"context"
	"fmt"

	"github.com/bagus2x/sirius-c/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PaperRepository -
type PaperRepository struct {
	ctx context.Context
	db  *mongo.Database
}

// NewPaperRepository -
func NewPaperRepository(ctx context.Context, db *mongo.Database) domain.PaperRepository {
	return &PaperRepository{ctx, db}
}

// InsertOne -
func (pr PaperRepository) InsertOne(p *domain.Paper) (err error) {
	p.ID = primitive.NewObjectID()
	p.Results = []domain.Result{}
	for i, v := range p.Questions {
		v.QstID = i
	}
	_, err = pr.db.Collection("paper").InsertOne(pr.ctx, p)
	return err
}

// FindByID -
func (pr PaperRepository) FindByID(id string) (res *domain.Paper, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.Paper{}, err
	}
	err = pr.db.Collection("paper").FindOne(pr.ctx, bson.M{"_id": _id}).Decode(&res)
	if err != nil {
		return &domain.Paper{}, err
	}
	return res, err
}

// GetPaper - without key
func (pr PaperRepository) GetPaper(id string) (res []map[string]interface{}, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	match := bson.M{"$match": bson.M{"_id": _id}}
	project := bson.M{"$project": bson.M{
		"_id": "$_id",
		"detail": bson.M{
			"title":       "$title",
			"description": "$description",
			"subject":     "$subject",
			"endAt":       "$endAt",
			"startFrom":   "$startFrom",
		},
		"questions": "$questions",
	}}
	project2 := bson.M{"$project": bson.M{"questions.key": 0, "questions.solution": 0}}
	curs, err := pr.db.Collection("paper").Aggregate(pr.ctx, bson.A{match, project, project2})
	if err != nil {
		return nil, err
	}
	err = curs.All(pr.ctx, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}

// PushExamResult -
func (pr PaperRepository) PushExamResult(id string, rst *domain.Result) (resid string, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	rst.ID = primitive.NewObjectID()
	res, err := pr.db.Collection("paper").UpdateOne(pr.ctx, bson.M{"_id": _id}, bson.M{"$push": bson.M{"results": rst}})
	if err != nil {
		return "", err
	}
	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return "", domain.ErrPaperIDNotFound
	}
	return rst.ID.Hex(), err
}

// GetExamResult -
func (pr PaperRepository) GetExamResult(id string, resid string) (res []map[string]interface{}, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	_resid, err := primitive.ObjectIDFromHex(resid)
	// Mom come pick me up, i'm scared
	if err != nil {
		return nil, err
	}
	match := bson.M{
		"$match": bson.M{"_id": _id},
	}
	unwind := bson.M{
		"$unwind": "$results",
	}
	match2 := bson.M{
		"$match": bson.M{"results._id": _resid},
	}
	unwind2 := bson.M{
		"$unwind": "$questions",
	}
	unwind3 := bson.M{
		"$unwind": "$results.selected",
	}
	redact := bson.M{
		"$redact": bson.M{
			"$cond": bson.A{bson.M{
				"$eq": bson.A{"$questions.qstID", "$results.selected.qstID"},
			}, "$$KEEP", "$$PRUNE"},
		},
	}
	addFields := bson.M{
		"$addFields": bson.M{
			"questions.selected": "$results.selected.option",
			"result_id":          "$results._id",
			"student_id":         "$results.student_id",
		},
	}
	group := bson.M{
		"$group": bson.M{
			"_id": bson.M{
				"_id":        "$_id",
				"title":      "$title",
				"subject":    "$subject",
				"teacher_id": "$teacher_id",
				"student_id": "$student_id",
				"startFrom":  "$startFrom",
				"endAt":      "$endAt",
				"result_id":  "$result_id",
			},
			"result":    bson.M{"$push": "$questions"},
			"questions": bson.M{"$push": "$questions"},
		},
	}
	addFields2 := bson.M{
		"$addFields": bson.M{
			"overall": bson.M{
				"$reduce": bson.M{
					"input": "$result",
					"initialValue": bson.M{
						"blank":     0,
						"correct":   0,
						"incorrect": 0,
					},
					"in": bson.M{
						"blank": bson.M{
							"$add": bson.A{
								"$$value.blank",
								bson.M{
									"$cond": bson.A{
										bson.M{
											"$eq": bson.A{"$$this.selected", ""}}, 1, 0,
									},
								},
							}},
						"correct": bson.M{
							"$add": bson.A{
								"$$value.correct",
								bson.M{
									"$cond": bson.A{
										bson.M{
											"$eq": bson.A{"$$this.selected", "$$this.key"},
										},
										1, 0,
									},
								},
							},
						},
						"incorrect": bson.M{
							"$add": bson.A{
								"$$value.incorrect",
								bson.M{
									"$cond": bson.A{
										bson.M{
											"$and": bson.A{
												bson.M{
													"$ne": bson.A{"$$this.key", "$$this.selected"},
												},
												bson.M{
													"$ne": bson.A{"", "$$this.selected"},
												},
											},
										}, 1, 0,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	unwind4 := bson.M{
		"$unwind": "$result",
	}
	group2 := bson.M{
		"$group": bson.M{
			"_id": "$result.category",
			"detail": bson.M{
				"$first": "$_id",
			},
			"overall": bson.M{
				"$first": "$overall",
			},
			"questions": bson.M{
				"$first": "$questions",
			},
			"categories": bson.M{
				"$push": bson.M{
					"label": "$result.category",
					"correct": bson.M{
						"$sum": bson.M{
							"$cond": bson.A{bson.M{"$eq": bson.A{"$result.key", "$result.selected"}}, 1, 0},
						},
					},
					"blank": bson.M{
						"$sum": bson.M{
							"$cond": bson.A{bson.M{"$eq": bson.A{"", "$result.selected"}}, 1, 0},
						},
					},
					"incorrect": bson.M{
						"$sum": bson.M{
							"$cond": bson.A{bson.M{
								"$and": bson.A{
									bson.M{"$ne": bson.A{"$result.key", "$result.selected"}},
									bson.M{"$ne": bson.A{"", "$result.selected"}}},
							}, 1, 0,
							}},
					},
				},
			},
		},
	}
	group3 := bson.M{
		"$group": bson.M{
			"_id": "$detail",
			"overall": bson.M{
				"$first": "$overall",
			},
			"questions": bson.M{
				"$first": "$questions",
			},
			"categories": bson.M{
				"$push": bson.M{
					"$reduce": bson.M{
						"input":        "$categories",
						"initialValue": bson.M{"correct": 0, "incorrect": 0, "blank": 0},
						"in": bson.M{
							"label":     "$$this.label",
							"correct":   bson.M{"$add": bson.A{"$$this.correct", "$$value.correct"}},
							"incorrect": bson.M{"$add": bson.A{"$$this.incorrect", "$$value.incorrect"}},
							"blank":     bson.M{"$add": bson.A{"$$this.blank", "$$value.blank"}},
						},
					},
				},
			},
		},
	}
	project := bson.M{
		"$project": bson.M{
			"_id":       0,
			"detail":    "$_id",
			"questions": 1,
			"result": bson.M{
				"categories": "$categories",
				"overall":    "$overall",
			},
		},
	}
	pip := bson.A{match, unwind, match2, unwind2, unwind3, redact, addFields, group, addFields2, unwind4, group2, group3, project}
	curs, err := pr.db.Collection("paper").Aggregate(pr.ctx, pip)
	if err != nil {
		return nil, err
	}
	err = curs.All(pr.ctx, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}
