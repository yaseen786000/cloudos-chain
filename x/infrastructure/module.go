package infrastructure

import "github.com/yaseen786000/cloudos-chain/x/infrastructure/keeper"

type AppModule struct {
	Keeper *keeper.Keeper
}

func NewAppModule(
	k *keeper.Keeper,
) AppModule {

	return AppModule{
		Keeper: k,
	}
}
