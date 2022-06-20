package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// AccountCToken holds the schema definition for the AccountCToken entity.
type AccountCToken struct {
	ent.Schema
}

// Fields of the AccountCToken.
func (AccountCToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("user").Default("").Comment("用户地址"),
		field.String("address").Default("").Comment("ctoken地址"),
		field.String("symbol").Default("").Comment("ctoken symbol"),
		field.String("name").Default("").Comment("ctoken name"),
		field.Bool("entered_market").Default(false).Comment("进入市场的标志"),
		field.String("ctoken_balance").Default("0").Comment("ctoken 余额"),
		field.String("stored_borrow_balance").Default("0").Comment("借款的余额"),
		field.String("borrow_index").Default("0").Comment("用户上一次的贷款指数"),
		field.String("total_underlying_supplied").Default("0").Comment("总的token 供应量"),
		field.String("total_underlying_redeemed").Default("0").Comment("累计提现"),
		field.String("total_underlying_borrowed").Default("0").Comment("累计借款"),
		field.String("total_underlying_repaid").Default("0").Comment("累计还款"),
		field.String("supply_balance_underlying").Default("0").Comment("转换为基础代币的 cTokenBalance * market.exchangeRate"),
		field.String("borrow_balance_underlying").Default("0").Comment("token的借款 borrowBalanceUnderlying = storedBorrowBalance * market.borrowIndex / accountBorrowIndex"),
		field.String("lifetime_supply_interest_accrued").Default("0").Comment("累积的供应利息 lifetimeSupplyInterestAccrued = supplyBalanceUnderlying - totalUnderlyingSupplied + totalUnderlyingRedeemed"),
		field.String("lifetime_borrow_interest_accrued").Default("0").Comment("生命周期内应计的借款利息金额 lifetimeSupplyInterestAccrued = supplyBalanceUnderlying - totalUnderlyingSupplied + totalUnderlyingRedeemed"),
		field.String("safe_withdraw_amount_underlying").Default("0").Comment("可以提取的供应量，以使用户的健康保持在 1.25 或更高"),
		field.String("collateral_value_in_usd").Default("0").Comment("抵押物价值，usd计价"),
		field.String("borrow_value_in_usd").Default("0").Comment("贷款价值，usd计价"),
		field.String("underlying_price_usd").Default("0").Comment("usd价格"),
		field.Uint64("block_number").Default(0).Comment("区块高度"),
	}
}

func (AccountCToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the AccountCToken.
func (AccountCToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("tokens").
			Unique(),
	}
}
