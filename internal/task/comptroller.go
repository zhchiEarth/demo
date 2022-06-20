package task

import "github.com/hasura/go-graphql-client"

type Comptroller struct {
    Id                   graphql.String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
    PriceOracle          graphql.String `protobuf:"bytes,2,opt,name=price_oracle,json=priceOracle,proto3" json:"price_oracle,omitempty"`
    CloseFactor          graphql.String `protobuf:"bytes,3,opt,name=close_factor,json=closeFactor,proto3" json:"close_factor,omitempty"`
    LiquidationIncentive graphql.String `protobuf:"bytes,4,opt,name=liquidation_incentive,json=liquidationIncentive,proto3" json:"liquidation_incentive,omitempty"`
    MaxAssets            graphql.String `protobuf:"bytes,5,opt,name=max_assets,json=maxAssets,proto3" json:"max_assets,omitempty"`
}