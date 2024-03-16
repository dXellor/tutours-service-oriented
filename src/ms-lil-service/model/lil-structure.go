package model

type Lil struct {
	//ID uuid.UUID // angy, why this, angy // i dont wanna shoo
	Id             int    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key;auto_increment;"`
	Name, Nickname string `gorm:"type:string;not null;default: null"` // cappital letter = exported field
	Age            int    `json: "Oldness"`
}

// func (l *Lil) BeforeCreate(db *gorm.DB) error {
// 	l.ID = uuid.New()
// 	return nil
// }

func GenerateNickname(l Lil) string {
	l.Nickname = string(l.Name[1])
	return string(l.Name[1])
}

func (l Lil) GenerateNicknameButBetter() string {
	l.Nickname = string(l.Name[1])
	return string(l.Name[1])
}

func (l *Lil) GenerateNicknameButEvenBetter() string {
	l.Nickname = string(l.Name[1])
	return string(l.Name[1])
}

type structure interface { // hidden outside of package
	AgePlus() int
	AgeMinus() int
}

type Structure interface {
	AgePlus() int
	AgeMinus() int
}

func (l Lil) AgePlus() int {
	return l.Age + 1
}

func (l Lil) AgeMinus() int {
	return l.Age + 1
}