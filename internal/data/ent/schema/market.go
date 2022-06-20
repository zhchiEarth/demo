package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Market holds the schema definition for the Market entity.
type Market struct {
	ent.Schema
}

// Fields of the Market.
func (Market) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Comment("ctoken地址 唯一的").Unique(),
		field.String("symbol").Default("").Comment("ctoken symbol"),
		field.String("name").Default("").Comment("ctoken name"),
		field.String("borrow_index").Default("").Comment("借款利息"),
		field.String("borrow_rate").Default("").Comment("借款利率"),
		field.String("supply_rate").Default("").Comment("供应利率"),
		field.String("cash").Default("").Comment("池子的流动性 cash"),
		field.String("collateral_factor").Default("").Comment("抵押率"),
		field.String("exchange_rate").Default("").Comment(" 兑换率"),
		field.String("reserve_factor").Default("").Comment("储备金率"),
		field.String("reserves").Default("").Comment("储备金"),
		field.String("total_borrows").Default("").Comment("总借款"),
		field.String("total_supply").Default("").Comment("总供应"),
		field.String("underlying_address").Default("").Comment("token 地址"),
		field.String("underlying_name").Default("").Comment("token 名称"),
		field.String("underlying_symbol").Default("").Comment("token 简称"),
		field.String("underlying_price").Default("").Comment("token 兑换成eth的价格"),
		field.Uint32("underlying_decimals").Default(0).Comment("token 位数"),
		field.String("underlying_price_usd").Default("").Comment("token usd价格"),
		field.Uint64("block_number").Default(0).Comment("区块高度"),
		field.Uint32("block_timestamp").Default(0).Comment("区块时间戳"),
	}
}

func (Market) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Market.
func (Market) Edges() []ent.Edge {
	return nil
}
