package controllers

import (
	"gen/api/pt"
	"goslib/logger"
	"goslib/player"
	"goslib/timertask"
	"time"
)

type EquipsController struct {
}

func (*EquipsController) Load(ctx *player.Player, params *pt.EquipLoadParams) (string, interface{}) {
	runAt := time.Now().Add(5 * time.Second).Unix()
	err := timertask.Add("fake_task_id", runAt, ctx.PlayerId, "EquipUnLoadParams", &pt.EquipUnLoadParams{
		PlayerID: "player_id",
		EquipId:  "equip_id",
	})
	if err != nil {
		logger.ERR("add timertask failed: ", err)
	}

	return pt.PT_EquipLoadResponse, &pt.EquipLoadResponse{PlayerID: "player_id", EquipId: "equip_id", Level: 10}
}

func (*EquipsController) UnLoad(ctx *player.Player, params *pt.EquipUnLoadParams) (string, interface{}) {
	logger.INFO("UnLoad equips")
	return pt.PT_EquipLoadResponse, &pt.EquipLoadResponse{PlayerID: "player_id", EquipId: "equip_id", Level: 10}
}
