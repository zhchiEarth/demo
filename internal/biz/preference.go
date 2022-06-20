package biz

import (
    "compound/internal/data/ent"
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "strconv"
)

const AccountCTokenLastBlockNumber = "account_ctoken_last_block_num"

type Preference struct {
    ID int
    Key string
    Value string
}

// PreferenceRepo is a Preference repo.
type PreferenceRepo interface {
    FindByKey(context.Context, string) (*Preference, error)
    Save(context.Context, string, string) (int, error)
}

// PreferenceUseCase is a Preference UseCase.
type PreferenceUseCase struct {
    repo PreferenceRepo
    log  *log.Helper
}

// NewPreferenceUseCase new a Preference UseCase.
func NewPreferenceUseCase(repo PreferenceRepo, logger log.Logger) *PreferenceUseCase {
    return &PreferenceUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PreferenceUseCase) SaveAccountCTokenLastBlockNumber(ctx context.Context, blockNumber uint64) (uint64, error) {
    id, err := uc.repo.Save(ctx, AccountCTokenLastBlockNumber, strconv.FormatUint(blockNumber, 10))
    if err != nil {
        uc.log.Errorf("SaveAccountCTokenLastBlockNumber 保存失败,Key:%s, value:%d, err:%s", AccountCTokenLastBlockNumber,
            blockNumber, err)
        return 0, err
    }
    return uint64(id), nil
}

func (uc *PreferenceUseCase) GetAccountCTokenLastBlockNumber(ctx context.Context) (uint64, error) {
    r, err := uc.repo.FindByKey(ctx, AccountCTokenLastBlockNumber)
    if ent.IsNotFound(err) {
        return 0, nil
    }
    if err != nil {
        uc.log.Errorf("GetAccountCTokenLastBlockNumber 失败,Key:%s, err:%s",
            AccountCTokenLastBlockNumber, err)
        return 0, err
    }
    val, err := strconv.ParseUint(r.Value, 10, 64)
    if err != nil {
        return 0, err
    }

    return val, nil
}