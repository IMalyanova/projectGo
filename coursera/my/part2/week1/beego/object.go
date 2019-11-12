package main

import "encoding/json"

func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
}



func (o *ObjectController) Get() {

	objectId := o.Ctx.Input.Param(":objectId")

	if objectId != "" {
		ob, err := models.GetOne(objectId)

		if err != nil {
			o.Data["json"] = err.Error()
		} else {

		}
	}
}
