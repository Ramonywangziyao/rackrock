package starter

import (
	"fmt"
	"github.com/farmerx/gorsa"
	"log"
	"rackrock/config"
	"rackrock/context"
	"rackrock/model"
	"rackrock/router"
	"rackrock/utils"
)

func serverRun() {
	//  permission
	context.AddBeforeHandler(context.PermissionHandle)
	context.AddAfterHandler(context.LoggerHandle)

	var engine = router.InitRouter()

	if err := gorsa.RSA.SetPrivateKey(model.Pirvatekey); err != nil {
		log.Fatalln(`set private key :`, err)
	}

	var port = "8080"
	if !utils.IsEmptyStr(config.Cfg.Port) {
		port = config.Cfg.Port
	}

	if err := engine.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(fmt.Sprintf("run server err: %s", err.Error()))
	}
}
