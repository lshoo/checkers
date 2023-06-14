package keeper

import (
	"context"
	"strconv"
	
	rules "github.com/lshoo/checkers/x/checkers/rules"
	"github.com/lshoo/checkers/x/checkers/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	// _ = ctx
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)

	if !found {
		panic("SystemInfo not found")
	}

	nextIndex := strconv.FormatUint(systemInfo.NextId, 10)

	newGame := rules.New()
	storeGame := types.StoredGame{
		Index: nextIndex,
		Board: newGame.String(),
		Turn:  rules.PieceStrings[newGame.Turn],
		Black: msg.Black,
		Red:   msg.Red,
	}

	err := storeGame.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredGame(ctx, storeGame)

	systemInfo.NextId++

	k.Keeper.SetSystemInfo(ctx, systemInfo)

	return &types.MsgCreateGameResponse{
		GameIndex: nextIndex,
	}, nil
}
