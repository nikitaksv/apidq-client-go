package address

const (
	// QltUniqUnique Uniq/Уникален: предлагается 1 эталонный адрес
	QltUniqUnique QltUniq = 0
	// QltUniqDoubtful Doubtful/Сомнителен: предлагается несколько эталонных адресов близких по написанию (возможен выбор)
	QltUniqDoubtful
	// QltUniqNotUniq Not unique/Неуникален: есть несколько эталонных записей, в равной степени соответствующих исходному адресу
	QltUniqNotUniq
)

const (
	// QltActualityActual Actual/Найдено по актуальной записи: название и административное подчинение, указанные в разбираемом адресе, соответствуют эталонному
	QltActualityActual QltActuality = iota
	// QltActualityRename Rename/Переименование: устаревшее название одного из адресных элементов, указанных в разбираемом адресе
	QltActualityRename
	// QltActualityReassignment Reassignment/Переподчинение: административное подчинение, указанное в разбираемом адресе, устарело
	QltActualityReassignment
)

const (
	// QltUndefNo No/Нет
	QltUndefNo QltUndef = iota
	// QltUndefInsignificant Insignificant/Малозначимый: информация, не влияющая на результаты распознавания при ручной проверке
	QltUndefInsignificant
	// QltUndefSignificant Significant/Значимый: информация, которая при ручной проверке может повлиять на результат сравнения разбираемого адреса с эталоном
	QltUndefSignificant
)

const (
	// QltLvlRegion To the region(state)/До региона
	QltLvlRegion QltLvl = iota + 1
	// QltLvlDistrict To the district/До района
	QltLvlDistrict
	// QltLvlCity To the city/До города
	QltLvlCity
	// QltLvlCityArea To the district in the city/До района в городе
	QltLvlCityArea
	// QltLvlSettlement To the settlement/До населенного пункта
	QltLvlSettlement
	// QltLvlPlanStruct To the planning structure/До планировочной структуры
	QltLvlPlanStruct
	// QltLvlStreet To the street/До улицы
	QltLvlStreet
	// QltLvlHouse To the house/До дома
	QltLvlHouse
)

const (
	// QltHouseNotFound Not found variants/Не найдено вариантов
	QltHouseNotFound QltHouse = iota
	_
	_
	// QltHouseExact Exact match of the house by reference/Точное определение дома по эталону
	QltHouseExact
	// QltHousePartial Partial house match by reference/Частичное определение дома по эталону
	QltHousePartial
	_
	_
	_
	_
	// QltHouseNonHouse The parsed address is missing a house number/В разбираемом адресе отсутствует номер дома
	QltHouseNonHouse
)

const (
	// QltGeoRegion To the region(state)/До региона
	QltGeoRegion QltGeo = iota + 1
	// QltGeoDistrict To the district/До района
	QltGeoDistrict
	// QltGeoCity To the city/До города
	QltGeoCity
	// QltGeoCityArea To the district in the city/До района в городе
	QltGeoCityArea
	// QltGeoSettlement To the settlement/До населенного пункта
	QltGeoSettlement
	// QltGeoPlanStruct To the planning structure/До планировочной структуры
	QltGeoPlanStruct
	// QltGeoStreet To the street/До улицы
	QltGeoStreet
	// QltGeoHouse To the house/До дома
	QltGeoHouse
)

type Quality struct {
	// Уровень уникальности найденного адреса
	Unique QltUniq `json:"unique"`
	// Статус актуальности исходного адреса
	Actuality QltActuality `json:"actuality"`
	// Разбор неадресной информации в исходном адресе
	Undefined QltUndef `json:"undefined"`
	// Уровень, до которого произведено сравнение исходного адреса с эталоном
	Level QltLvl `json:"level"`
	// Степень совпадения номера дома в исходном адресе с эталоном
	House QltHouse `json:"house"`
	// Уровень, до которого разобраны координаты адреса
	Geo QltGeo `json:"geo"`
}

// QltUniq Уровень уникальности найденного адреса
type QltUniq int

// Values Все возможные значения
func (q QltUniq) Values() []QltUniq {
	return []QltUniq{QltUniqUnique, QltUniqDoubtful, QltUniqNotUniq}
}

// QltActuality Статус актуальности исходного адреса
type QltActuality int

// Values Все возможные значения
func (q QltActuality) Values() []QltActuality {
	return []QltActuality{QltActualityActual, QltActualityRename, QltActualityReassignment}
}

// QltUndef Разбор неадресной информации в исходном адресе
type QltUndef int

// Values Все возможные значения
func (q QltUndef) Values() []QltUndef {
	return []QltUndef{QltUndefNo, QltUndefInsignificant, QltUndefSignificant}
}

// QltLvl Уровень, до которого произведено сравнение исходного адреса с эталоном
type QltLvl int

// Values Все возможные значения
func (q QltLvl) Values() []QltLvl {
	return []QltLvl{
		QltLvlRegion,
		QltLvlDistrict,
		QltLvlCity,
		QltLvlCityArea,
		QltLvlSettlement,
		QltLvlPlanStruct,
		QltLvlStreet,
		QltLvlHouse,
	}
}

// QltHouse Степень совпадения номера дома в исходном адресе с эталоном
type QltHouse int

// Values Все возможные значения
func (q QltHouse) Values() []QltHouse {
	return []QltHouse{QltHouseNotFound, QltHouseExact, QltHousePartial, QltHouseNonHouse}
}

// QltGeo Уровень, до которого разобраны координаты адреса
type QltGeo int

// Values Все возможные значения
func (q QltGeo) Values() []QltGeo {
	return []QltGeo{
		QltGeoRegion,
		QltGeoDistrict,
		QltGeoCity,
		QltGeoCityArea,
		QltGeoSettlement,
		QltGeoPlanStruct,
		QltGeoStreet,
		QltGeoHouse,
	}
}
