package entity

//Skin contains data about a skin
type Skin struct {
	skinID       string
	skinData     string
	capeData     string
	geometryName string
	geometryData string
}

//New returns a new skin
func (*Skin) New(skinID, skinData, capeData, gemoetryName, geometryData string) *Skin {
	return &Skin{
		skinID:       skinID,
		skinData:     skinData,
		capeData:     capeData,
		geometryName: geometryName,
		geometryData: geometryData,
	}
}

//IsValid returns whether a skin is valid or not
func (skin *Skin) IsValid() bool {
	if skin.skinID != "" && (len(skin.skinData) == 16384 || len(skin.skinData) == 8192 || len(skin.skinData) == 65535) && (skin.capeData == "" || len(skin.capeData) == 8192) {
		return true
	}
	return false
}

//GetSkinID returns a skin's ID
func (skin *Skin) GetSkinID() string {
	return skin.skinID
}

//GetSkinData returns a skin's data
func (skin *Skin) GetSkinData() string {
	return skin.skinData
}

//GetCapeData returns a skin's cape data
func (skin *Skin) GetCapeData() string {
	return skin.capeData
}

//GetGeometryName returns a skin's geometry name
func (skin *Skin) GetGeometryName() string {
	return skin.geometryName
}

//GetGeometryData returns a skin's geometry data
func (skin *Skin) GetGeometryData() string {
	return skin.geometryData
}
