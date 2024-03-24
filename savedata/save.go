package savedata

type SaveData struct {
	Shop    ShopInfo
	Options OptionInfo
}

func Load() (*SaveData, error) {
	save := &SaveData{}
	return save, nil
}

func (s *SaveData) Save() error {
	return nil
}
