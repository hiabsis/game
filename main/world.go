package main

/**
 *新手村地图
 */

type VillageMap struct {
	name    string
	context *WordContext
	body    []string
}

func getVillageBody() []string {
	return []string{"武器铺"}
}
func getWeaponBody() []string {
	return []string{"1.小刀\t[100]", "2.大剑\t[100]", "3.盔甲\t[50]"}
}
func getHomeBody() []string {
	return []string{"1.睡觉\t[100]", "2.仓库\t[100]", "3.宠物\t[50]"}
}
func (v *VillageMap) getMapBody() []string {
	return v.body
}

func (v *VillageMap) enter(wc *WordContext) {
	v.context = wc

	go func() {
		for {
			cmd := scanner()
			switch cmd {
			case "武器铺":
				v.body = getWeaponBody()
				wc.draw()
			}

		}
	}()
}
func (v *VillageMap) exit() {
}

func (v *VillageMap) queryMapName() string {
	return v.name
}
