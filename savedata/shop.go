package savedata

type Equip struct {
	Name        string
	IconURL     string
	Description string
	Equipped    bool
	Level       int
	MaxLevel    int
	Power       int
	Projectiles int
	Frequency   float32
	CostFunc    func(float32) float32
}

type Consumable struct {
	Name        string
	IconURL     string
	Description string
	Quantity    int
	MaxQuantity int
	Cost        float32
}

type Upgrades struct {
	HP               float32
	Def              float32
	Pow              float32
	MaxAcceleration  float32
	MaxVelocity      float32
	HPCostFunc       func(float32) float32
	DefPowCostFunc   func(float32) float32
	AccelVelCostFunc func(float32) float32
}

type ShopInfo struct {
	Equipment   []Equip
	Consumables []Consumable
	Upgrades    Upgrades
	Money       float32
}

var defaultShopInfo = ShopInfo{
	Equipment: []Equip{
		{
			Name:        "Sweeper",
			IconURL:     "ui/sweepericon.png",
			Description: "Sweeps away nearby bunnies.",
			Level:       1,
			MaxLevel:    15,
			Power:       5,
			Projectiles: 1,
			Frequency:   5,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Cleaning Bomb",
			IconURL: "ui/cleanbomb.png",
			Description: "Drops a cleaning bomb nearby. Damages bunnies hit, then leaves a clean aura that damages any bunnies inside.",
			Level: 1,
			MaxLevel: 10,
			Power: 2,
			Projectiles: 1,
			Frequency:  20,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Cleaning Spray",
			IconURL: "ui/spray.png",
			Description: "Sprays a mist in front of Scrubbert. Leaves behind a clean aura that damages any bunnies inside.",
			Level: 1,
			MaxLevel: 10,
			Power: 2,
			Projectiles: 1,
			Frequency: 5,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Wax",
			IconURL: "ui/wax.png",
			Description: "Leaves behind a trail of wax, making the floor slippery!",
			Level: 1,
			MaxLevel: 5,
			Power: 0,
			Projectiles: 1,
			Frequency: 1,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Soap Turret",
			IconURL: "ui/soapturret.png",
			Description: "Fires a bar of soap at the nearest dust bunny!",
			Level: 1,
			MaxLevel: 15,
			Power: 3,
			Projectiles: 1,
			Frequency: 3,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Bath Bomb Lobber",
			IconURL: "ui/bomblobber.png",
			Description: "Shoots a bath bomb in a random direction. Explodes on impact with a dust bunnies, damaging nearby bunnies and leaving a clean aura behind.",
			Level: 1,
			MaxLevel: 10,
			Power: 5,
			Projectiles: 1,
			Frequency: 5,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Drone",
			IconURL: "ui/drone.png",
			Description: "A smaller scrubber follows you around, damaging dust bunnies it comes in contact with.",
			Level: 1,
			MaxLevel: 10,
			Power: 5,
			Projectiles: 1,
			Frequency: 1,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Repair Droid",
			IconURL: "ui/droid.png",
			Description: "A repair bot slowly heals Scrubbert!",
			Level: 1,
			MaxLevel: 15,
			Power: 1,
			Projectiles: 0,
			Frequency: 2.5,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Reactive Armor",
			IconURL: "ui/armor.png",
			Description: "If you would die, instead explodes dealing heavy damage to nearby dust bunnies and leaving Scrubbert with some health.",
			Level: 1,
			MaxLevel: 10,
			Power: 15,
			Projectiles: 1,
			Frequency: 30,
			CostFunc: func(x float32) float32 {
				return 0
			},
		},
		{
			Name: "Satchel",
			IconURL: "ui/satchel.png",
			Description: "Adds an extra slot for consumables!",
			Level: 1,
			MaxLevel: 1,
			Power: 0,
			Projectiles: 0,
			Frequency: 0,
			CostFunc: func(x float32) float32 {
				return 0
			},
cd .		},
	},
	Consumables: []Consumable{
		{},
		{},
		{},
	},
	Upgrades: Upgrades{
		HP:              10,
		Def:             1,
		Pow:             1,
		MaxAcceleration: 5,
		MaxVelocity:     5,
		HPCostFunc: func(x float32) float32 {
			return 0
		},
		DefPowCostFunc: func(x float32) float32 {
			return 0
		},
		AccelVelCostFunc: func(x float32) float32 {
			return 0
		},
	},
	Money: 0,
}
