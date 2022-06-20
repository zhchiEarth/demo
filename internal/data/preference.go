package data

import (
	"compound/internal/data/ent"
	"compound/internal/data/ent/preference"
	"context"
	"github.com/pkg/errors"

	"compound/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.PreferenceRepo = (*preferenceRepo)(nil)

type preferenceRepo struct {
	data *Data
	log  *log.Helper
}

// NewPreferenceRepo .
func NewPreferenceRepo(data *Data, logger log.Logger) biz.PreferenceRepo {
	return &preferenceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *preferenceRepo) FindByKey(ctx context.Context, key string) (*biz.Preference, error) {
	p, err := r.data.db.Preference.Query().
		Where(preference.Key(key)).
		First(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &biz.Preference{
		ID:    p.ID,
		Key:   p.Key,
		Value: p.Value,
	}, nil
}

func (r *preferenceRepo) Save(ctx context.Context, key string, value string) (int, error) {
	p, err := r.data.db.Preference.Query().
		Where(preference.Key(key)).
		First(ctx)
	if !ent.IsNotFound(err) && err != nil {
		return 0, errors.Wrap(err, "")
	}
	// 区块高度 大于数据库的区块高度
	if p == nil {
		return r.Create(ctx, key, value)
	}
	return r.Update(ctx, p.ID, value)
}

func (r *preferenceRepo) Create(ctx context.Context, key string, value string) (int, error) {
	p, err := r.data.db.Preference.Create().
		SetKey(key).
		SetValue(value).
		Save(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return p.ID, nil
}

func (r *preferenceRepo) Update(ctx context.Context, id int, value string) (int, error) {
	p, err := r.data.db.Preference.UpdateOneID(id).
		SetValue(value).
		Save(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	return p.ID, nil
}
