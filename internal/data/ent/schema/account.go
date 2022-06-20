package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Comment("用户地址").Unique(),
		field.Bool("has_borrowed").Default(false),
		field.Uint32("count_liquidated").Default(0).Comment("Count user has been liquidated"),
		field.Uint32("count_liquidator").Default(0).Comment("Count user has liquidated others"),
		field.String("health").Default("0").Comment("用户健康度 total_collateral_value_in_eth / total_borrow_value_in_eth。如果此值小于 1.0，则该帐户将被清算"),
		field.String("total_collateral_value_in_usd").Default("0").Comment("账户提供的所有抵押品的价值。计算为持有的cToken • 汇率 • 抵押因子。注：资产可以提供并获得利息，不计入抵押品"),
		field.String("total_borrow_value_in_usd").Default("0").Comment("累积利息的所有未偿还借款的价值"),
		field.Uint64("block_number").Default(0).Comment("区块高度"),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", AccountCToken.Type),
	}
}
